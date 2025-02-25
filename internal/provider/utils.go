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
	if cluster.Id.IsNull() || cluster.Id.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	describeInstanceResult, describeInstanceResponse, describeInstanceErr := client.MultiCloudOpenAPI.DescribeInstance(ctx, cluster.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, describeInstanceResult, describeInstanceResponse, describeInstanceErr)
	describeInstancesResult, describeInstancesResponse, describeInstancesErr := client.MultiCloudOpenAPI.DescribeInstances(ctx).InstanceId(cluster.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, describeInstancesResult, describeInstancesResponse, describeInstancesErr)

	describeInstanceSuccess := describeInstanceErr == nil && describeInstanceResponse.StatusCode == http.StatusOK && describeInstanceResult.GetSuccess()
	describeInstancesSuccess := describeInstancesErr == nil && describeInstancesResponse.StatusCode == http.StatusOK && describeInstancesResult.GetSuccess()

	if !describeInstanceSuccess && !describeInstancesSuccess {
		return errors.New(fmt.Sprintf("Failed to call api to refresh cluster, describeInstanceErr: %v, describeInstancesErr: %v, describeInstanceResponse: %v, describeInstancesResponse: %v", describeInstanceErr, describeInstancesErr, describeInstanceResponse, describeInstancesResponse))
	}

	if describeInstanceSuccess {
		if !strings.HasSuffix(strings.ToUpper(describeInstanceResult.Data.GetInstanceType()), InstanceTypeCluster) {
			return errors.New("Instance type mismatch")
		}
		cluster.RefreshClusterFromInstance(describeInstanceResult.Data)
	}
	if describeInstancesSuccess {
		if *describeInstancesResult.GetData().Total != 1 {
			return errors.New("Multiple instances returned")
		}
		if !strings.HasSuffix(strings.ToUpper(describeInstancesResult.GetData().DataList[0].GetInstanceType()), InstanceTypeCluster) {
			return errors.New("Instance type mismatch")
		}
		cluster.RefreshClusterFromInstanceListItem(&describeInstancesResult.GetData().DataList[0])
	}
	return nil
}

func refreshTenant(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) error {
	tflog.Debug(ctx, "start to refresh tenant")
	if tenant.Id.IsNull() || tenant.Id.IsUnknown() {
		return errors.New("Tenant id is invalid")
	}

	if tenant.ClusterId.IsNull() || tenant.ClusterId.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	result, response, err := client.MultiCloudOpenAPI.DescribeTenant(ctx, tenant.ClusterId.ValueString(), tenant.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
		return errors.New("Failed to call api to refresh tenant")
	}
	tenant.RefreshTenant(result.Data)
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
	if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
		return errors.New("Failed to call api to refresh tenant security group")
	}
	tenantSecurityGroups := result.GetData().DataList
	securityGroupFound := false
	for _, securityGroup := range tenantSecurityGroups {
		if securityGroup.GetSecurityIpGroupName() == tenantSecurityGroup.Name.ValueString() {
			cidrs := make([]types.String, 0)
			for _, cidr := range strings.Split(securityGroup.GetSecurityIps(), ",") {
				cidrs = append(cidrs, types.StringValue(cidr))
			}
			tenantSecurityGroup.Cidrs = cidrs
			securityGroupFound = true
			break
		}
	}
	if !securityGroupFound {
		return errors.New(fmt.Sprintf("Security group %s not found", tenantSecurityGroup.Name.ValueString()))
	}
	return nil
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
	if tenantConnection.Id.IsNull() || tenantConnection.Id.IsUnknown() {
		return errors.New("Tenant connection address id is invalid")
	}

	if tenantConnection.ClusterId.IsNull() || tenantConnection.ClusterId.IsUnknown() {
		return errors.New("Cluster id is invalid")
	}
	if tenantConnection.TenantId.IsNull() || tenantConnection.TenantId.IsUnknown() {
		return errors.New("Tenant id is invalid")
	}

	switch strings.ToUpper(tenantConnection.Type.ValueString()) {
	case model.ConnectionTypeIntranet:
		result, response, err := client.MultiCloudOpenAPI.DescribeTenantPrivateLink(ctx, tenantConnection.ClusterId.ValueString(), tenantConnection.TenantId.ValueString()).AddressId(tenantConnection.Id.ValueString()).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() || len(result.GetData()) != 1 {
			return errors.New("Failed to call api to refresh tenant connection")
		}
		tenantConnectionList := result.GetData()
		tenantConnection.Refresh(&tenantConnectionList[0])
	case model.ConnectionTypeInternet:
		result, response, err := client.MultiCloudOpenAPI.DescribeTenantPublicAddress(ctx, tenantConnection.ClusterId.ValueString(), tenantConnection.TenantId.ValueString()).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
			return errors.New("Failed to call api to refresh tenant connection")
		}
		connectionData := result.GetData()
		tenantConnection.Refresh(&connectionData)
	}
	return nil
}
