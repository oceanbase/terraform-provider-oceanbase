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

type Resource struct {
	Cpu    *ResourceCpu    `tfsdk:"cpu"`
	Disk   *ResourceDisk   `tfsdk:"disk"`
	Memory *ResourceMemory `tfsdk:"memory"`
}

type ResourceCpu struct {
	TotalCpu types.Float64 `tfsdk:"total_cpu"`
	UsedCpu  types.Float64 `tfsdk:"used_cpu"`
}

type ResourceDisk struct {
	TotalDiskSize types.Float64 `tfsdk:"total_disk_size"`
	UsedDiskSize  types.Float64 `tfsdk:"used_disk_size"`
}

type ResourceMemory struct {
	TotalMemory types.Float64 `tfsdk:"total_memory"`
	UsedMemory  types.Float64 `tfsdk:"used_memory"`
}

func newResource(r obcloudsdk.ResourceDoV2) *Resource {
	cpuResource := r.GetCpu()
	memoryResource := r.GetMemory()
	diskResource := r.GetDisk()
	return &Resource{
		Cpu: &ResourceCpu{
			TotalCpu: types.Float64Value(cpuResource.GetTotalCpu()),
			UsedCpu:  types.Float64Value(cpuResource.GetUsedCpu()),
		},
		Memory: &ResourceMemory{
			TotalMemory: types.Float64Value(memoryResource.GetTotalMemory()),
			UsedMemory:  types.Float64Value(memoryResource.GetUsedMemory()),
		},
		Disk: &ResourceDisk{
			TotalDiskSize: types.Float64Value(diskResource.GetTotalDiskSize()),
			UsedDiskSize:  types.Float64Value(diskResource.GetUsedDiskSize()),
		},
	}
}

func newResourceObject(r obcloudsdk.ResourceDoV2) basetypes.ObjectValue {
	cpuResource := r.GetCpu()
	memoryResource := r.GetMemory()
	diskResource := r.GetDisk()
	cpuAttributeTypes := map[string]attr.Type{"total_cpu": basetypes.Float64Type{}, "used_cpu": basetypes.Float64Type{}}
	memoryAttributeTypes := map[string]attr.Type{"total_memory": basetypes.Float64Type{}, "used_memory": basetypes.Float64Type{}}
	diskAttributeTypes := map[string]attr.Type{"total_disk_size": basetypes.Float64Type{}, "used_disk_size": basetypes.Float64Type{}}
	cpuAttributes := map[string]attr.Value{"total_cpu": types.Float64Value(cpuResource.GetTotalCpu()), "used_cpu": types.Float64Value(cpuResource.GetUsedCpu())}
	memoryAttributes := map[string]attr.Value{"total_memory": types.Float64Value(memoryResource.GetTotalMemory()), "used_memory": types.Float64Value(memoryResource.GetUsedMemory())}
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
