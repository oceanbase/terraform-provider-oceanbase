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
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.Provider = &oceanbaseProvider{}

type oceanbaseProvider struct {
	clientGenerator *api.Generator
	configured      bool
	version         string
}

type providerData struct {
	AccessKey types.String `tfsdk:"access_key"`
	SecretKey types.String `tfsdk:"secret_key"`
	Site      types.String `tfsdk:"site"`
}

func (p *oceanbaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring oceanbase provider")

	// get providerData
	var data providerData
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a access_key to the provider
	var accessKey string
	if data.AccessKey.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as access_key",
		)
		return
	}

	if data.AccessKey.IsNull() {
		accessKey = os.Getenv(EnvNameAccessKey)
	} else {
		accessKey = data.AccessKey.ValueString()
	}

	if accessKey == "" {
		resp.Diagnostics.AddError(
			"Unable to find access_key",
			"Access_key cannot be empty",
		)
		return
	}

	// User must provide a secret_key to the provider
	var secretKey string
	if data.SecretKey.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as secret_key",
		)
		return
	}

	if data.SecretKey.IsNull() {
		secretKey = os.Getenv(EnvNameSecretKey)
	} else {
		secretKey = data.SecretKey.ValueString()
	}

	if secretKey == "" {
		resp.Diagnostics.AddError(
			"Unable to find secret_key",
			"Secret_key cannot be empty",
		)
		return
	}
	site := data.Site.ValueString()
	ctx = tflog.SetField(ctx, "site", site)
	tflog.Debug(ctx, "Creating oceanbaseapi client")
	p.clientGenerator = &api.Generator{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Site:      site,
	}

	p.configured = true
	resp.ResourceData = p
	resp.DataSourceData = p
}

// Metadata returns the provider type name.
func (p *oceanbaseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "oceanbase"
	resp.Version = p.version
}

func (p *oceanbaseProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClusterResource,
		NewTenantResource,
		NewTenantSecurityGroupResource,
		NewTenantUserResource,
		NewDatabaseResource,
		NewTenantConnectionResource,
	}
}

func (p *oceanbaseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewClusterDataSource,
		NewTenantDataSource,
		NewTenantConnectionDataSource,
		NewTenantUserDataSource,
		NewTenantSecurityGroupDataSource,
		NewDatabaseDataSource,
	}
}

// Schema defines the provider-level schema for configuration data.
func (p *oceanbaseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_key": schema.StringAttribute{
				MarkdownDescription: "access key to access oceanbase cloud",
				Required:            true,
				Sensitive:           true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "secret key to access oceanbase cloud",
				Required:            true,
				Sensitive:           true,
			},
			"site": schema.StringAttribute{
				MarkdownDescription: "oceanbase cloud site, Enum: \"CHINA\", \"GLOBAL\"",
				Required:            true,
				Sensitive:           false,
			},
		},
	}
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &oceanbaseProvider{
			version: version,
		}
	}
}
