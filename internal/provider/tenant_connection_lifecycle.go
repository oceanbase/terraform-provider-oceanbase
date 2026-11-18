package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

const (
	tenantAllocatingInternetAddress     = "ALLOCATING_INTERNET_ADDRESS"
	tenantPendingOfflineInternetAddress = "PENDING_OFFLINE_INTERNET_ADDRESS"
)

// The generated SDK decodes both null and {} list elements into a zero-value
// TenantConnectionDTO. Inspect its retained response only for that ambiguous case.
func connectionResponseHasNullEntry(response *http.Response) (bool, error) {
	if response == nil || response.Body == nil {
		return false, errors.New("tenant connection query response has no body")
	}
	var body struct {
		Data []json.RawMessage `json:"data"`
	}
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return false, fmt.Errorf("decode tenant connection query response: %w", err)
	}
	return len(body.Data) == 1 && bytes.Equal(bytes.TrimSpace(body.Data[0]), []byte("null")), nil
}

func confirmPublicConnectionAbsent(ctx context.Context, connection *model.TenantConnection, client *obcloudsdk.APIClient) error {
	result, response, err := client.MultiCloudOpenAPI.DescribeTenantPublicAddress(ctx, connection.ClusterId.ValueString(), connection.TenantId.ValueString()).RequestId("").Execute()
	if err != nil || response != nil {
		api.LogAPIResult(ctx, result, response, err)
	}
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		return fmt.Errorf("could not confirm tenant public address is absent: %v", err)
	}
	address := result.Data
	if id := address.GetTenantId(); id != "" && id != connection.TenantId.ValueString() {
		return errors.New("tenant public address belongs to a different tenant")
	}
	if kind := address.GetNetworkType(); kind != "" && !strings.EqualFold(kind, model.ConnectionTypeInternet) {
		return fmt.Errorf("tenant public address has unexpected network type %q", kind)
	}
	if id := address.GetAddressId(); id != "" {
		return fmt.Errorf("tenant public address still reports address %q", id)
	}
	if address.GetInternetAddress() != "" || address.GetInternetDomain() != "" || address.GetInternetPort() != 0 ||
		address.GetInternetAddressStatus() != "" || address.GetAddressStatus() != "" {
		return errors.New("tenant public address contains address details without an id")
	}
	return nil
}

func tenantStatusForConnection(ctx context.Context, connection *model.TenantConnection, client *obcloudsdk.APIClient) (string, error) {
	result, response, err := client.MultiCloudOpenAPI.DescribeTenant(ctx, connection.ClusterId.ValueString(), connection.TenantId.ValueString()).RequestId("").Execute()
	if err != nil || response != nil {
		api.LogAPIResult(ctx, result, response, err)
	}
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		return "", fmt.Errorf("could not check tenant status for public address: %v", err)
	}
	if result.Data.GetTenantId() != connection.TenantId.ValueString() || strings.TrimSpace(result.Data.GetStatus()) == "" {
		return "", errors.New("tenant status response omitted the expected tenant id or status")
	}
	return result.Data.GetStatus(), nil
}

func tenantConnectionReleasePending(ctx context.Context, connection *model.TenantConnection, client *obcloudsdk.APIClient) (bool, error) {
	status, err := tenantStatusForConnection(ctx, connection, client)
	return strings.EqualFold(status, tenantPendingOfflineInternetAddress), err
}

func waitForTenantConnectionStatus(ctx context.Context, connection *model.TenantConnection, client *obcloudsdk.APIClient, timeout time.Duration, pendingStatus string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		status, err := tenantStatusForConnection(ctx, connection, client)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		if strings.EqualFold(status, pendingStatus) {
			return retry.RetryableError(fmt.Errorf("tenant status is still %s", pendingStatus))
		}
		return nil
	})
}

// Removal is complete only when the address is gone and its tenant has left
// the asynchronous release state. Keep the polled copy local to the callback.
func readTenantConnectionRemoval(ctx context.Context, connection *model.TenantConnection, client *obcloudsdk.APIClient) error {
	current := *connection
	err := refreshTenantConnection(ctx, &current, client)
	if !errors.Is(err, errResourceNotFound) {
		return err
	}
	if strings.EqualFold(connection.Type.ValueString(), model.ConnectionTypeInternet) {
		pending, statusErr := tenantConnectionReleasePending(ctx, connection, client)
		if statusErr != nil {
			return statusErr
		}
		if pending {
			return nil
		}
	}
	return err
}
