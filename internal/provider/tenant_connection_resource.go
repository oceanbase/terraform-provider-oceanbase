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
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &tenantConnectionResource{}
	_ resource.ResourceWithConfigure   = &tenantConnectionResource{}
	_ resource.ResourceWithImportState = &tenantConnectionResource{}
)

type tenantConnectionResource struct {
	provider *oceanbaseProvider
}

func NewTenantConnectionResource() resource.Resource {
	return &tenantConnectionResource{}
}

func (r *tenantConnectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant Connection Resource",
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster id of the tenant",
				Required:            true,
			},
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant id",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the connection",
				Required:            true,
			},
			"private_link": schema.SingleNestedAttribute{
				MarkdownDescription: "Private link configuration",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"user_account": schema.StringAttribute{
						MarkdownDescription: "User account",
						Required:            true,
					},
					"user_vpc": schema.StringAttribute{
						MarkdownDescription: "User vpc",
						Required:            true,
					},
					"user_vswitch": schema.StringAttribute{
						MarkdownDescription: "User vswitch",
						Required:            true,
					},
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the connection",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the connection",
				Computed:            true,
			},
			"connection_zones": schema.ListAttribute{
				MarkdownDescription: "Zones of the connection",
				Computed:            true,
				ElementType:         types.StringType,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: "Whether ssl is enabled",
				Optional:            true,
				Computed:            true,
			},
			"enable_rpc": schema.BoolAttribute{
				MarkdownDescription: "Whether rpc is enabled",
				Optional:            true,
				Computed:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Connection role",
				Computed:            true,
			},
			"user_name_format": schema.StringAttribute{
				MarkdownDescription: "User name format, available values are USER, USER_AND_TENANT, USER_AND_TENANT_AND_CLUSTER",
				Computed:            true,
			},
			"proxy_cluster_id": schema.StringAttribute{
				MarkdownDescription: "Obproxy cluster id",
				Computed:            true,
			},
			"intranet": schema.SingleNestedAttribute{
				MarkdownDescription: "Intranet properties",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"domain": schema.StringAttribute{
						MarkdownDescription: "Domain name",
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: "Status",
						Computed:            true,
					},
					"rpc_port": schema.Int32Attribute{
						MarkdownDescription: "Rpc port",
						Computed:            true,
					},
					"peering_max_connections": schema.Int32Attribute{
						MarkdownDescription: "Perring max connections",
						Computed:            true,
					},
				},
			},
			"internet": schema.SingleNestedAttribute{
				MarkdownDescription: "Internet properties",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"domain": schema.StringAttribute{
						MarkdownDescription: "Domain name",
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: "Status",
						Computed:            true,
					},
					"rpc_port": schema.Int32Attribute{
						MarkdownDescription: "Rpc port",
						Computed:            true,
					},
					"max_connections": schema.Int32Attribute{
						MarkdownDescription: "Max connections",
						Computed:            true,
					},
					"address": schema.StringAttribute{
						MarkdownDescription: "Connection address",
						Computed:            true,
					},
					"port": schema.Int32Attribute{
						MarkdownDescription: "Sql port",
						Computed:            true,
					},
				},
			},
		},
	}
}

func (r *tenantConnectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_connection"
}

func (r *tenantConnectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (r *tenantConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create tenant connection is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.TenantConnection

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// only internet is supported for now
	if strings.ToUpper(plan.Type.ValueString()) != model.ConnectionTypeInternet {
		resp.Diagnostics.AddError("Invalid type", "Only internet is supported for now")
		return
	}

	client := r.provider.clientGenerator.GetClient()
	serviceType := model.ServiceTypeInternet

	switch strings.ToUpper(plan.Type.ValueString()) {
	case model.ConnectionTypeIntranet:
		serviceType = model.ServiceTypeIntranet
	case model.ConnectionTypeInternet:
		serviceType = model.ServiceTypeInternet
	}
	body := obcloudsdk.NewCreateMcTenantAddressRequestWithDefaults()
	body.SetVipServiceType(serviceType)
	// switch plan.Type.ValueString() {
	// case model.ConnectionTypeIntranet:
	// 	tflog.Info(ctx, "Create private connection")
	// 	privateLinkConfig, err := plan.GetPrivateLinkConfig()
	// 	if err != nil {
	// 		resp.Diagnostics.AddError(
	// 			"Error getting tenant connection private_link config",
	// 			fmt.Sprintf("Could not parse private_link configuration err: %v", err),
	// 		)
	// 		return
	// 	}
	// 	body.SetUserVpcOwnerId(privateLinkConfig.UserAccount)
	// 	body.SetUserVpcId(privateLinkConfig.UserVpc)
	// 	body.SetUserVswitchId(privateLinkConfig.UserVswitch)
	// case model.ConnectionTypeInternet:
	// 	tflog.Info(ctx, "Create public connection")
	// default:
	// 	resp.Diagnostics.AddError(
	// 		"Unsupported connection type",
	// 		fmt.Sprintf("Connection type is not supported: %s", plan.Type.ValueString()),
	// 	)
	// 	return
	// }

	result, response, err := client.MultiCloudOpenAPI.CreateTenantAddress(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).Body(*body).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(result.GetSuccess()) {
		resp.Diagnostics.AddError(
			"Error creating tenant connection",
			fmt.Sprintf("Could not create tenant connection err: %v, response: %v", err, response),
		)
		return
	}

	data := result.GetData()
	state := plan
	state.Id = types.StringValue(data.GetAddressId())

	timeoutDuration := DefaultResourceCreateTimeout
	envTimeoutSeconds := os.Getenv(EnvResourceCreateTimeoutSeconds)
	if envTimeoutSeconds != "" {
		timeoutSeconds, err := strconv.ParseInt(envTimeoutSeconds, 10, 64)
		if err == nil {
			timeoutDuration = time.Duration(timeoutSeconds) * time.Second
		}
	}
	err = retry.RetryContext(ctx, timeoutDuration,
		waitConnectionOnline(ctx, &state, client, isTenantConnectioOnline))
	// Just add an error log here but consider the resource is created successfully
	// state is already refreshed while waiting, no need to refresh again here
	if err != nil {
		tflog.Error(ctx, "Wait for tenant connection online timed out, just save current state")
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error after create the tenant", "Set the state error")
		return
	}
}

func (r *tenantConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant connection is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	tflog.Debug(ctx, "Start Read state...")

	var tenantConnection model.TenantConnection
	diags := req.State.Get(ctx, &tenantConnection)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := refreshTenantConnection(ctx, &tenantConnection, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant connection info failed",
			fmt.Sprintf("Failed to refresh tenant connection info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantConnection)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *tenantConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// nothing to update, just add a log
	tflog.Debug(ctx, "Update tenant connection is called")
}

func (r *tenantConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete tenant connection is called")
	var tenantConnection model.TenantConnection
	diags := req.State.Get(ctx, &tenantConnection)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if tenantConnection.Id.IsNull() || tenantConnection.Id.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant connection id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantConnection.ClusterId.IsNull() || tenantConnection.ClusterId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Cluster id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if tenantConnection.TenantId.IsNull() || tenantConnection.TenantId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()

	result, response, err := client.MultiCloudOpenAPI.DeleteTenantAddress(ctx, tenantConnection.ClusterId.ValueString(), tenantConnection.TenantId.ValueString()).AddressId(tenantConnection.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
		resp.Diagnostics.AddError(
			"Failed to delete tenant connection",
			fmt.Sprintf("Unexpected error while delete tenant connection: %v, response: %v", err, response))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *tenantConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import tenant connection is called")
	parts := strings.Split(req.ID, ".")
	if len(parts) != 4 {
		resp.Diagnostics.AddError(
			"Invalid id format",
			fmt.Sprintf("Unexpected error while parsing resource id %s, the expected format is <cluster_id>.<tenant_id>.<type>.<address_id>", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("cluster_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tenant_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("type"), parts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[3])...)
}

func isTenantConnectioOnline(state *model.TenantConnection) bool {
	return state.Status.ValueString() == model.ConnectionStatusOnline
}

func waitConnectionOnline(ctx context.Context, state *model.TenantConnection, client *obcloudsdk.APIClient, statusValidationFunction func(*model.TenantConnection) bool) retry.RetryFunc {
	return func() *retry.RetryError {
		err := refreshTenantConnection(ctx, state, client)
		if err != nil {
			return retry.RetryableError(fmt.Errorf("Failed to refresh tenant connection info, err: %v", err))
		}
		if !statusValidationFunction(state) {
			return retry.RetryableError(fmt.Errorf("Tenant connection status is still not online, current status: %s", state.Status.ValueString()))
		}
		return nil
	}
}
