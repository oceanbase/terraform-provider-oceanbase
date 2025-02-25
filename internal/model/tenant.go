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
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type TenantSpec struct {
	Name                types.String          `tfsdk:"name"`
	ClusterId           types.String          `tfsdk:"cluster_id"`
	Mode                types.String          `tfsdk:"mode"`
	PrimaryZone         types.String          `tfsdk:"primary_zone"`
	Unit                basetypes.ObjectValue `tfsdk:"unit"`
	UnitNum             types.Int32           `tfsdk:"unit_num"`
	Charset             types.String          `tfsdk:"charset"`
	LowerCaseTableNames types.Int32           `tfsdk:"lower_case_table_names"`
	Timezone            types.String          `tfsdk:"timezone"`
	EnablePublicAddress types.Bool            `tfsdk:"enable_public_address"`
	CompatibleMode      types.String          `tfsdk:"compatible_mode"`
	Whitelist           types.String          `tfsdk:"whitelist"`
	Description         types.String          `tfsdk:"description"`
}

type TenantInfo struct {
	Id             types.String          `tfsdk:"id"`
	PrimaryRegion  types.String          `tfsdk:"primary_region"`
	DeployType     types.String          `tfsdk:"deploy_type"`
	CreateTime     types.String          `tfsdk:"create_time"`
	Status         types.String          `tfsdk:"status"`
	Vpc            types.String          `tfsdk:"vpc"`
	DailyMergeTime types.String          `tfsdk:"daily_merge_time"`
	Connections    basetypes.ListValue   `tfsdk:"connections"`
	Locality       basetypes.ListValue   `tfsdk:"locality"`
	Tags           basetypes.ListValue   `tfsdk:"tags"`
	Resource       basetypes.ObjectValue `tfsdk:"resource"`
}

type Tenant struct {
	TenantSpec
	TenantInfo
}

type TenantDataSource struct {
	Tenant
}

type TenantResource struct {
	Tenant
}

type TenantUnitConfig struct {
	Cpu    types.Int32 `tfsdk:"cpu"`
	Memory types.Int32 `tfsdk:"memory"`
}

func (t *Tenant) GetUnitConfig() (*TenantUnitConfig, error) {
	attributes := t.Unit.Attributes()
	cpuAttr, cpuExists := attributes["cpu"]
	memoryAttr, memoryExists := attributes["memory"]
	if cpuExists && memoryExists {
		cpuValue, cpuOk := cpuAttr.(types.Int32)
		memoryValue, memoryOk := memoryAttr.(types.Int32)
		if cpuOk && memoryOk {
			return &TenantUnitConfig{
				Cpu:    cpuValue,
				Memory: memoryValue,
			}, nil
		}
	}
	return nil, errors.New("Invalid unit config")
}

func extractConnections(data *obcloudsdk.TenantDTO) basetypes.ListValue {
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{"type": basetypes.StringType{}, "address_id": basetypes.StringType{}, "status": basetypes.StringType{}},
	}
	connections := make([]attr.Value, len(data.GetTenantConnections()))
	for idx, connection := range data.GetTenantConnections() {
		attributeTypes := map[string]attr.Type{"type": basetypes.StringType{}, "address_id": basetypes.StringType{}, "status": basetypes.StringType{}}
		attributes := map[string]attr.Value{"type": types.StringValue(connection.GetNetworkType()), "address_id": types.StringValue(connection.GetAddressId()), "status": types.StringValue(connection.GetAddressStatus())}
		connectionObject, _ := types.ObjectValue(attributeTypes, attributes)
		connections[idx] = connectionObject
	}
	result, _ := types.ListValue(objectType, connections)
	return result
}

func extractTags(data *obcloudsdk.TenantDTO) basetypes.ListValue {
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{"key": basetypes.StringType{}, "value": basetypes.StringType{}},
	}

	tagList := make([]attr.Value, len(data.GetTagList()))
	for idx, tag := range data.GetTagList() {
		attributeTypes := map[string]attr.Type{"key": basetypes.StringType{}, "value": basetypes.StringType{}}
		attributes := map[string]attr.Value{"key": types.StringValue(tag.GetTagKey()), "value": types.StringValue(tag.GetTagValue())}
		tagObject, _ := types.ObjectValue(attributeTypes, attributes)
		tagList[idx] = tagObject
	}
	result, _ := types.ListValue(objectType, tagList)
	return result
}

func extractLocality(data *obcloudsdk.TenantDTO) basetypes.ListValue {
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{"zone_name": basetypes.StringType{}, "replica_type": basetypes.StringType{}, "region": basetypes.StringType{}, "role": basetypes.StringType{}},
	}

	locality := make([]attr.Value, len(data.GetTenantZones()))
	for idx, zone := range data.GetTenantZones() {
		attributeTypes := map[string]attr.Type{"zone_name": basetypes.StringType{}, "replica_type": basetypes.StringType{}, "region": basetypes.StringType{}, "role": basetypes.StringType{}}
		attributes := map[string]attr.Value{"zone_name": types.StringValue(zone.GetName()), "replica_type": types.StringValue(zone.GetReplicaType()), "region": types.StringValue(zone.GetRegion()), "role": types.StringValue(zone.GetRole())}
		localityObject, _ := types.ObjectValue(attributeTypes, attributes)
		locality[idx] = localityObject
	}
	result, _ := types.ListValue(objectType, locality)
	return result
}

func extractUnitConfig(data *obcloudsdk.TenantDTO) basetypes.ObjectValue {
	attributeTypes := map[string]attr.Type{"cpu": basetypes.Int32Type{}, "memory": basetypes.Int32Type{}}
	attributes := map[string]attr.Value{"cpu": types.Int32Value(data.GetUnitCpu()), "memory": types.Int32Value(data.GetUnitMem())}
	unit, _ := basetypes.NewObjectValue(attributeTypes, attributes)
	return unit
}

func extractResource(data *obcloudsdk.TenantDTO) basetypes.ObjectValue {
	tenantResource := data.GetTenantResource()
	cpuResource := tenantResource.GetCpu()
	memoryResource := tenantResource.GetMemory()
	diskResource := tenantResource.GetDisk()

	cpuAttributeTypes := map[string]attr.Type{"total_cpu": basetypes.Float64Type{}, "used_cpu": basetypes.Float64Type{}, "unit_cpu": basetypes.Float64Type{}}
	memoryAttributeTypes := map[string]attr.Type{"total_memory": basetypes.Float64Type{}, "used_memory": basetypes.Float64Type{}, "unit_memory": basetypes.Float64Type{}}
	diskAttributeTypes := map[string]attr.Type{"total_disk_size": basetypes.Float64Type{}, "used_disk_size": basetypes.Float64Type{}}
	cpuAttributes := map[string]attr.Value{"total_cpu": types.Float64Value(cpuResource.GetTotalCpu()), "used_cpu": types.Float64Value(cpuResource.GetUsedCpu()), "unit_cpu": types.Float64Value(cpuResource.GetUnitCpu())}
	memoryAttributes := map[string]attr.Value{"total_memory": types.Float64Value(memoryResource.GetTotalMemory()), "used_memory": types.Float64Value(memoryResource.GetUsedMemory()), "unit_memory": types.Float64Value(memoryResource.GetUnitMemory())}
	diskAttributes := map[string]attr.Value{"total_disk_size": types.Float64Value(diskResource.GetTotalDiskSize()), "used_disk_size": types.Float64Value(diskResource.GetUsedDiskSize())}

	cpu, _ := basetypes.NewObjectValue(cpuAttributeTypes, cpuAttributes)
	memory, _ := basetypes.NewObjectValue(memoryAttributeTypes, memoryAttributes)
	disk, _ := basetypes.NewObjectValue(diskAttributeTypes, diskAttributes)

	cpuType := basetypes.ObjectType{
		AttrTypes: cpuAttributeTypes,
	}
	memoryType := basetypes.ObjectType{
		AttrTypes: memoryAttributeTypes,
	}
	diskType := basetypes.ObjectType{
		AttrTypes: diskAttributeTypes,
	}
	attributeTypes := map[string]attr.Type{"cpu": cpuType, "memory": memoryType, "disk": diskType}
	attributes := map[string]attr.Value{"cpu": cpu, "memory": memory, "disk": disk}
	resource, _ := basetypes.NewObjectValue(attributeTypes, attributes)
	return resource
}

func (t *Tenant) RefreshTenant(data *obcloudsdk.TenantDTO) {
	t.Id = types.StringValue(data.GetTenantId())
	t.Name = types.StringValue(data.GetTenantName())
	t.CreateTime = types.StringValue(data.GetCreateTime())
	t.Mode = types.StringValue(data.GetTenantMode())
	t.PrimaryRegion = types.StringValue(data.GetPrimaryRegion())
	t.PrimaryZone = types.StringValue(data.GetPrimaryZone())
	t.Status = types.StringValue(data.GetStatus())
	t.Vpc = types.StringValue(data.GetVpcId())
	t.DeployType = types.StringValue(data.GetDeployType())
	t.Description = types.StringValue(data.GetDescription())
	t.EnablePublicAddress = types.BoolValue(data.GetEnableProxyPublicLink())
	t.DailyMergeTime = types.StringValue(data.GetDataMergeTime())
	t.Timezone = types.StringValue(data.GetTimeZone())
	t.Charset = types.StringValue(data.GetCharset())
	t.CompatibleMode = types.StringValue(data.GetTenantCompatibilityMode())
	t.UnitNum = types.Int32Value(data.GetUnitNum())
	t.Unit = extractUnitConfig(data)
	t.Connections = extractConnections(data)
	t.Tags = extractTags(data)
	t.LowerCaseTableNames = types.Int32Value(data.GetLowerCaseTableNamesParam())
	t.Locality = extractLocality(data)
	t.Resource = extractResource(data)
}
