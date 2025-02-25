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
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &tenantUserResource{}
	_ resource.ResourceWithConfigure   = &tenantUserResource{}
	_ resource.ResourceWithImportState = &tenantUserResource{}
)

type tenantUserResource struct {
	provider *oceanbaseProvider
}

func NewTenantUserResource() resource.Resource {
	return &tenantUserResource{}
}

func (r *tenantUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant User Resource",
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
				MarkdownDescription: "Name of the user",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password of the user",
				Required:            true,
				Sensitive:           true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the user",
				Required:            true,
			},
			"privileges": schema.ListNestedAttribute{
				MarkdownDescription: "Privileges of the user",
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"database": schema.StringAttribute{
							MarkdownDescription: "Database name",
							Required:            true,
						},
						"table": schema.StringAttribute{
							MarkdownDescription: "Table name",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString(""),
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "User role",
							Required:            true,
						},
						"privileges": schema.StringAttribute{
							MarkdownDescription: "Privileges",
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString(""),
						},
						"with_grant": schema.BoolAttribute{
							MarkdownDescription: "With grant privilege",
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
					},
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the user",
				Optional:            true,
				Computed:            true,
			},
			"is_locked": schema.BoolAttribute{
				MarkdownDescription: "Whether the user is locked",
				Optional:            true,
				Computed:            true,
			},
			"encryption_type": schema.StringAttribute{
				MarkdownDescription: "Encryption type of the user",
				Optional:            true,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the user",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the user",
				Computed:            true,
			},
		},
	}
}

func (r *tenantUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_user"
}

func (r *tenantUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func generateUserId(clusterId, tenantId, name string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s:%s", clusterId, tenantId, name)))
}

func (r *tenantUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create tenant user is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.TenantUser

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	createTenantUserBody := obcloudsdk.NewCreateUserParamDoWithDefaults()
	createTenantUserBody.SetUserName(plan.Name.ValueString())
	createTenantUserBody.SetInstanceId(plan.ClusterId.ValueString())
	createTenantUserBody.SetTenantId(plan.TenantId.ValueString())
	createTenantUserBody.SetPassword(plan.Password.ValueString())
	createTenantUserBody.SetUserType(plan.Type.ValueString())
	createTenantUserBody.SetEncryptionType(plan.EncryptionType.ValueString())
	createTenantUserBody.SetDescription(plan.Description.ValueString())
	// TODO: set roles when user type is normal
	privilegeStr, err := plan.GetPrivilegesStr()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting tenant user's privileges",
			fmt.Sprintf("Could not get tenant user privilege configuration err: %v", err),
		)
		return
	}
	createTenantUserBody.SetRoles(privilegeStr)
	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateTenantUser(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).RequestId("").Body(*createTenantUserBody).Execute()
	api.LogAPIResult(ctx, result, response, err)

	if err != nil || response.StatusCode != http.StatusOK || !(result.GetSuccess()) {
		resp.Diagnostics.AddError(
			"Error creating tenant user",
			fmt.Sprintf("Could not create tenant user err: %v, response: %v", err, response),
		)
		return
	}

	state := plan
	state.Id = types.StringValue(generateUserId(plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()))
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error after create the tenant", "Set the state error")
		return
	}
}

func (r *tenantUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant user is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenantUser model.TenantUser
	diags := req.State.Get(ctx, &tenantUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := refreshTenantUser(ctx, &tenantUser, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant user info failed",
			fmt.Sprintf("Failed to refresh tenant user info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getTenantUserOperation(ctx context.Context, state, plan *model.TenantUser) (string, error) {
	modifiedSpecs := make([]string, 0)
	if !(plan.Description.IsUnknown() || plan.Description.IsNull()) && !plan.Description.Equal(state.Description) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantUserDescription)
	}
	if !(plan.Password.IsUnknown() || plan.Password.IsNull()) && !plan.Password.Equal(state.Password) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantUserPassword)
	}
	if !(plan.IsLocked.IsUnknown() || plan.IsLocked.IsNull()) && !(plan.IsLocked.Equal(state.IsLocked)) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantUserLockedStatus)
	}
	// compare privileges
	statePrivileges := make(map[string]struct{})
	planPrivileges := make(map[string]struct{})
	for _, privilege := range state.Privileges.Elements() {
		statePrivileges[privilege.String()] = struct{}{}
	}
	for _, privilege := range plan.Privileges.Elements() {
		planPrivileges[privilege.String()] = struct{}{}
	}

	matched := true
	for privilege, _ := range statePrivileges {
		_, found := planPrivileges[privilege]
		if !found {
			matched = false
			break
		}
	}
	for privilege, _ := range planPrivileges {
		_, found := statePrivileges[privilege]
		if !found {
			matched = false
			break
		}
	}
	if !matched {
		modifiedSpecs = append(modifiedSpecs, SpecTenantUserPrivileges)
	}

	if len(modifiedSpecs) > 1 {
		return OperationUnsupported, errors.New(fmt.Sprintf("Modify %s at once is not supported", strings.Join(modifiedSpecs, ", ")))
	}

	if len(modifiedSpecs) == 1 {
		switch modifiedSpecs[0] {
		case SpecTenantUserDescription:
			return OperationModifyTenantUserDescription, nil
		case SpecTenantUserPassword:
			return OperationModifyTenantUserPassword, nil
		case SpecTenantUserPrivileges:
			return OperationModifyTenantUserPrivileges, nil
		case SpecTenantUserLockedStatus:
			return OperationModifyTenantUserLockedStatus, nil
		default:
			return OperationUnsupported, errors.New("Unsupported operation")
		}
	}
	return OperationNothing, nil
}

func (r *tenantUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update tenant user is called")

	var plan model.TenantUser
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state model.TenantUser
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := getTenantUserOperation(ctx, &state, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update is not supported",
			fmt.Sprintf("Got error when get operation err: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Got operation: %s", operation))

	client := r.provider.clientGenerator.GetClient()
	switch operation {
	case OperationModifyTenantUserDescription:
		tflog.Info(ctx, "Modify tenant user description")
		body := obcloudsdk.NewModifyTenantUserDescriptionParamDoWithDefaults()
		body.SetDescription(plan.Description.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantUserDescription(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify tenant user description failed")
			resp.Diagnostics.AddError(
				"Modify tenant user description failed",
				fmt.Sprintf("Got error when update tenant user description err: %v, response: %v", err, response))
			return
		}
	case OperationModifyTenantUserPassword:
		tflog.Info(ctx, "Modify tenant user password")
		body := obcloudsdk.NewModifyUserPasswordParamDoWithDefaults()
		body.SetUserPassword(plan.Password.ValueString())
		body.SetUserName(plan.Name.ValueString())
		body.SetTenantId(plan.TenantId.ValueString())
		body.SetInstanceId(plan.ClusterId.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantUserPassword(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify tenant user password failed")
			resp.Diagnostics.AddError(
				"Modify tenant user password failed",
				fmt.Sprintf("Got error when update tenant user password err: %v, response: %v", err, response))
			return
		}
		// set password manually, the api won't return password
		state.Password = plan.Password
	case OperationModifyTenantUserLockedStatus:
		tflog.Info(ctx, "Modify tenant user locked status")
		body := obcloudsdk.NewModifyUserStatusParamDoWithDefaults()
		body.SetInstanceId(plan.ClusterId.ValueString())
		body.SetTenantId(plan.TenantId.ValueString())
		body.SetUserName(plan.Name.ValueString())
		targetStatus := model.TenantUserStatusOnline
		if plan.IsLocked.ValueBool() {
			targetStatus = model.TenantUserStatusLocked
		}
		body.SetUserStatus(targetStatus)
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantUserStatus(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify tenant user locked status failed")
			resp.Diagnostics.AddError(
				"Modify tenant user status failed",
				fmt.Sprintf("Got error when update tenant user status err: %v, response: %v", err, response))
			return
		}
	case OperationModifyTenantUserPrivileges:
		tflog.Info(ctx, "Modify tenant user privileges")
		body := obcloudsdk.NewModifyUserRolesParamDoWithDefaults()
		body.SetModifyType("UPDATE")
		privilegeStr, err := plan.GetPrivilegesStr()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting tenant user's privileges",
				fmt.Sprintf("Could not get tenant user privilege configuration err: %v", err),
			)
			return
		}
		body.SetUserRole(privilegeStr)
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantUserRole(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()).Body(*body).Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify tenant user privileges failed")
			resp.Diagnostics.AddError(
				"Modify tenant user privileges failed",
				fmt.Sprintf("Got error when update tenant user privileges err: %v, response: %v", err, response))
			return
		}
		state.Privileges = plan.Privileges
	case OperationNothing:
		tflog.Info(ctx, "Nothing to do")
	default:
		tflog.Info(ctx, "Unsupported operation")
	}

	err = refreshTenantUser(ctx, &state, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant user info failed",
			fmt.Sprintf("Failed to refresh tenant user info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *tenantUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete tenant user is called")
	var tenantUser model.TenantUser
	diags := req.State.Get(ctx, &tenantUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if tenantUser.TenantId.IsNull() || tenantUser.TenantId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant id can't be null",
			"The tenant_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantUser.ClusterId.IsNull() || tenantUser.ClusterId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant's cluster id can't be null",
			"The cluster_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantUser.Name.IsNull() || tenantUser.Name.IsUnknown() {
		resp.Diagnostics.AddError(
			"User name can't be null",
			"The name field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()
	deleteUserBody := obcloudsdk.NewDeleteTenantUserParamDoWithDefaults()
	deleteUserBody.SetInstanceId(tenantUser.ClusterId.ValueString())
	deleteUserBody.SetTenantId(tenantUser.TenantId.ValueString())
	deleteUserBody.SetUsers(fmt.Sprintf("[\"%s\"]", tenantUser.Name.ValueString()))

	result, response, err := client.MultiCloudOpenAPI.DeleteTenantUsers(ctx, tenantUser.ClusterId.ValueString(), tenantUser.TenantId.ValueString()).Body(*deleteUserBody).Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
		resp.Diagnostics.AddError(
			"Failed to delete tenant user",
			fmt.Sprintf("Unexpected error while delete tenant user: %v, response: %v", err, response))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *tenantUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import tenant user is called")
	parts := strings.Split(req.ID, ".")
	if len(parts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid id format",
			fmt.Sprintf("Unexpected error while parsing resource id %s, the expected format is <cluster_id>.<tenant_id>.<name>", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tenant_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parts[2])...)
}
