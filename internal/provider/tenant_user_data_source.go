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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
)

var (
	_ datasource.DataSource              = &tenantUserDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantUserDataSource{}
)

type tenantUserDataSource struct {
	provider *oceanbaseProvider
}

func NewTenantUserDataSource() datasource.DataSource {
	return &tenantUserDataSource{}
}

func (d *tenantUserDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant User Datasource",
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
			// computed fields to create tenant
			"password": schema.StringAttribute{
				MarkdownDescription: "Password of the user",
				Computed:            true,
				Sensitive:           true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type of the user",
				Computed:            true,
			},
			"privileges": schema.ListNestedAttribute{
				MarkdownDescription: "Privileges of the user",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"database": schema.StringAttribute{
							MarkdownDescription: "Database name",
							Computed:            true,
						},
						"table": schema.StringAttribute{
							MarkdownDescription: "Table name",
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "User role",
							Required:            true,
						},
						"privileges": schema.StringAttribute{
							MarkdownDescription: "Privileges",
							Computed:            true,
						},
						"with_grant": schema.BoolAttribute{
							MarkdownDescription: "With grant privilege",
							Computed:            true,
						},
					},
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the user",
				Computed:            true,
			},
			"is_locked": schema.BoolAttribute{
				MarkdownDescription: "Whether the user is locked",
				Computed:            true,
			},
			"encryption_type": schema.StringAttribute{
				MarkdownDescription: "Encryption type of the user",
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

func (d *tenantUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_user"
}

// Configure adds the provider configured client to the data source.
func (d *tenantUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *tenantUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant user is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenantUser model.TenantUser
	diags := req.Config.Get(ctx, &tenantUser)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the tenant user", "")
		return
	}

	client := d.provider.clientGenerator.GetClient()
	err := refreshTenantUser(ctx, &tenantUser, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant user info failed",
			fmt.Sprintf("Failed to refresh tenant user info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantUser)
	resp.Diagnostics.Append(diags...)
}
