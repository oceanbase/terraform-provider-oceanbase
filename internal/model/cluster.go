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
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type ClusterSpec struct {
	CloudProvider types.String   `tfsdk:"cloud_provider"`
	DiskSize      types.Int32    `tfsdk:"disk_size"`
	Region        types.String   `tfsdk:"region"`
	Zones         []types.String `tfsdk:"zones"`
	InstanceClass types.String   `tfsdk:"instance_class"`
	Version       types.String   `tfsdk:"version"`
	ReplicaMode   types.String   `tfsdk:"replica_mode"`
	InstanceType  types.String   `tfsdk:"instance_type"`
	SaleChannel   types.String   `tfsdk:"sale_channel"`
	ProjectId     types.String   `tfsdk:"project_id"`
	IsStopped     types.Bool     `tfsdk:"is_stopped"`
	Series        types.String   `tfsdk:"series"`
	CpuArch       types.String   `tfsdk:"cpu_arch"`
	PaymentType   types.String   `tfsdk:"payment_type"`
}

type ClusterInfo struct {
	Id         types.String          `tfsdk:"id"`
	Name       types.String          `tfsdk:"name"`
	DeployMode types.String          `tfsdk:"deploy_mode"`
	DeployType types.String          `tfsdk:"deploy_type"`
	NodeNum    types.Int32           `tfsdk:"node_num"`
	CreateTime types.String          `tfsdk:"create_time"`
	Status     types.String          `tfsdk:"status"`
	Vpc        types.String          `tfsdk:"vpc"`
	StopTime   types.String          `tfsdk:"stop_time"`
	Resource   basetypes.ObjectValue `tfsdk:"resource"`
}

type Cluster struct {
	ClusterSpec
	ClusterInfo
}

type ClusterDataSource struct {
	Cluster
}

type ClusterResource struct {
	Cluster
}

func convertZoneNames(region string, zones []string) []types.String {
	regionParts := strings.Split(region, "-")
	_, parseErr := strconv.Atoi(regionParts[len(regionParts)-1])
	connector := ""
	if parseErr != nil {
		connector = "-"
	}
	zoneList := make([]types.String, len(zones))
	for idx, zone := range zones {
		zoneList[idx] = types.StringValue(fmt.Sprintf("%s%s%s", region, connector, zone))
	}
	return zoneList
}

func (c *Cluster) RefreshClusterFromInstance(data *obcloudsdk.DescribeInstanceResponseV2OpenAPI) {
	c.CloudProvider = types.StringValue(data.GetCloudProvider())
	c.CpuArch = types.StringValue(data.GetCpuArchitecture())
	c.CreateTime = types.StringValue(data.GetCreateTime())
	c.DeployMode = types.StringValue(data.GetDeployMode())
	c.DeployType = types.StringValue(data.GetDeployType())
	c.InstanceClass = types.StringValue(data.GetInstanceClass())
	c.InstanceType = types.StringValue(strings.ToUpper(data.GetInstanceType()))
	c.IsStopped = types.BoolValue(data.GetStatus() == InstanceStatusStopped)
	c.Name = types.StringValue(data.GetInstanceName())
	c.NodeNum = types.Int32Value(data.GetNodeNum())
	c.PaymentType = types.StringValue(data.GetPayType())
	c.ProjectId = types.StringValue(data.GetProjectId())
	c.Region = types.StringValue(data.GetRegion())
	c.ReplicaMode = types.StringValue(data.GetReplicaMode())
	c.Resource = newResourceObject(data.GetResource())
	c.SaleChannel = types.StringValue(data.GetSaleChannel())
	c.Status = types.StringValue(data.GetStatus())
	c.Series = types.StringValue(data.GetSeries())
	c.StopTime = types.StringValue(data.GetStopTime())
	c.Version = types.StringValue(data.GetVersion())
	c.Vpc = types.StringValue(data.GetVpcId())
	zones := convertZoneNames(c.Region.ValueString(), data.GetAvailableZones())
	// sort zones
	c.Zones = sortListByOriginalOrder(c.Zones, zones)
}

func (c *Cluster) RefreshClusterFromInstanceListItem(data *obcloudsdk.DescribeInstancesResponseV2OpenAPI) {
	c.CloudProvider = types.StringValue(data.GetCloudProvider())
	c.CpuArch = types.StringValue(data.GetCpuArchitecture())
	c.CreateTime = types.StringValue(data.GetCreateTime())
	c.DeployMode = types.StringValue(data.GetDeployMode())
	c.DeployType = types.StringValue(data.GetDeployType())
	c.DiskSize = types.Int32Value(int32(data.GetDiskSize()))
	// terraform's disk_size means disk size of each observer, it need to keep unchanged after adding server
	if !c.NodeNum.IsNull() && !c.NodeNum.IsUnknown() && c.NodeNum.ValueInt32() != 0 {
		c.DiskSize = types.Int32Value(int32(data.GetDiskSize() * 3 / int64(c.NodeNum.ValueInt32())))
	}
	c.InstanceType = types.StringValue(strings.ToUpper(data.GetInstanceType()))
	c.InstanceClass = types.StringValue(data.GetInstanceClass())
	c.IsStopped = types.BoolValue(data.GetStatus() == InstanceStatusStopped)
	c.Name = types.StringValue(data.GetInstanceName())
	c.PaymentType = types.StringValue(data.GetPayType())
	c.Region = types.StringValue(data.GetRegion())
	c.Resource = newResourceObject(data.GetResource())
	c.SaleChannel = types.StringValue(data.GetSaleChannel())
	c.Series = types.StringValue(data.GetSeries())
	c.Status = types.StringValue(data.GetStatus())
	c.StopTime = types.StringValue(data.GetStopTime())
	c.Version = types.StringValue(data.GetVersion())
	c.Vpc = types.StringValue(data.GetVpcId())
	c.Vpc = types.StringValue(data.GetVpcId())
	zones := convertZoneNames(c.Region.ValueString(), data.GetAvailableZones())
	// sort zones
	c.Zones = sortListByOriginalOrder(c.Zones, zones)
}
