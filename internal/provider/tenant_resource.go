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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                 = &tenantResource{}
	_ resource.ResourceWithConfigure    = &tenantResource{}
	_ resource.ResourceWithImportState  = &tenantResource{}
	_ resource.ResourceWithModifyPlan   = &tenantResource{}
	_ resource.ResourceWithUpgradeState = &tenantResource{}
)

type tenantResource struct {
	provider *oceanbaseProvider
}

func NewTenantResource() resource.Resource {
	return &tenantResource{}
}

// Version 0 is retained for decoding state written by older provider versions.
func tenantSchemaV0() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Tenant Resource",
		Attributes: map[string]schema.Attribute{
			// required fields to create tenant
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Id of cluster to create the tenant",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the tenant",
				Required:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Tenant mode",
				Required:            true,
			},
			"primary_zone": schema.StringAttribute{
				MarkdownDescription: "Primary zone of the tenant",
				Required:            true,
			},
			"charset": schema.StringAttribute{
				MarkdownDescription: "Charset of the tenant",
				Required:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "Timezone of the tenant",
				Required:            true,
			},
			"unit_num": schema.Int32Attribute{
				MarkdownDescription: "Unit num of the tenant",
				Required:            true,
			},
			"unit": schema.SingleNestedAttribute{
				Required:            true,
				MarkdownDescription: "Unit config of the tenant",
				Attributes: map[string]schema.Attribute{
					"cpu": schema.Int32Attribute{
						MarkdownDescription: "Cpu config of the unit",
						Required:            true,
					},
					"memory": schema.Int32Attribute{
						MarkdownDescription: "Memory config of the unit",
						Required:            true,
					},
				},
			},
			"enable_public_address": schema.BoolAttribute{
				MarkdownDescription: "Whether to create a public address with the tenant. This is a creation-only setting; manage later address changes with oceanbase_tenant_connection. Imported resources initially reflect the existing address.",
				Optional:            true,
				Computed:            true,
			},
			"compatible_mode": schema.StringAttribute{
				MarkdownDescription: "Compatible mode of the tenant",
				Optional:            true,
				Computed:            true,
			},
			"whitelist": schema.StringAttribute{
				MarkdownDescription: "Whitelist passed when creating the tenant. This setting cannot be updated; manage additional IP groups with oceanbase_tenant_security_group.",
				Optional:            true,
				Computed:            true,
			},

			// optional fields to create tenant
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the tenant",
				Optional:            true,
				Computed:            true,
			},
			"lower_case_table_names": schema.Int32Attribute{
				MarkdownDescription: "Configuration for table names",
				Optional:            true,
				Computed:            true,
			},

			// computed fields, known after creation
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the tenant",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"primary_region": schema.StringAttribute{
				MarkdownDescription: "Primary region of the tenant",
				Computed:            true,
			},
			"create_time": schema.StringAttribute{
				MarkdownDescription: "Create time of the tenant",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the tenant",
				Computed:            true,
			},
			"vpc": schema.StringAttribute{
				MarkdownDescription: "VPC where the tenant is created",
				Computed:            true,
			},
			"deploy_type": schema.StringAttribute{
				MarkdownDescription: "Deploy type of the the tenant",
				Computed:            true,
			},
			"daily_merge_time": schema.StringAttribute{
				MarkdownDescription: "Tenant daily merge time",
				Computed:            true,
			},
			"connections": schema.ListNestedAttribute{
				MarkdownDescription: "Tenant connections",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Connection type",
							Computed:            true,
						},
						"address_id": schema.StringAttribute{
							MarkdownDescription: "Address id",
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: "Address status",
							Computed:            true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"locality": schema.ListNestedAttribute{
				MarkdownDescription: "Tenant locality",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"zone_name": schema.StringAttribute{
							MarkdownDescription: "Zone name",
							Computed:            true,
						},
						"replica_type": schema.StringAttribute{
							MarkdownDescription: "Replica type",
							Computed:            true,
						},
						"region": schema.StringAttribute{
							MarkdownDescription: "Region",
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "Role",
							Computed:            true,
						},
					},
				},
			},
			"tags": schema.ListNestedAttribute{
				MarkdownDescription: "Tenant locality",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							MarkdownDescription: "Tag key",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Tag value",
							Computed:            true,
						},
					},
				},
			},
			"resource": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cpu": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Cpu resource info of the tenant",
						Attributes: map[string]schema.Attribute{
							"total_cpu": schema.Float64Attribute{
								MarkdownDescription: "Total cpu of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_cpu": schema.Float64Attribute{
								MarkdownDescription: "Used cpu of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"unit_cpu": schema.Float64Attribute{
								MarkdownDescription: "Unit config cpu of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
					"disk": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Disk resource info of the tenant",
						Attributes: map[string]schema.Attribute{
							"total_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Total disk size of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Used disk size of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
					"memory": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Memory resource info of the tenant",
						Attributes: map[string]schema.Attribute{
							"total_memory": schema.Float64Attribute{
								MarkdownDescription: "Total memory of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_memory": schema.Float64Attribute{
								MarkdownDescription: "Used memory of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"unit_memory": schema.Float64Attribute{
								MarkdownDescription: "Unit config memory of the tenant",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *tenantResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tenantSchemaV0()
	resp.Schema.Version = 1
}

func (r *tenantResource) UpgradeState(_ context.Context) map[int64]resource.StateUpgrader {
	prior := tenantSchemaV0()
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &prior,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Version 0 stored a capability flag here, not address presence.
				// Preserve all other state and let Read discover the real value.
				resp.State.Raw = req.State.Raw
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("enable_public_address"), types.BoolNull())...)
			},
		},
	}
}

func (r *tenantResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

func (r *tenantResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (r *tenantResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var state, plan model.Tenant
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := validateTenantCreationInputs(&state, &plan); err != nil {
		resp.Diagnostics.AddError("Tenant creation-only setting cannot be updated", err.Error())
	}
}

func validateTenantCreationInputs(state, plan *model.Tenant) error {
	if !plan.EnablePublicAddress.IsNull() && !plan.EnablePublicAddress.IsUnknown() && !plan.EnablePublicAddress.Equal(state.EnablePublicAddress) {
		return errors.New("enable_public_address is used only when creating a tenant. Manage subsequent public addresses with oceanbase_tenant_connection")
	}
	if !plan.Whitelist.IsNull() && !plan.Whitelist.IsUnknown() && !plan.Whitelist.Equal(state.Whitelist) {
		return errors.New("whitelist is used only when creating a tenant and cannot be updated through this resource. Use oceanbase_tenant_security_group to manage a separate IP group")
	}
	return nil
}

func (r *tenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create tenant is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.Tenant

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add checks
	// check instance id exists
	unitConfig, err := plan.GetUnitConfig()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unit config is invalid",
			fmt.Sprintf("Failed to parse unit config err: %v", err))
		return
	}

	createTenantBody := obcloudsdk.NewCreateTenantRequestWithDefaults()
	createTenantBody.SetInstanceId(plan.ClusterId.ValueString())
	createTenantBody.SetTenantName(plan.Name.ValueString())
	createTenantBody.SetTenantMode(plan.Mode.ValueString())
	createTenantBody.SetPrimaryZone(plan.PrimaryZone.ValueString())
	createTenantBody.SetCharset(plan.Charset.ValueString())
	createTenantBody.SetDescription(plan.Description.ValueString())
	createTenantBody.SetWhitelist(plan.Whitelist.ValueString())
	createTenantBody.SetTimeZone(plan.Timezone.ValueString())
	createTenantBody.SetCpu(unitConfig.Cpu.ValueInt32())
	createTenantBody.SetMemory(unitConfig.Memory.ValueInt32())
	createTenantBody.SetUnitNum(plan.UnitNum.ValueInt32())
	// which value to set
	// createTenantBody.SetEnvironment("dev")
	// createTenantBody.SetTenantCompatibilityModeEnum(unitConfig.CompatibleMode.ValueString())
	createTenantBody.SetEnablePublicLink(plan.EnablePublicAddress.ValueBool())
	createTenantBody.SetCreateParams(map[string]string{"lower_case_table_names": fmt.Sprintf("%d", plan.LowerCaseTableNames.ValueInt32())})

	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateTenant(ctx, plan.ClusterId.ValueString()).RequestId("").Body(*createTenantBody).Execute()
	api.LogAPIResult(ctx, result, response, err)

	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		resp.Diagnostics.AddError(
			"Error creating tenant",
			fmt.Sprintf("Could not create tenant err: %v, response: %v", err, response),
		)
		return
	}

	state := plan
	createdTenant := result.GetData()
	if createdTenant.GetTenantId() == "" {
		resp.Diagnostics.AddError("Error creating tenant", "The create response did not contain a tenant id.")
		return
	}
	state.Id = types.StringValue(createdTenant.GetTenantId())
	if state.EnablePublicAddress.IsNull() || state.EnablePublicAddress.IsUnknown() {
		state.EnablePublicAddress = types.BoolValue(false)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "instance_id", state.ClusterId)
	ctx = tflog.SetField(ctx, "tenant_id", state.Id)
	tflog.Info(ctx, "Successfully created oceanbase tenant")

	timeoutDuration := operationTimeout(EnvResourceCreateTimeoutSeconds, DefaultResourceCreateTimeout)
	err = waitForTenantCreation(ctx, &state, client, plan.EnablePublicAddress, timeoutDuration)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if err != nil {
		resp.Diagnostics.AddError("Tenant creation did not complete", fmt.Sprintf("Tenant %s was created, but its readiness could not be confirmed: %v", state.Id.ValueString(), err))
	}
}

func (r *tenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenant model.Tenant
	diags := req.State.Get(ctx, &tenant)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := readTenant(ctx, &tenant, client)
	if errors.Is(err, errResourceNotFound) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant info failed",
			fmt.Sprintf("Failed to refresh tenant info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenant)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		for _, diag := range diags {
			tflog.Error(ctx, fmt.Sprintf("State set error: %s", diag.Detail()))
		}
		return
	}
}

func getTenantOperation(ctx context.Context, state, plan *model.Tenant) (string, error) {
	if err := validateTenantCreationInputs(state, plan); err != nil {
		return OperationUnsupported, err
	}
	// TODO: Is it necessary to check cluster status is ONLINE to perform operations
	modifiedSpecs := make([]string, 0)
	// modify tenant unit config, including cpu, memory and unit num
	if (!(plan.UnitNum.IsUnknown() || plan.UnitNum.IsNull()) && !plan.UnitNum.Equal(state.UnitNum)) || (!(plan.Unit.IsUnknown() || plan.Unit.IsNull()) && !plan.Unit.Equal(state.Unit)) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantUnit)
	}
	// modify tenant name
	if !(plan.Name.IsUnknown() || plan.Name.IsNull()) && (!plan.Name.Equal(state.Name)) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantName)
	}
	// modify tenant primary zone
	if !(plan.PrimaryZone.IsUnknown() || plan.PrimaryZone.IsNull()) && (!plan.PrimaryZone.Equal(state.PrimaryZone)) {
		modifiedSpecs = append(modifiedSpecs, SpecTenantPrimaryZone)
	}
	if len(modifiedSpecs) > 1 {
		return OperationUnsupported, errors.New(fmt.Sprintf("Modify %s at once is not supported", strings.Join(modifiedSpecs, ", ")))
	}
	if len(modifiedSpecs) == 1 {
		switch modifiedSpecs[0] {
		case SpecTenantUnit:
			return OperationModifyTenantUnit, nil
		case SpecTenantName:
			return OperationModifyTenantName, nil
		case SpecTenantPrimaryZone:
			return OperationModifyTenantPrimaryZone, nil
		default:
			return OperationUnsupported, errors.New("Unsupported operation")
		}
	}
	return OperationNothing, nil
}

func (r *tenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update tenant is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.Tenant
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state model.Tenant
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := getTenantOperation(ctx, &state, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update is not supported",
			fmt.Sprintf("Got error when get operation err: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Got operation: %s", operation))

	needWaitTask := false
	client := r.provider.clientGenerator.GetClient()
	switch operation {
	case OperationModifyTenantUnit:
		tflog.Info(ctx, "Modify tenant unit")
		needWaitTask = true
		body := obcloudsdk.NewUpdateTenantRequestWithDefaults()
		unitConfig, err := plan.GetUnitConfig()
		if err != nil {
			resp.Diagnostics.AddError(
				"Unit config is invalid",
				fmt.Sprintf("Failed to parse unit config err: %v", err))
			return
		}
		body.SetUnitNum(plan.UnitNum.ValueInt32())
		body.SetCpu(unitConfig.Cpu.ValueInt32())
		body.SetMemory(unitConfig.Memory.ValueInt32())
		result, response, err := client.MultiCloudOpenAPI.UpdateTenant(ctx, plan.ClusterId.ValueString(), plan.Id.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() {
			tflog.Error(ctx, "Modify tenant unit failed")
			resp.Diagnostics.AddError(
				"Modify tenant unit failed",
				fmt.Sprintf("Got error when update tenant unit err: %v, response: %v", err, response))
			return
		}
	case OperationModifyTenantName:
		tflog.Info(ctx, "Modify tenant name")
		body := obcloudsdk.NewModifyTenantNameRequestV2WithDefaults()
		body.SetTenantName(plan.Name.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantName(ctx, plan.ClusterId.ValueString(), plan.Id.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() {
			tflog.Error(ctx, "Modify tenant name failed")
			resp.Diagnostics.AddError(
				"Modify tenant name failed",
				fmt.Sprintf("Got error when update tenant name err: %v, response: %v", err, response))
			return
		}
	case OperationModifyTenantPrimaryZone:
		tflog.Info(ctx, "Modify tenant primary zone")
		needWaitTask = true
		body := obcloudsdk.NewModifyTenantPrimaryZoneRequestWithDefaults()
		body.SetPrimaryZone(plan.PrimaryZone.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyTenantPrimaryZone(ctx, plan.ClusterId.ValueString(), plan.Id.ValueString()).Body(*body).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() {
			tflog.Error(ctx, "Modify tenant primary zone failed")
			resp.Diagnostics.AddError(
				"Modify tenant primary zone failed",
				fmt.Sprintf("Got error when update tenant primary zone err: %v, response: %v", err, response))
			return
		}
	case OperationNothing:
		tflog.Info(ctx, "Nothing to do")
	default:
		tflog.Info(ctx, "Unsupported operation")
	}

	if needWaitTask {
		timeoutDuration := operationTimeout(EnvResourceModifyTimeoutSeconds, DefaultResourceModifyTimeout)
		err = waitForTenantUpdate(ctx, &state, &plan, operation, client, timeoutDuration)
		if err != nil {
			resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			resp.Diagnostics.AddError("Tenant update did not complete", fmt.Sprintf("Tenant %s did not reach the requested state: %v", state.Id.ValueString(), err))
			return
		}
	}

	err = refreshTenant(ctx, &state, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant info error",
			"Failed to refresh tenant info after update",
		)
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *tenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete tenant is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var tenant model.Tenant
	diags := req.State.Get(ctx, &tenant)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tenant.Id.IsNull() || tenant.Id.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenant.ClusterId.IsNull() || tenant.ClusterId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant's cluster id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()
	status, err := tenantListStatus(ctx, &tenant, client)
	if errors.Is(err, errResourceNotFound) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Failed to confirm tenant existence", err.Error())
		return
	}

	// OCP uses PENDING_OFFLINE while the accepted tenant deletion is running.
	if !strings.EqualFold(status, "PENDING_OFFLINE") {
		result, response, err := client.MultiCloudOpenAPI.DeleteTenant(ctx, tenant.ClusterId.ValueString(), tenant.Id.ValueString()).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() {
			resp.Diagnostics.AddError(
				"Failed to delete tenant",
				fmt.Sprintf("Unexpected error while delete tenant: %v, response: %v", err, response))
			return
		}
	}

	timeout := operationTimeout(EnvResourceDeleteTimeoutSeconds, DefaultResourceModifyTimeout)
	if err := waitForRemoval(ctx, timeout, func(ctx context.Context) error {
		return findTenant(ctx, &tenant, client)
	}); err != nil {
		resp.Diagnostics.AddError("Tenant deletion did not complete", fmt.Sprintf("Tenant %s has not been confirmed deleted: %v", tenant.Id.ValueString(), err))
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r *tenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import tenant is called")
	parts := strings.Split(req.ID, ".")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid id format",
			fmt.Sprintf("Unexpected error while parsing resource id %s, the expected format is <cluster_id>.<tenant_id>", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}

func isTenantOnline(tenant *model.Tenant) bool {
	tflog.Debug(context.Background(), fmt.Sprintf("Tenant status when validate: %s", tenant.Status.ValueString()))
	return tenant.Status.ValueString() == model.TenantStatusOnline
}
