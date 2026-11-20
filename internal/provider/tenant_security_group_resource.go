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
	"errors"
	"fmt"
	"net/http"
	"net/netip"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant id",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the security group",
				Required:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
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

	body, err := securityGroupRequest(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Invalid CIDRs", err.Error())
		return
	}

	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateTenantSecurityIpGroup(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).RequestId("").Body(*body).Execute()
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
	state.Id = types.StringValue(generateSecurityGroupId(plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()))
	timeout := operationTimeout(EnvResourceCreateTimeoutSeconds, DefaultResourceCreateTimeout)
	err = retry.RetryContext(ctx, timeout, waitSecurityGroupMatch(ctx, &state, client, plan.Cidrs))
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if err != nil {
		resp.Diagnostics.AddError("Waiting for tenant security group creation failed", err.Error())
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
	if errors.Is(err, errResourceNotFound) {
		resp.State.RemoveResource(ctx)
		return
	}
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

// The API accepts single IPv4 addresses without the /32 suffix.
func securityGroupIPs(cidrs []types.String) []string {
	ips := make([]string, len(cidrs))
	for i, cidr := range cidrs {
		value := cidr.ValueString()
		if prefix, err := netip.ParsePrefix(value); err == nil && prefix.Addr().Is4() && prefix.Bits() == 32 {
			value = prefix.Addr().String()
		}
		ips[i] = value
	}
	return ips
}

func sameSecurityGroupIPs(a, b []types.String) bool {
	left, right := securityGroupIPs(a), securityGroupIPs(b)
	slices.Sort(left)
	slices.Sort(right)
	return slices.Equal(left, right)
}

func securityGroupRequest(plan *model.TenantSecurityGroup) (*obcloudsdk.McModifyTenantSecurityIpGroupRequest, error) {
	ips, err := json.Marshal(securityGroupIPs(plan.Cidrs))
	if err != nil {
		return nil, err
	}
	body := obcloudsdk.NewMcModifyTenantSecurityIpGroupRequestWithDefaults()
	body.SetSecurityIpGroupName(plan.Name.ValueString())
	body.SetSecurityIps(string(ips))
	return body, nil
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

	if !sameSecurityGroupIPs(state.Cidrs, plan.Cidrs) {
		body, err := securityGroupRequest(&plan)
		if err != nil {
			resp.Diagnostics.AddError("Invalid CIDRs", err.Error())
			return
		}
		client := r.provider.clientGenerator.GetClient()
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantSecurityIpGroup(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
			resp.Diagnostics.AddError("Modify tenant security group CIDRs failed", fmt.Sprintf("Error: %v, response: %v", err, response))
			return
		}
		timeout := operationTimeout(EnvResourceModifyTimeoutSeconds, DefaultResourceModifyTimeout)
		if err := retry.RetryContext(ctx, timeout, waitSecurityGroupMatch(ctx, &state, client, plan.Cidrs)); err != nil {
			resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			resp.Diagnostics.AddError("Waiting for tenant security group update failed", err.Error())
			return
		}
	}
	state.Cidrs = plan.Cidrs
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
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
	if err != nil || response.StatusCode != http.StatusOK || !result.GetSuccess() {
		resp.Diagnostics.AddError(
			"Failed to delete tenant security group",
			fmt.Sprintf("Unexpected error while delete tenant security group: %v, response: %v", err, response))
		return
	}

	timeout := operationTimeout(EnvResourceDeleteTimeoutSeconds, DefaultResourceModifyTimeout)
	if err := waitForRemoval(ctx, timeout, func(ctx context.Context) error {
		return refreshTenantSecurityGroup(ctx, &tenantSecurityGroup, client)
	}); err != nil {
		resp.Diagnostics.AddError("Waiting for tenant security group deletion failed", err.Error())
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
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), generateSecurityGroupId(parts[0], parts[1], parts[2]))...)
}

func waitSecurityGroupMatch(ctx context.Context, state *model.TenantSecurityGroup, client *obcloudsdk.APIClient, desired []types.String) retry.RetryFunc {
	return func() *retry.RetryError {
		err := refreshTenantSecurityGroup(ctx, state, client)
		if err != nil {
			if errors.Is(err, errResourceNotFound) {
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		if !sameSecurityGroupIPs(state.Cidrs, desired) {
			return retry.RetryableError(errors.New("tenant security group CIDRs have not converged"))
		}
		state.Cidrs = desired
		return nil
	}
}
