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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &databaseResource{}
	_ resource.ResourceWithConfigure   = &databaseResource{}
	_ resource.ResourceWithImportState = &databaseResource{}
)

type databaseResource struct {
	provider *oceanbaseProvider
}

func NewDatabaseResource() resource.Resource {
	return &databaseResource{}
}

func (r *databaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Database Resource",
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Cluster id of the tenant",
				Required:            true,
			},
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant id",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the database",
				Required:            true,
			},
			"charset": schema.StringAttribute{
				MarkdownDescription: "Charset of the database",
				Optional:            true,
			},
			"privileged_users": schema.ListNestedAttribute{
				MarkdownDescription: "Privileged users of the database",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "User name",
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "User role",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Privilege type",
							Computed:            true,
						},
					},
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the database",
				Optional:            true,
				Computed:            true,
			},
			"create_time": schema.StringAttribute{
				MarkdownDescription: "Create time of the database",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the database",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the database",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Database compatible type",
				Computed:            true,
			},
			"collation": schema.StringAttribute{
				MarkdownDescription: "Database collation",
				Computed:            true,
			},
			"required_size": schema.Float64Attribute{
				MarkdownDescription: "Required size of the database",
				Computed:            true,
			},
		},
	}
}

func (r *databaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database"
}

func (r *databaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func generateDatabaseId(clusterId, tenantId, name string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s:%s", clusterId, tenantId, name)))
}

func (r *databaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create database is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.Database

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	createDatabaseBody := obcloudsdk.NewCreateDatabaseRequestV2WithDefaults()
	createDatabaseBody.SetInstanceId(plan.ClusterId.ValueString())
	createDatabaseBody.SetTenantId(plan.TenantId.ValueString())
	createDatabaseBody.SetDatabaseName(plan.Name.ValueString())
	createDatabaseBody.SetEncoding(plan.Charset.ValueString())
	createDatabaseBody.SetDescription(plan.Description.ValueString())

	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateDatabase(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString()).Body(*createDatabaseBody).Execute()
	api.LogAPIResult(ctx, result, response, err)

	if err != nil || response.StatusCode != http.StatusOK || !(result.GetSuccess()) {
		resp.Diagnostics.AddError(
			"Error creating database",
			fmt.Sprintf("Could not create database err: %v, response: %v", err, response),
		)
		return
	}

	state := plan
	state.Id = types.StringValue(generateDatabaseId(plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()))
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error after create the tenant", "Set the state error")
		return
	}
}

func (r *databaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read database is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var database model.Database
	diags := req.State.Get(ctx, &database)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := refreshDatabase(ctx, &database, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh database info failed",
			fmt.Sprintf("Failed to refresh database info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, database)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getDatabaseOperation(ctx context.Context, state, plan *model.Database) (string, error) {
	modifiedSpecs := make([]string, 0)
	if !(plan.Description.IsUnknown() || plan.Description.IsNull()) && !plan.Description.Equal(state.Description) {
		modifiedSpecs = append(modifiedSpecs, SpecDatabaseDescription)
	}
	// currently, it's only one operation, maybe add more operations in the future
	if len(modifiedSpecs) > 1 {
		return OperationUnsupported, errors.New(fmt.Sprintf("Modify %s at once is not supported", strings.Join(modifiedSpecs, ", ")))
	}

	if len(modifiedSpecs) == 1 {
		switch modifiedSpecs[0] {
		case SpecDatabaseDescription:
			return OperationModifyDatabaseDescription, nil
		default:
			return OperationUnsupported, errors.New("Unsupported operation")
		}
	}
	return OperationNothing, nil
}

func (r *databaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update database is called")

	var plan model.Database
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state model.Database
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := getDatabaseOperation(ctx, &state, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update is not supported",
			fmt.Sprintf("Got error when get operation err: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Got operation: %s", operation))

	client := r.provider.clientGenerator.GetClient()
	switch operation {
	case OperationModifyDatabaseDescription:
		tflog.Info(ctx, "Modify database description")
		modifyDescriptionBody := obcloudsdk.NewModifyDatabaseDescriptionRequestV2WithDefaults()
		modifyDescriptionBody.SetDescription(plan.Description.ValueString())
		result, response, err := client.MultiCloudOpenAPI.ModifyDatabaseDescription(ctx, plan.ClusterId.ValueString(), plan.TenantId.ValueString(), plan.Name.ValueString()).Body(*modifyDescriptionBody).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify database description failed")
			resp.Diagnostics.AddError(
				"Modify database description failed",
				fmt.Sprintf("Got error when update database description err: %v, response: %v", err, response))
			return
		}
	case OperationNothing:
		tflog.Info(ctx, "Nothing to do")
	default:
		tflog.Info(ctx, "Unsupported operation")
	}

	// no need to wait, refresh database after modified
	err = refreshDatabase(ctx, &state, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh database info failed",
			fmt.Sprintf("Failed to refresh database info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *databaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete database is called")
	var database model.Database
	diags := req.State.Get(ctx, &database)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if database.TenantId.IsNull() || database.TenantId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant id can't be null",
			"The tenant_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if database.ClusterId.IsNull() || database.ClusterId.IsUnknown() {
		resp.Diagnostics.AddError(
			"Tenant's cluster id can't be null",
			"The cluster_id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	if database.Name.IsNull() || database.Name.IsUnknown() {
		resp.Diagnostics.AddError(
			"User name can't be null",
			"The name field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()
	deleteUserBody := obcloudsdk.NewDeleteDatabasesParamDoWithDefaults()
	deleteUserBody.SetInstanceId(database.ClusterId.ValueString())
	deleteUserBody.SetTenantId(database.TenantId.ValueString())
	deleteUserBody.SetDatabaseNames(fmt.Sprintf("[\"%s\"]", database.Name.ValueString()))
	deleteUserBody.SetForce(true)

	result, response, err := client.MultiCloudOpenAPI.DeleteDatabase(ctx, database.ClusterId.ValueString(), database.TenantId.ValueString()).Body(*deleteUserBody).Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
		resp.Diagnostics.AddError(
			"Failed to delete database",
			fmt.Sprintf("Unexpected error while delete database: %v, response: %v", err, response))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *databaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import database is called")
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
