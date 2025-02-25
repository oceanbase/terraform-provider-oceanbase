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
	_ datasource.DataSource              = &tenantDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantDataSource{}
)

type tenantDataSource struct {
	provider *oceanbaseProvider
}

func NewTenantDataSource() datasource.DataSource {
	return &tenantDataSource{}
}

func (d *tenantDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant Resource",
		Attributes: map[string]schema.Attribute{
			// required fields to query tenant
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the tenant",
				Required:            true,
			},
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Id of cluster to create the tenant",
				Required:            true,
			},

			// tenant attributes
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the tenant",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Tenant mode",
				Computed:            true,
			},
			"primary_zone": schema.StringAttribute{
				MarkdownDescription: "Primary zone of the tenant",
				Computed:            true,
			},
			"charset": schema.StringAttribute{
				MarkdownDescription: "Charset of the tenant",
				Computed:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "Timezone of the tenant",
				Computed:            true,
			},
			"unit_num": schema.Int32Attribute{
				MarkdownDescription: "Unit num of the tenant",
				Computed:            true,
			},
			"unit": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Unit config of the tenant",
				Attributes: map[string]schema.Attribute{
					"cpu": schema.Int32Attribute{
						MarkdownDescription: "Cpu config of the unit",
						Computed:            true,
					},
					"memory": schema.Int32Attribute{
						MarkdownDescription: "Memory config of the unit",
						Computed:            true,
					},
				},
			},
			"enable_public_address": schema.BoolAttribute{
				MarkdownDescription: "Whether create a public address for the tenant",
				Computed:            true,
			},
			"compatible_mode": schema.StringAttribute{
				MarkdownDescription: "Compatible mode of the tenant",
				Computed:            true,
			},
			"whitelist": schema.StringAttribute{
				MarkdownDescription: "Whitelist of the tenant",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the tenant",
				Computed:            true,
			},
			"lower_case_table_names": schema.Int32Attribute{
				MarkdownDescription: "Configuration for table names",
				Computed:            true,
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
								Computed:            true,
							},
							"used_cpu": schema.Float64Attribute{
								MarkdownDescription: "Used cpu of the tenant",
								Computed:            true,
							},
							"unit_cpu": schema.Float64Attribute{
								MarkdownDescription: "Unit config cpu of the tenant",
								Computed:            true,
							},
						},
					},
					"disk": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Disk resource info of the tenant",
						Attributes: map[string]schema.Attribute{
							"total_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Total disk size of the tenant",
								Computed:            true,
							},
							"used_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Used disk size of the tenant",
								Computed:            true,
							},
						},
					},
					"memory": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Memory resource info of the tenant",
						Attributes: map[string]schema.Attribute{
							"total_memory": schema.Float64Attribute{
								MarkdownDescription: "Total memory of the tenant",
								Computed:            true,
							},
							"used_memory": schema.Float64Attribute{
								MarkdownDescription: "Used memory of the tenant",
								Computed:            true,
							},
							"unit_memory": schema.Float64Attribute{
								MarkdownDescription: "Unit config memory of the tenant",
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *tenantDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

// Configure adds the provider configured client to the data source.
func (d *tenantDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *tenantDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenant model.Tenant
	diags := req.Config.Get(ctx, &tenant)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the tenant", "")
		return
	}

	client := d.provider.clientGenerator.GetClient()
	err := refreshTenant(ctx, &tenant, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant info failed",
			fmt.Sprintf("Failed to refresh tenant info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenant)
	resp.Diagnostics.Append(diags...)
}
