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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &tenantSecurityGroupResource{}
	_ resource.ResourceWithConfigure   = &tenantSecurityGroupResource{}
	_ resource.ResourceWithImportState = &tenantSecurityGroupResource{}
)

type tenantSecurityGroupResource struct {
	provider *oceanbaseProvider
}

func NewTenantSecurityGroupResource() resource.Resource {
	return &tenantSecurityGroupResource{}
}

func (r *tenantSecurityGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant Security Group Resource",
		Attributes: map[string]schema.Attribute{
			// required fields to create tenant
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster id of the tenant",
				Required:            true,
			},
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant id",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the security group",
				Required:            true,
			},
			"cidrs": schema.ListAttribute{
				MarkdownDescription: "Security group cidrs",
				Required:            true,
				ElementType:         types.StringType,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the security group",
				Computed:            true,
			},
		},
	}
}

func (r *tenantSecurityGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_security_group"
}

func (r *tenantSecurityGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func generateSecurityGroupId(clusterId, tenantId, name string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s:%s", clusterId, tenantId, name)))
}

func (r *tenantSecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create tenant security group is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.TenantSecurityGroup

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	ips := make([]string, len(plan.Cidrs))
	for idx, ip := range plan.Cidrs {
		ips[idx] = ip.ValueString()
	}
	createTenantSecurityGroupBody := obcloudsdk.NewMcModifyTenantSecurityIpGroupRequestWithDefaults()
	createTenantSecurityGroupBody.SetSecurityIpGroupName(plan.Name.ValueString())
	ipsBytes, err := json.Marshal(ips)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid cidrs config",
			fmt.Sprintf("Could not marshal cidrs err: %v", err),
		)
		return
	}
	ipsStr := string(ipsBytes)
	createTenantSecurityGroupBody.SetSecurityIps(ipsStr)

	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateTenantSecurityIpGroup(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).RequestId("").Body(*createTenantSecurityGroupBody).Execute()
	api.LogAPIResult(ctx, result, response, err)

	// TODO: add retry policy according to result and response
	if err != nil || response.StatusCode != http.StatusOK || !(result.GetSuccess()) {
		resp.Diagnostics.AddError(
			"Error creating tenant security group",
			fmt.Sprintf("Could not create tenant security group err: %v, response: %v", err, response),
		)
		return
	}

	state := plan
	timeoutDuration := DefaultResourceCreateTimeout
	envTimeoutSeconds := os.Getenv(EnvResourceCreateTimeoutSeconds)
	if envTimeoutSeconds != "" {
		timeoutSeconds, err := strconv.ParseInt(envTimeoutSeconds, 10, 64)
		if err == nil {
			timeoutDuration = time.Duration(timeoutSeconds) * time.Second
		}
	}
	err = retry.RetryContext(ctx, timeoutDuration,
		waitSecurityGroupMatch(ctx, &state, client, cidrsMatchedFuncGenerator(&plan)))
	// Just add an error log here but consider the resource is created successfully
	// state is already refreshed while waiting, no need to refresh again here
	if err != nil {
		tflog.Error(ctx, "Wait for tenant security match timed out, just save current state")
	}

	state.Id = types.StringValue(generateSecurityGroupId(plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()))
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error after create the tenant", "Set the state error")
		return
	}
}

func (r *tenantSecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant security group is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenantSecurityGroup model.TenantSecurityGroup
	diags := req.State.Get(ctx, &tenantSecurityGroup)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := refreshTenantSecurityGroup(ctx, &tenantSecurityGroup, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant security group info failed",
			fmt.Sprintf("Failed to refresh tenant security group info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantSecurityGroup)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getTenantSecurityGroupOperation(ctx context.Context, state, plan *model.TenantSecurityGroup) (string, error) {
	stateCidrs := make(map[string]struct{})
	planCidrs := make(map[string]struct{})
	for _, stateCidr := range state.Cidrs {
		stateCidrs[stateCidr.ValueString()] = struct{}{}
	}
	for _, planCidr := range plan.Cidrs {
		planCidrs[planCidr.ValueString()] = struct{}{}
	}

	matched := true
	for cidr, _ := range stateCidrs {
		_, found := planCidrs[cidr]
		if !found {
			matched = false
			break
		}
	}
	for cidr, _ := range planCidrs {
		_, found := stateCidrs[cidr]
		if !found {
			matched = false
			break
		}
	}
	if !matched {
		return OperationModifyTenantSecurityGroupCidrs, nil
	}
	return OperationNothing, nil
}

func (r *tenantSecurityGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update tenant security group is called")

	var plan model.TenantSecurityGroup
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state model.TenantSecurityGroup
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := getTenantSecurityGroupOperation(ctx, &state, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update is not supported",
			fmt.Sprintf("Got error when get operation err: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Got operation: %s", operation))

	client := r.provider.clientGenerator.GetClient()
	switch operation {
	case OperationModifyTenantSecurityGroupCidrs:
		tflog.Info(ctx, "Modify tenant security group cidrs")
		body := obcloudsdk.NewMcModifyTenantSecurityIpGroupRequestWithDefaults()
		planCidrs := make([]string, len(plan.Cidrs))
		for idx, cidr := range plan.Cidrs {
			planCidrs[idx] = cidr.ValueString()
		}
		cidrBytes, err := json.Marshal(planCidrs)
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid cidrs config",
				fmt.Sprintf("Could not marshal cidrs err: %v", err),
			)
			return
		}
		cidrStr := string(cidrBytes)
		body.SetSecurityIps(cidrStr)
		body.SetSecurityIpGroupName(plan.Name.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantSecurityIpGroup(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify tenant security group cidrs failed")
			resp.Diagnostics.AddError(
				"Modify tenant security group cidrs failed",
				fmt.Sprintf("Got error when update tenant security group cidrs err: %v, response: %v", err, response))
			return
		}

		timeoutDuration := DefaultResourceCreateTimeout
		envTimeoutSeconds := os.Getenv(EnvResourceCreateTimeoutSeconds)
		if envTimeoutSeconds != "" {
			timeoutSeconds, err := strconv.ParseInt(envTimeoutSeconds, 10, 64)
			if err == nil {
				timeoutDuration = time.Duration(timeoutSeconds) * time.Second
			}
		}
		err = retry.RetryContext(ctx, timeoutDuration,
			waitSecurityGroupMatch(ctx, &state, client, cidrsMatchedFuncGenerator(&plan)))
		// Just add an error log here but consider the resource is created successfully
		// state is already refreshed while waiting, no need to refresh again here
		if err != nil {
			tflog.Error(ctx, "Wait for tenant security match timed out, just save current state")
		}
		state.Cidrs = plan.Cidrs

	case OperationNothing:
		tflog.Info(ctx, "Nothing to do")
	default:
		tflog.Info(ctx, "Unsupported operation")
	}

	err = refreshTenantSecurityGroup(ctx, &state, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant security group info failed",
			fmt.Sprintf("Failed to refresh tenant security group info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *tenantSecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete tenant security group is called")
	var tenantSecurityGroup model.TenantSecurityGroup
	diags := req.State.Get(ctx, &tenantSecurityGroup)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tenantSecurityGroup.TenantId.IsNull() || tenantSecurityGroup.TenantId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant id can't be null",
			"The tenant_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantSecurityGroup.ClusterId.IsNull() || tenantSecurityGroup.ClusterId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant's cluster id can't be null",
			"The cluster_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantSecurityGroup.Name.IsNull() || tenantSecurityGroup.Name.IsUnknown() {
		resp.Diagnostics.AddError(
			"Security group name can't be null",
			"The name field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()
	deleteSecurityGroupBody := obcloudsdk.NewMcDeleteTenantSecurityIpGroupRequestWithDefaults()
	deleteSecurityGroupBody.SetSecurityIpGroupName(tenantSecurityGroup.Name.ValueString())

	result, response, err := client.MultiCloudOpenAPI.DeleteTenantSecurityIpGroup(ctx, tenantSecurityGroup.ClusterId.ValueString(), tenantSecurityGroup.TenantId.ValueString()).RequestId("").Body(*deleteSecurityGroupBody).Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
		resp.Diagnostics.AddError(
			"Failed to delete tenant security group",
			fmt.Sprintf("Unexpected error while delete tenant security group: %v, response: %v", err, response))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *tenantSecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import tenant security group is called")
	parts := strings.Split(req.ID, ".")
	if len(parts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid id format",
			fmt.Sprintf("Unexpected error while parsing resource id %s, the expected format is <cluster_id>.<tenant_id>.<security_group_name>", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tenant_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parts[2])...)
}

func cidrsMatchedFuncGenerator(plan *model.TenantSecurityGroup) func(*model.TenantSecurityGroup) bool {
	return func(state *model.TenantSecurityGroup) bool {
		planCidrs := make([]string, 0)
		stateCidrs := make([]string, 0)
		for _, pc := range plan.Cidrs {
			planCidrs = append(planCidrs, pc.ValueString())
		}
		for _, cc := range state.Cidrs {
			stateCidrs = append(stateCidrs, cc.ValueString())
		}
		sort.Strings(planCidrs)
		sort.Strings(stateCidrs)
		return strings.Join(planCidrs, ",") == strings.Join(stateCidrs, ",")
	}
}

func waitSecurityGroupMatch(ctx context.Context, state *model.TenantSecurityGroup, client *obcloudsdk.APIClient, statusValidationFunction func(*model.TenantSecurityGroup) bool) retry.RetryFunc {
	return func() *retry.RetryError {
		err := refreshTenantSecurityGroup(ctx, state, client)
		if err != nil {
			return retry.RetryableError(fmt.Errorf("Failed to refresh tenant security group info, err: %v", err))
		}
		if !statusValidationFunction(state) {
			return retry.RetryableError(fmt.Errorf("Tenant security group is still not match, current status: %v", state))
		}
		return nil
	}
}
