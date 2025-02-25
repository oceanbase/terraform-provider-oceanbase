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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
)

var (
	_ datasource.DataSource              = &tenantConnectionDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantConnectionDataSource{}
)

type tenantConnectionDataSource struct {
	provider *oceanbaseProvider
}

func NewTenantConnectionDataSource() datasource.DataSource {
	return &tenantConnectionDataSource{}
}

func (d *tenantConnectionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the connection",
				Required:            true,
			},
			"private_link": schema.SingleNestedAttribute{
				MarkdownDescription: "Private link configuration",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"user_account": schema.StringAttribute{
						MarkdownDescription: "User account",
						Computed:            true,
					},
					"user_vpc": schema.StringAttribute{
						MarkdownDescription: "User vpc",
						Computed:            true,
					},
					"user_vswitch": schema.StringAttribute{
						MarkdownDescription: "User vswitch",
						Computed:            true,
					},
				},
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the connection",
				Computed:            true,
			},
			"connection_zones": schema.ListAttribute{
				MarkdownDescription: "Zones of the connection",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: "Whether ssl is enabled",
				Computed:            true,
			},
			"enable_rpc": schema.BoolAttribute{
				MarkdownDescription: "Whether rpc is enabled",
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

func (d *tenantConnectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_connection"
}

// Configure adds the provider configured client to the data source.
func (d *tenantConnectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *tenantConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant connection is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenantConnection model.TenantConnection
	diags := req.Config.Get(ctx, &tenantConnection)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the tenant user", "")
		return
	}

	client := d.provider.clientGenerator.GetClient()
	err := refreshTenantConnection(ctx, &tenantConnection, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant connection info failed",
			fmt.Sprintf("Failed to refresh tenant connection info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantConnection)
	resp.Diagnostics.Append(diags...)
}
