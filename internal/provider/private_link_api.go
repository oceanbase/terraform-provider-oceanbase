package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type privateLinkTarget struct {
	clusterID string
	tenantID  string
	projectID string
	addressID string
}

func (t privateLinkTarget) validate(requireAddress bool) error {
	if strings.TrimSpace(t.clusterID) == "" || strings.TrimSpace(t.tenantID) == "" {
		return errors.New("cluster_id and tenant_id must not be empty")
	}
	if requireAddress && strings.TrimSpace(t.addressID) == "" {
		return errors.New("PrivateLink address ID must not be empty")
	}
	return nil
}

func privateLinkClient(p *oceanbaseProvider, projectID string) *obcloudsdk.APIClient {
	client := p.clientGenerator.GetClient()
	if projectID != "" {
		client.GetConfig().AddDefaultHeader(api.HeaderProjectId, projectID)
	}
	return client
}

func privateLinkAPIError(operation string, response *http.Response, err error, success bool, code, message string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	if response == nil {
		return fmt.Errorf("%s: no HTTP response", operation)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: unexpected HTTP status %d", operation, response.StatusCode)
	}
	if !success {
		return fmt.Errorf("%s: %s: %s", operation, code, message)
	}
	return nil
}

func describePrivateLinks(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget) ([]obcloudsdk.TenantConnectionDTO, error) {
	req := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.DescribeTenantPrivateLink(ctx, target.clusterID, target.tenantID).RequestId("")
	if target.addressID != "" {
		req = req.AddressId(target.addressID)
	}
	result, response, err := req.Execute()
	if result == nil {
		return nil, privateLinkAPIError("describe PrivateLink", response, err, false, "", "empty response")
	}
	if err := privateLinkAPIError("describe PrivateLink", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage()); err != nil {
		if result.GetErrorCode() == "Tenant.Address.NotExist" {
			return nil, errResourceNotFound
		}
		return nil, err
	}
	if !result.HasData() {
		return nil, errors.New("describe PrivateLink returned no data")
	}
	addresses := result.GetData()
	for _, address := range addresses {
		if tenantID := address.GetTenantId(); tenantID != "" && tenantID != target.tenantID {
			return nil, errors.New("describe PrivateLink returned an address for a different tenant")
		}
		if networkType := address.GetNetworkType(); networkType != "" && !strings.EqualFold(networkType, "INTRANET") {
			return nil, fmt.Errorf("describe PrivateLink returned unexpected network type %q", networkType)
		}
		if target.addressID != "" {
			if address.GetAddressId() != target.addressID {
				return nil, fmt.Errorf("describe PrivateLink returned a different address from %q", target.addressID)
			}
		}
		for _, entry := range address.GetPrivateLinkList() {
			if entry.GetPrivateLinkStatus() == "" {
				return nil, errors.New("describe PrivateLink returned an entry without a status")
			}
		}
	}
	return addresses, nil
}

func privateLinkEntries(addresses []obcloudsdk.TenantConnectionDTO) []obcloudsdk.TenantPrivateLinkDTO {
	var entries []obcloudsdk.TenantPrivateLinkDTO
	for _, address := range addresses {
		entries = append(entries, address.GetPrivateLinkList()...)
	}
	return entries
}

func privateLinkActive(status string) bool {
	return status != "" && !strings.EqualFold(status, "UNCREATED") && !strings.EqualFold(status, "DELETED")
}

func waitPrivateLink(ctx context.Context, timeout time.Duration, check func(context.Context) (bool, error)) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		done, err := check(ctx)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		if !done {
			return retry.RetryableError(errors.New("PrivateLink operation is still in progress"))
		}
		return nil
	})
}

func privateLinkTenantReady(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget) (bool, error) {
	result, response, err := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.DescribeTenant(ctx, target.clusterID, target.tenantID).RequestId("").Execute()
	if result == nil {
		return false, privateLinkAPIError("describe tenant", response, err, false, "", "empty response")
	}
	if err := privateLinkAPIError("describe tenant", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage()); err != nil {
		if result.GetErrorCode() == "Tenant.NotExists" {
			return true, nil
		}
		return false, err
	}
	if !result.HasData() || result.Data.GetTenantId() != target.tenantID {
		return false, errors.New("describe tenant did not return the expected tenant")
	}
	switch strings.ToUpper(result.Data.GetStatus()) {
	case "ONLINE":
		return true, nil
	case "DELETING_MASTER_ADDRESS", "DELETING_READONLY_ADDRESS", "DELETING_READWRITE_ADDRESS", "ALLOCATING_MASTER_ADDRESS":
		return false, nil
	default:
		return false, fmt.Errorf("tenant is not ready after PrivateLink deletion: %q", result.Data.GetStatus())
	}
}

func createPrivateLinkService(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget, userAccount string, proxySetting *obcloudsdk.PrivateLinkProxySetting) (string, error) {
	body := obcloudsdk.NewCreateTenantPlServiceRequestWithDefaults()
	body.SetUserAccount(userAccount)
	if proxySetting != nil {
		body.SetProxySetting(*proxySetting)
	}
	result, response, err := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.CreatePrivateLinkService(ctx, target.clusterID, target.tenantID).RequestId("").Body(*body).Execute()
	if result == nil {
		return "", privateLinkAPIError("create PrivateLink service", response, err, false, "", "empty response")
	}
	if err := privateLinkAPIError("create PrivateLink service", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage()); err != nil {
		return "", err
	}
	if !result.HasData() || result.Data.GetAddressId() == "" {
		return "", errors.New("create PrivateLink service succeeded without an address ID; inspect the tenant before retrying")
	}
	return result.Data.GetAddressId(), nil
}

func connectPrivateLink(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget, endpointID, endpointAddress string) error {
	body := obcloudsdk.NewConnectTenantPlServiceRequestWithDefaults()
	body.SetAddressId(target.addressID)
	body.SetEndpointId(endpointID)
	if endpointAddress != "" {
		body.SetUserEndpointAddress(endpointAddress)
	}
	result, response, err := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.ConnectPrivateLinkService(ctx, target.clusterID, target.tenantID).RequestId("").Body(*body).Execute()
	if result == nil {
		return privateLinkAPIError("connect PrivateLink endpoint", response, err, false, "", "empty response")
	}
	return privateLinkAPIError("connect PrivateLink endpoint", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage())
}

func confirmPrivateLink(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget, endpointID string) error {
	body := obcloudsdk.NewConfirmTenantPlServiceRequestWithDefaults()
	body.SetAddressId(target.addressID)
	body.SetEndpointId(endpointID)
	result, response, err := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.ConfirmPrivatelinkConnection(ctx, target.clusterID, target.tenantID).RequestId("").Body(*body).Execute()
	if result == nil {
		return privateLinkAPIError("confirm PrivateLink endpoint", response, err, false, "", "empty response")
	}
	if err := privateLinkAPIError("confirm PrivateLink endpoint", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage()); err != nil {
		return err
	}
	if result.HasData() && !result.GetData() {
		return errors.New("confirm PrivateLink endpoint returned false")
	}
	return nil
}

func deletePrivateLink(ctx context.Context, p *oceanbaseProvider, target privateLinkTarget, endpointID, status string) error {
	req := privateLinkClient(p, target.projectID).MultiCloudOpenAPI.DeletePrivatelinkEndpoint(ctx, target.clusterID, target.tenantID).RequestId("").AddressId(target.addressID)
	if endpointID != "" {
		req = req.EndpointId(endpointID)
	}
	if status != "" {
		req = req.Status(status)
	}
	result, response, err := req.Execute()
	if result == nil {
		return privateLinkAPIError("delete PrivateLink", response, err, false, "", "empty response")
	}
	if err := privateLinkAPIError("delete PrivateLink", response, err, result.GetSuccess(), result.GetErrorCode(), result.GetErrorMessage()); err != nil {
		return err
	}
	if result.HasData() && !result.GetData() {
		return errors.New("delete PrivateLink returned false")
	}
	return nil
}
