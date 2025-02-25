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
	_ datasource.DataSource              = &databaseDataSource{}
	_ datasource.DataSourceWithConfigure = &databaseDataSource{}
)

type databaseDataSource struct {
	provider *oceanbaseProvider
}

func NewDatabaseDataSource() datasource.DataSource {
	return &databaseDataSource{}
}

func (d *databaseDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				Computed:            true,
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

func (d *databaseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database"
}

// Configure adds the provider configured client to the data source.
func (d *databaseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *databaseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read database is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var database model.Database
	diags := req.Config.Get(ctx, &database)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the database", "")
		return
	}

	client := d.provider.clientGenerator.GetClient()
	err := refreshDatabase(ctx, &database, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh database info failed",
			fmt.Sprintf("Failed to refresh database info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, database)
	resp.Diagnostics.Append(diags...)
}
