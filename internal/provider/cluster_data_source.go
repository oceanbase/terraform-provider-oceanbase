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
	_ datasource.DataSource              = &clusterDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterDataSource{}
)

type clusterDataSource struct {
	provider *oceanbaseProvider
}

func NewClusterDataSource() datasource.DataSource {
	return &clusterDataSource{}
}

func (d *clusterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Cluster DataSource",
		Attributes: map[string]schema.Attribute{
			// parameters to query obcluster
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the cluster",
				Required:            true,
			},

			// obcluster attributes
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Project id of the cluster",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster",
				Computed:            true,
			},
			"cloud_provider": schema.StringAttribute{
				MarkdownDescription: "Cloud provider of the cluster",
				Computed:            true,
			},
			"instance_class": schema.StringAttribute{
				MarkdownDescription: "Instance class of the cluster, e.g. 4c16g",
				Computed:            true,
			},
			"vpc": schema.StringAttribute{
				MarkdownDescription: "VPC where the cluster is created",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Version of the cluster",
				Computed:            true,
			},
			"create_time": schema.StringAttribute{
				MarkdownDescription: "Create time of the cluster",
				Computed:            true,
			},
			"series": schema.StringAttribute{
				MarkdownDescription: "Series of the cluster, e.g. normal, normal_kv",
				Computed:            true,
			},
			"payment_type": schema.StringAttribute{
				MarkdownDescription: "Payment type of the cluster",
				Computed:            true,
			},
			"sale_channel": schema.StringAttribute{
				MarkdownDescription: "Sales channel",
				Computed:            true,
			},
			"instance_type": schema.StringAttribute{
				MarkdownDescription: "Business scenario of cluster, e.g. CLUSTER, ANALYTICAL_CLUSTER, KV_CLUSTER, SHARED_STORAGE_CLUSTER",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of the cluster, such as ONLINE, PENDING etc",
				Computed:            true,
			},
			"is_stopped": schema.BoolAttribute{
				MarkdownDescription: "Whether the cluster is supposed to be stopped",
				Computed:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region, such as us-east-1",
				Computed:            true,
			},
			"zones": schema.ListAttribute{
				MarkdownDescription: "Zones of the cluster",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"replica_mode": schema.StringAttribute{
				MarkdownDescription: "Replica mode of the cluster",
				Computed:            true,
			},
			"cpu_arch": schema.StringAttribute{
				MarkdownDescription: "cpu architecture, x86_64 or aarch64",
				Computed:            true,
			},
			"node_num": schema.Int32Attribute{
				MarkdownDescription: "Total node number of the cluster, it's calculated by summing up the observer numbers in each zone",
				Computed:            true,
			},
			"disk_size": schema.Int32Attribute{
				MarkdownDescription: "Disk size of observer",
				Computed:            true,
			},
			"deploy_mode": schema.StringAttribute{
				MarkdownDescription: "Deploy mode of the cluster, such as 1-1-1",
				Computed:            true,
			},
			"deploy_type": schema.StringAttribute{
				MarkdownDescription: "Deploy type of the cluster, Enums: SINGLE, MULTIPLE",
				Computed:            true,
			},
			"stop_time": schema.StringAttribute{
				MarkdownDescription: "Cluster stop time",
				Computed:            true,
			},
			"resource": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cpu": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Cpu resource info of the cluster",
						Attributes: map[string]schema.Attribute{
							"total_cpu": schema.Float64Attribute{
								MarkdownDescription: "Total cpu of the cluster",
								Computed:            true,
							},
							"used_cpu": schema.Float64Attribute{
								MarkdownDescription: "Used cpu of the cluster",
								Computed:            true,
							},
						},
					},
					"disk": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Disk resource info of the cluster",
						Attributes: map[string]schema.Attribute{
							"total_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Total disk size of the cluster",
								Computed:            true,
							},
							"used_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Used disk size of the cluster",
								Computed:            true,
							},
						},
					},
					"memory": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Memory resource info of the cluster",
						Attributes: map[string]schema.Attribute{
							"total_memory": schema.Float64Attribute{
								MarkdownDescription: "Total memory of the cluster",
								Computed:            true,
							},
							"used_memory": schema.Float64Attribute{
								MarkdownDescription: "Used memory of the cluster",
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *clusterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Configure adds the provider configured client to the data source.
func (d *clusterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *clusterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read cluster is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	tflog.Info(ctx, "Start data source read...")

	var cluster model.Cluster
	diags := req.Config.Get(ctx, &cluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the cluster", "")
		return
	}

	if cluster.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Cluster's id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	ctx = tflog.SetField(ctx, "Id", cluster.Id.ValueString())
	tflog.Debug(ctx, "Read cluster data source")

	client := d.provider.clientGenerator.GetClient()
	err := refreshCluster(ctx, &cluster, client)

	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh obcluster info failed",
			fmt.Sprintf("Failed to refresh obcluster info err: %v", err))

	}
	diags = resp.State.Set(ctx, cluster)
	resp.Diagnostics.Append(diags...)
}
