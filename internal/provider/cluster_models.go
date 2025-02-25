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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClusterResourceData struct {
	CloudProvider  types.String   `tfsdk:"cloud_provider"`
	CreateTime     types.String   `tfsdk:"create_time"`
	DeployMode     types.String   `tfsdk:"deploy_mode"`
	StorageGib     types.Float64  `tfsdk:"storage_gib"`
	NodeSize       types.String   `tfsdk:"node_size"`
	ClusterId      types.String   `tfsdk:"id"`
	ClusterName    types.String   `tfsdk:"cluster_name"`
	ClusterType    types.String   `tfsdk:"cluster_type"`
	MaintainTime   types.String   `tfsdk:"maintain_time"`
	PayType        types.String   `tfsdk:"pay_type"`
	Region         types.String   `tfsdk:"region"`
	Resource       *Resource      `tfsdk:"resource"` // 为了处理null, 这里用指针
	Status         types.String   `tfsdk:"status"`
	Version        types.String   `tfsdk:"version"`
	AvailableZones []types.String `tfsdk:"available_zones"`
	TenantIds      types.List     `tfsdk:"tenant_ids"`
}

type Resource struct {
	Cpu    Cpu    `tfsdk:"cpu"`
	Disk   Disk   `tfsdk:"disk"`
	Memory Memory `tfsdk:"memory"`
}

type Cpu struct {
	TotalCpu types.Float64 `tfsdk:"total_cpu"`
}

type Disk struct {
	TotalDiskSize types.Float64 `tfsdk:"total_disk_size"`
}

type Memory struct {
	TotalMemory types.Float64 `tfsdk:"total_memory"`
}
