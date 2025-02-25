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
	_ datasource.DataSource              = &tenantSecurityGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &tenantSecurityGroupDataSource{}
)

type tenantSecurityGroupDataSource struct {
	provider *oceanbaseProvider
}

func NewTenantSecurityGroupDataSource() datasource.DataSource {
	return &tenantSecurityGroupDataSource{}
}

func (d *tenantSecurityGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tenant Security Group Resource",
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
				MarkdownDescription: "Name of the security group",
				Required:            true,
			},
			"cidrs": schema.ListAttribute{
				MarkdownDescription: "Security group cidrs",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the security group",
				Computed:            true,
			},
		},
	}
}

func (d *tenantSecurityGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_security_group"
}

// Configure adds the provider configured client to the data source.
func (d *tenantSecurityGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if d.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (d *tenantSecurityGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read tenant security group is called")
	if d.provider == nil || !d.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var tenantSecurityGroup model.TenantSecurityGroup
	diags := req.Config.Get(ctx, &tenantSecurityGroup)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error loading the tenant user", "")
		return
	}

	client := d.provider.clientGenerator.GetClient()
	err := refreshTenantSecurityGroup(ctx, &tenantSecurityGroup, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh tenant security group info failed",
			fmt.Sprintf("Failed to refresh tenant security group info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, tenantSecurityGroup)
	resp.Diagnostics.Append(diags...)
}
