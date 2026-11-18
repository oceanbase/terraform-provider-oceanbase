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
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

func clusterCreateRequest(plan *model.Cluster) (*obcloudsdk.MultiCloudCreateInstanceRequest, error) {
	if len(plan.Zones) == 0 || plan.Region.IsNull() || plan.Region.IsUnknown() {
		return nil, fmt.Errorf("region and at least one zone are required")
	}
	if !plan.NodeNum.IsNull() && !plan.NodeNum.IsUnknown() {
		return nil, fmt.Errorf("node_num cannot be specified during creation; set it in a subsequent update")
	}
	if !plan.Series.IsNull() && !plan.Series.IsUnknown() {
		return nil, fmt.Errorf("series cannot be specified during creation")
	}
	if !plan.CpuArch.IsNull() && !plan.CpuArch.IsUnknown() {
		return nil, fmt.Errorf("cpu_arch cannot be specified during creation")
	}
	if !plan.IsStopped.IsNull() && !plan.IsStopped.IsUnknown() && plan.IsStopped.ValueBool() {
		return nil, fmt.Errorf("is_stopped cannot be true during creation; stop the cluster in a subsequent update")
	}

	body := obcloudsdk.NewMultiCloudCreateInstanceRequestWithDefaults()
	body.SetCloudProvider(plan.CloudProvider.ValueString())
	body.SetDiskSize(plan.DiskSize.ValueInt32())
	body.SetRegion(plan.Region.ValueString())
	body.SetObVersion(plan.Version.ValueString())
	zones := make([]string, 0, len(plan.Zones))
	for _, zone := range plan.Zones {
		zones = append(zones, zone.ValueString())
	}
	body.SetZones(strings.Join(zones, ","))
	body.SetInstanceClass(plan.InstanceClass.ValueString())
	body.SetReplicaMode(plan.ReplicaMode.ValueString())
	body.SetInstanceType(plan.InstanceType.ValueString())
	if !plan.SaleChannel.IsNull() && !plan.SaleChannel.IsUnknown() {
		body.SetSaleChannel(plan.SaleChannel.ValueString())
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		body.SetInstanceName(plan.Name.ValueString())
	}
	if !plan.PaymentType.IsNull() && !plan.PaymentType.IsUnknown() {
		body.SetChargeType(plan.PaymentType.ValueString())
	}
	if !plan.DiskType.IsNull() && !plan.DiskType.IsUnknown() {
		body.SetDiskType(plan.DiskType.ValueString())
	}
	if !plan.StorageArchitecture.IsNull() && !plan.StorageArchitecture.IsUnknown() {
		body.SetStorageArchitecture(plan.StorageArchitecture.ValueString())
	}
	return body, nil
}

func sameClusterZones(old, next []types.String) bool {
	if len(old) != len(next) {
		return false
	}
	zones := make(map[string]bool, len(old))
	for _, zone := range old {
		if zone.IsNull() || zone.IsUnknown() {
			return false
		}
		zones[zone.ValueString()] = true
	}
	for _, zone := range next {
		if zone.IsNull() || zone.IsUnknown() || !zones[zone.ValueString()] {
			return false
		}
	}
	return true
}

func clusterUpdateComplete(operation string, current, plan *model.Cluster) bool {
	switch operation {
	case OperationStopCluster:
		return isClusterStopped(current)
	case OperationStartCluster:
		return isClusterOnline(current)
	case OperationModifyNodeNum:
		return isClusterOnline(current) && current.NodeNum.Equal(plan.NodeNum)
	case OperationModifyInstanceSpec:
		return isClusterOnline(current) && current.InstanceClass.Equal(plan.InstanceClass) && current.DiskSize.Equal(plan.DiskSize)
	default:
		return false
	}
}

// findClusterInProject enumerates the authorized project inventory without the
// API's fuzzy instanceId filter. Absence is returned only after all pages.
func findClusterInProject(ctx context.Context, client *obcloudsdk.APIClient, projectID, instanceID string) (*obcloudsdk.DescribeInstancesResponseV2OpenAPI, error) {
	if projectID == "" {
		return nil, fmt.Errorf("project ID and cluster ID are required to enumerate the project")
	}
	return findClusterInInventory(ctx, client, projectID, instanceID, false)
}

// The API's instanceId filter is a substring match, so still compare exact IDs.
func findClusterByID(ctx context.Context, client *obcloudsdk.APIClient, projectID, instanceID string) (*obcloudsdk.DescribeInstancesResponseV2OpenAPI, error) {
	return findClusterInInventory(ctx, client, projectID, instanceID, true)
}

func findClusterInInventory(ctx context.Context, client *obcloudsdk.APIClient, projectID, instanceID string, filter bool) (*obcloudsdk.DescribeInstancesResponseV2OpenAPI, error) {
	if instanceID == "" {
		return nil, fmt.Errorf("cluster ID is required to query the inventory")
	}
	const pageSize int32 = 100
	var total int64 = -1
	seen := make(map[string]bool)
	for page := int32(1); ; page++ {
		request := client.MultiCloudOpenAPI.DescribeInstances(ctx).RequestId("").PageNumber(page).PageSize(pageSize)
		if projectID != "" {
			request = request.XObProjectId(projectID)
		}
		if filter {
			request = request.InstanceId(instanceID)
		}
		result, response, err := request.Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || !result.GetSuccess() || result.Data == nil || result.Data.Total == nil {
			return nil, fmt.Errorf("cannot enumerate project %s instances: error %v, response %v, error code %s", projectID, err, response, result.GetErrorCode())
		}
		data := result.Data
		if *data.Total < 0 || (total >= 0 && total != *data.Total) {
			return nil, fmt.Errorf("project instance inventory changed while enumerating")
		}
		total = *data.Total
		for i := range data.DataList {
			item := &data.DataList[i]
			id := item.GetInstanceId()
			if id == "" || seen[id] {
				return nil, fmt.Errorf("project instance inventory contains a missing or duplicate ID")
			}
			seen[id] = true
			if id == instanceID {
				return item, nil
			}
		}
		if int64(len(seen)) == total {
			return nil, errResourceNotFound
		}
		if int64(len(seen)) > total || len(data.DataList) == 0 {
			return nil, fmt.Errorf("incomplete project instance inventory: received %d of %d", len(seen), total)
		}
	}
}
