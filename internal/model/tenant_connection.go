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

package model

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type PrivateLinkConfig struct {
	UserAccount string `tfsdk:"user_account"`
	UserVpc     string `tfsdk:"user_vpc"`
	UserVswitch string `tfsdk:"user_vswitch"`
}

func (tc *TenantConnection) GetPrivateLinkConfig() (*PrivateLinkConfig, error) {
	return parseObjectValue[PrivateLinkConfig](tc.PrivateLink)
}

type TenantConnection struct {
	Id              types.String          `tfsdk:"id"`
	ClusterId       types.String          `tfsdk:"cluster_id"`
	TenantId        types.String          `tfsdk:"tenant_id"`
	Type            types.String          `tfsdk:"type"`
	PrivateLink     basetypes.ObjectValue `tfsdk:"private_link"`
	Status          types.String          `tfsdk:"status"`
	ConnectionZones []types.String        `tfsdk:"connection_zones"`
	EnableSsl       types.Bool            `tfsdk:"enable_ssl"`
	EnableRpc       types.Bool            `tfsdk:"enable_rpc"`
	Role            types.String          `tfsdk:"role"`
	UserNameFormat  types.String          `tfsdk:"user_name_format"`
	ProxyClusterId  types.String          `tfsdk:"proxy_cluster_id"`
	Intranet        basetypes.ObjectValue `tfsdk:"intranet"`
	Internet        basetypes.ObjectValue `tfsdk:"internet"`
}

func extractIntranetInfo(data *obcloudsdk.TenantConnectionDTO) basetypes.ObjectValue {
	attributeTypes := map[string]attr.Type{"domain": basetypes.StringType{}, "status": basetypes.StringType{}, "rpc_port": basetypes.Int32Type{}, "peering_max_connections": basetypes.Int32Type{}}
	attributes := map[string]attr.Value{"domain": types.StringValue(data.GetIntranetDomain()), "status": types.StringValue(data.GetIntranetAddressStatus()), "rpc_port": types.Int32Value(data.GetIntranetRpcPort()), "peering_max_connections": types.Int32Value(data.GetIntranetPeeringMaxConnectionNum())}
	intranet, _ := basetypes.NewObjectValue(attributeTypes, attributes)
	return intranet
}

func extractInternetInfo(data *obcloudsdk.TenantConnectionDTO) basetypes.ObjectValue {
	attributeTypes := map[string]attr.Type{"domain": basetypes.StringType{}, "status": basetypes.StringType{}, "rpc_port": basetypes.Int32Type{}, "max_connections": basetypes.Int32Type{}, "address": basetypes.StringType{}, "port": basetypes.Int32Type{}}
	attributes := map[string]attr.Value{"domain": types.StringValue(data.GetInternetDomain()), "status": types.StringValue(data.GetInternetAddressStatus()), "rpc_port": types.Int32Value(data.GetInternetRpcPort()), "max_connections": types.Int32Value(data.GetInternetMaxConnectionNum()), "address": types.StringValue(data.GetInternetAddress()), "port": types.Int32Value(data.GetInternetPort())}
	internet, _ := basetypes.NewObjectValue(attributeTypes, attributes)
	return internet
}

func (tc *TenantConnection) Refresh(data *obcloudsdk.TenantConnectionDTO) {
	tc.Id = types.StringValue(data.GetAddressId())
	// tc.Type = types.String(data.GetAddressType())
	tc.Status = types.StringValue(data.GetAddressStatus())
	// tc.ConnectionZones = types.ListValue(data.GetConnectionZones())
	tc.EnableSsl = types.BoolValue(data.GetUseSSL())
	tc.EnableRpc = types.BoolValue(data.GetEnableRPc())
	tc.Role = types.StringValue(data.GetRole())
	tc.UserNameFormat = types.StringValue(data.GetUserNameFormat())
	tc.ProxyClusterId = types.StringValue(data.GetProxyClusterId())
	switch strings.ToUpper(tc.Type.ValueString()) {
	case ConnectionTypeIntranet:
		tc.Intranet = extractIntranetInfo(data)
	case ConnectionTypeInternet:
		tc.Internet = extractInternetInfo(data)
	}
}
