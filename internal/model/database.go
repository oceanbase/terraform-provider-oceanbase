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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type Database struct {
	Id              types.String        `tfsdk:"id"`
	Name            types.String        `tfsdk:"name"`
	ClusterId       types.String        `tfsdk:"cluster_id"`
	TenantId        types.String        `tfsdk:"tenant_id"`
	Charset         types.String        `tfsdk:"charset"`
	CreateTime      types.String        `tfsdk:"create_time"`
	Description     types.String        `tfsdk:"description"`
	Status          types.String        `tfsdk:"status"`
	Type            types.String        `tfsdk:"type"`
	Collation       types.String        `tfsdk:"collation"`
	RequiredSize    types.Float64       `tfsdk:"required_size"`
	PrivilegedUsers basetypes.ListValue `tfsdk:"privileged_users"`
}

func extractPrivilegedUsers(data *obcloudsdk.DescribeDatabasesResponseV2) basetypes.ListValue {
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{"name": basetypes.StringType{}, "role": basetypes.StringType{}, "type": basetypes.StringType{}},
	}
	privilegedUsers := make([]attr.Value, 0)
	for _, user := range data.GetUsers() {
		attributeTypes := map[string]attr.Type{"name": basetypes.StringType{}, "role": basetypes.StringType{}, "type": basetypes.StringType{}}
		attributes := map[string]attr.Value{"name": types.StringValue(user.GetUserName()), "role": types.StringValue(user.GetRole()), "type": types.StringValue(user.GetUserType())}
		privilegedUserObject, _ := basetypes.NewObjectValue(attributeTypes, attributes)
		privilegedUsers = append(privilegedUsers, privilegedUserObject)
	}
	result, _ := types.ListValue(objectType, privilegedUsers)
	return result
}

func (d *Database) Refresh(data *obcloudsdk.DescribeDatabasesResponseV2) {
	d.TenantId = types.StringValue(data.GetTenantId())
	d.Name = types.StringValue(data.GetDatabaseName())
	d.CreateTime = types.StringValue(data.GetGmtCreate())
	d.Charset = types.StringValue(data.GetEncoding())
	d.Status = types.StringValue(data.GetStatus())
	d.Description = types.StringValue(data.GetDescription())
	d.RequiredSize = types.Float64Value(data.GetRequiredSize())
	d.Collation = types.StringValue(data.GetCollation())
	d.Type = types.StringValue(data.GetDbType())
	d.PrivilegedUsers = extractPrivilegedUsers(data)
}
