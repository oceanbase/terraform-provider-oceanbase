/*
 * Copyright 2025 OceanBase
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

func addConfigureProviderErr(diagnostics *diag.Diagnostics) {
	diagnostics.AddError(
		"Provider not configured",
		"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
	)
}

func refreshCluster(ctx context.Context, cluster *model.Cluster, client *obcloudsdk.APIClient) error {
	if cluster.Id.IsNull() || cluster.Id.IsUnknown() || cluster.Id.ValueString() == "" {
		return errors.New("Cluster id is invalid")
	}
	describeInstanceResult, describeInstanceResponse, describeInstanceErr := client.MultiCloudOpenAPI.DescribeInstance(ctx, cluster.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, describeInstanceResult, describeInstanceResponse, describeInstanceErr)
	detailOK := describeInstanceErr == nil && describeInstanceResponse != nil && describeInstanceResponse.StatusCode == http.StatusOK && describeInstanceResult.GetSuccess() && describeInstanceResult.Data != nil
	detailMissing := (describeInstanceResponse != nil && describeInstanceResponse.StatusCode == http.StatusNotFound) ||
		(describeInstanceErr == nil && describeInstanceResult != nil && describeInstanceResult.GetErrorCode() == "Instance.NotExists")
	if !detailOK && !detailMissing {
		return fmt.Errorf("failed to describe cluster: error %v, response %v, error code %s", describeInstanceErr, describeInstanceResponse, describeInstanceResult.GetErrorCode())
	}
	projectID := cluster.ProjectId.ValueString()
	if detailOK {
		detail := describeInstanceResult.Data
		if detail.GetInstanceId() != cluster.Id.ValueString() {
			return errors.New("DescribeInstance returned a different or missing instance ID")
		}
		if projectID != "" && detail.GetProjectId() != "" && projectID != detail.GetProjectId() {
			return errors.New("DescribeInstance returned a different project ID")
		}
		if projectID == "" {
			projectID = detail.GetProjectId()
		}
	}
	var item *obcloudsdk.DescribeInstancesResponseV2OpenAPI
	var listErr error
	if detailMissing {
		item, listErr = findClusterInProject(ctx, client, projectID, cluster.Id.ValueString())
	} else {
		item, listErr = findClusterByID(ctx, client, projectID, cluster.Id.ValueString())
	}
	if errors.Is(listErr, errResourceNotFound) {
		if detailMissing {
			return errResourceNotFound
		}
		return errors.New("cluster exists in detail but is absent from project inventory")
	}
	if listErr != nil {
		return listErr
	}
	if !detailOK {
		return errors.New("cluster is present in project inventory but detail reports it missing")
	}
	isClusterType := func(instanceType string) bool {
		switch strings.ToUpper(instanceType) {
		case "CLUSTER", "PRIMARY_STANDBY_CLUSTER", "ANALYTICAL_CLUSTER", "KV_CLUSTER", "SHARED_STORAGE_CLUSTER", "K8S_CLUSTER", "K8S_KV_CLUSTER", "FLAG_SHIP":
			return true
		default:
			return false
		}
	}
	if !isClusterType(describeInstanceResult.Data.GetInstanceType()) || !isClusterType(item.GetInstanceType()) {
		return errors.New("Instance type mismatch")
	}
	cluster.RefreshClusterFromInstance(describeInstanceResult.Data)
	cluster.RefreshClusterFromInstanceListItem(item)
	return nil
}

func refreshTenant(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) error {
	tflog.Debug(ctx, "start to refresh tenant")
	if tenant.Id.IsNull() || tenant.Id.IsUnknown() || tenant.Id.ValueString() == "" {
		return errors.New("Tenant id is invalid")
	}

	if tenant.ClusterId.IsNull() || tenant.ClusterId.IsUnknown() || tenant.ClusterId.ValueString() == "" {
		return errors.New("Cluster id is invalid")
	}
	result, response, err := client.MultiCloudOpenAPI.DescribeTenant(ctx, tenant.ClusterId.ValueString(), tenant.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err == nil && response != nil && response.StatusCode == http.StatusOK && result != nil && !result.GetSuccess() && result.GetErrorCode() == "Tenant.NotExists" {
		return fmt.Errorf("tenant %s: %w", tenant.Id.ValueString(), errResourceNotFound)
	}
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		if result != nil {
			detailErr := fmt.Errorf("failed to describe tenant: %v, code: %s, message: %s", err, result.GetErrorCode(), result.GetErrorMessage())
			if err == nil && response != nil && response.StatusCode == http.StatusOK && !result.GetSuccess() && result.GetErrorCode() == "IllegalOperation.Resource" {
				return &tenantDetailResourceStateError{detailErr}
			}
			return detailErr
		}
		return fmt.Errorf("failed to describe tenant: %v", err)
	}
	if result.Data.GetTenantId() != tenant.Id.ValueString() {
		return errors.New("describe tenant returned a different or missing tenant id")
	}
	tenant.RefreshTenant(result.Data)
	// The detail API's enableProxyPublicLink is a capability, not an address.
	// A known value on a resource is its creation input; a data source or an
	// import with no prior value needs to discover the current address instead.
	if tenant.EnablePublicAddress.IsNull() || tenant.EnablePublicAddress.IsUnknown() {
		enabled, err := tenantHasPublicAddress(ctx, tenant, client)
		if err != nil {
			return err
		}
		tenant.EnablePublicAddress = types.BoolValue(enabled)
	}
	tflog.Debug(ctx, fmt.Sprintf("Generated tenant, %v", tenant))
	tflog.Debug(ctx, "finished to refresh tenant")
	return nil
}

func refreshTenantSecurityGroup(ctx context.Context, tenantSecurityGroup *model.TenantSecurityGroup, client *obcloudsdk.APIClient) error {
	if tenantSecurityGroup.TenantId.IsNull() || tenantSecurityGroup.TenantId.IsUnknown() {
		return errors.New("Tenant id is invalid")
	}

	if tenantSecurityGroup.ClusterId.IsNull() || tenantSecurityGroup.ClusterId.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	if tenantSecurityGroup.Name.IsNull() || tenantSecurityGroup.Name.IsUnknown() {
		return errors.New("Security group name invalid")
	}
	result, response, err := client.MultiCloudOpenAPI.DescribeTenantSecurityIpGroups(ctx, tenantSecurityGroup.ClusterId.ValueString(), tenantSecurityGroup.TenantId.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil {
		return fmt.Errorf("describe tenant security groups: %w", err)
	}
	if response.StatusCode != http.StatusOK || !result.GetSuccess() {
		return fmt.Errorf("describe tenant security groups: %s (%s)", response.Status, result.GetErrorCode())
	}
	page := result.Data
	if page == nil || page.DataList == nil && (page.Total == nil || page.GetTotal() != 0) {
		return errors.New("describe tenant security groups returned no list")
	}
	for _, securityGroup := range page.DataList {
		if securityGroup.GetSecurityIpGroupName() == "" {
			return errors.New("describe tenant security groups returned an unnamed group")
		}
		if securityGroup.GetSecurityIpGroupName() == tenantSecurityGroup.Name.ValueString() {
			if securityGroup.SecurityIps == nil {
				return errors.New("describe tenant security group returned no securityIps")
			}
			cidrs := make([]types.String, 0)
			for _, cidr := range strings.Split(securityGroup.GetSecurityIps(), ",") {
				if ip := strings.TrimSpace(cidr); ip != "" {
					cidrs = append(cidrs, types.StringValue(ip))
				}
			}
			if !sameSecurityGroupIPs(tenantSecurityGroup.Cidrs, cidrs) {
				tenantSecurityGroup.Cidrs = cidrs
			}
			return nil
		}
	}
	if page.GetHasMore() || page.Total != nil && page.GetTotal() != int64(len(page.DataList)) {
		return errors.New("describe tenant security groups returned an incomplete list")
	}
	return fmt.Errorf("security group %s: %w", tenantSecurityGroup.Name.ValueString(), errResourceNotFound)
}

func refreshTenantUser(ctx context.Context, tenantUser *model.TenantUser, client *obcloudsdk.APIClient) error {
	if tenantUser.TenantId.IsNull() || tenantUser.TenantId.IsUnknown() {
		return errors.New("Tenant id is invalid")
	}
	if tenantUser.ClusterId.IsNull() || tenantUser.ClusterId.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	if tenantUser.Name.IsNull() || tenantUser.Name.IsUnknown() {
		return errors.New("User name is invalid")
	}
	result, response, err := client.MultiCloudOpenAPI.DescribeTenantUsers(ctx, tenantUser.ClusterId.ValueString(), tenantUser.TenantId.ValueString()).RequestId("").UserName(tenantUser.Name.ValueString()).Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
		return errors.New("Failed to call api to refresh tenant user")
	}
	tenantUsers := result.GetData().DataList
	for idx, user := range tenantUsers {
		if user.GetUserName() == tenantUser.Name.ValueString() {
			tenantUser.Refresh(&tenantUsers[idx])
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Tenant user %s does not exist", tenantUser.Name.ValueString()))

}

func refreshDatabase(ctx context.Context, database *model.Database, client *obcloudsdk.APIClient) error {
	if database.TenantId.IsNull() || database.TenantId.IsUnknown() {
		return errors.New("Tenant id is invalid")
	}

	if database.ClusterId.IsNull() || database.ClusterId.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	if database.Name.IsNull() || database.Name.IsUnknown() {
		return errors.New("Database name is invalid")
	}
	// only one result will be returned since database name is passed
	result, response, err := client.MultiCloudOpenAPI.DescribeDatabases(ctx, database.ClusterId.ValueString(), database.TenantId.ValueString()).DatabaseName(database.Name.ValueString()).PageSize(10).PageNumber(1).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
		return errors.New("Failed to call api to refresh database")
	}

	databases := result.GetData().DataList
	for idx, db := range databases {
		if db.GetDatabaseName() == database.Name.ValueString() {
			database.Refresh(&databases[idx])
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Database %s does not exist", database.Name.ValueString()))
}

func refreshTenantConnection(ctx context.Context, tenantConnection *model.TenantConnection, client *obcloudsdk.APIClient) error {
	if tenantConnection.Id.IsNull() || tenantConnection.Id.IsUnknown() || tenantConnection.Id.ValueString() == "" {
		return errors.New("Tenant connection address id is invalid")
	}

	if tenantConnection.ClusterId.IsNull() || tenantConnection.ClusterId.IsUnknown() || tenantConnection.ClusterId.ValueString() == "" {
		return errors.New("Cluster id is invalid")
	}
	if tenantConnection.TenantId.IsNull() || tenantConnection.TenantId.IsUnknown() || tenantConnection.TenantId.ValueString() == "" {
		return errors.New("Tenant id is invalid")
	}

	var result *obcloudsdk.OBCloudResultListTenantConnectionDTO
	var response *http.Response
	var err error
	switch strings.ToUpper(tenantConnection.Type.ValueString()) {
	case model.ConnectionTypeIntranet:
		result, response, err = client.MultiCloudOpenAPI.DescribeTenantPrivateLink(ctx, tenantConnection.ClusterId.ValueString(), tenantConnection.TenantId.ValueString()).AddressId(tenantConnection.Id.ValueString()).RequestId("").Execute()
	case model.ConnectionTypeInternet:
		result, response, err = client.MultiCloudOpenAPI.DescribeTenantAddress(ctx, tenantConnection.ClusterId.ValueString(), tenantConnection.TenantId.ValueString()).AddressId(tenantConnection.Id.ValueString()).RequestId("").Execute()
	default:
		return fmt.Errorf("unsupported tenant connection type %q", tenantConnection.Type.ValueString())
	}
	if err != nil || response != nil {
		api.LogAPIResult(ctx, result, response, err)
	}
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil {
		return fmt.Errorf("failed to query tenant connection %q: response: %v, error: %v", tenantConnection.Id.ValueString(), response, err)
	}
	if !result.GetSuccess() {
		if result.GetErrorCode() == "Tenant.Address.NotExist" {
			return fmt.Errorf("tenant connection %q: %w", tenantConnection.Id.ValueString(), errResourceNotFound)
		}
		return fmt.Errorf("failed to query tenant connection %q: %s: %s", tenantConnection.Id.ValueString(), result.GetErrorCode(), result.GetErrorMessage())
	}
	connections := result.GetData()
	if connections == nil {
		return errors.New("tenant connection query returned no data")
	}
	if len(connections) == 0 {
		return fmt.Errorf("tenant connection %q: %w", tenantConnection.Id.ValueString(), errResourceNotFound)
	}
	if len(connections) == 1 && connections[0].GetAddressId() == "" && strings.EqualFold(tenantConnection.Type.ValueString(), model.ConnectionTypeInternet) {
		isNull, err := connectionResponseHasNullEntry(response)
		if err != nil {
			return err
		}
		if isNull {
			if err := confirmPublicConnectionAbsent(ctx, tenantConnection, client); err != nil {
				return err
			}
			return fmt.Errorf("tenant connection %q: %w", tenantConnection.Id.ValueString(), errResourceNotFound)
		}
	}
	if len(connections) != 1 || connections[0].GetAddressId() != tenantConnection.Id.ValueString() {
		return fmt.Errorf("tenant connection query did not return exactly address %q", tenantConnection.Id.ValueString())
	}
	address := &connections[0]
	if id := address.GetTenantId(); id != "" && id != tenantConnection.TenantId.ValueString() {
		return fmt.Errorf("address %q belongs to a different tenant", tenantConnection.Id.ValueString())
	}
	if kind := address.GetNetworkType(); kind != "" && !strings.EqualFold(kind, tenantConnection.Type.ValueString()) {
		return fmt.Errorf("address %q has unexpected network type %q", tenantConnection.Id.ValueString(), kind)
	}
	tenantConnection.Refresh(address)
	return nil
}
