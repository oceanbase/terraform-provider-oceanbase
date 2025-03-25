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
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &clusterResource{}
	_ resource.ResourceWithConfigure   = &clusterResource{}
	_ resource.ResourceWithImportState = &clusterResource{}
)

type clusterResource struct {
	provider *oceanbaseProvider
}

func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

func (r *clusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Cluster Resource",
		Attributes: map[string]schema.Attribute{
			// required fields to create cluster
			"cloud_provider": schema.StringAttribute{
				MarkdownDescription: "Cloud provider to create the cluster",
				Required:            true,
			},
			"disk_size": schema.Int32Attribute{
				MarkdownDescription: "Disk size of observer",
				Required:            true,
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region id of the cluster",
				Required:            true,
			},
			"zones": schema.SetAttribute{
				MarkdownDescription: "Zones of the cluster",
				Required:            true,
				ElementType:         types.StringType,
			},
			"instance_class": schema.StringAttribute{
				MarkdownDescription: "Instance class of the cluster, e.g. 4C16G",
				Required:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Version of the cluster",
				Required:            true,
			},
			"replica_mode": schema.StringAttribute{
				MarkdownDescription: "Replica mode of the cluster",
				Required:            true,
			},
			"instance_type": schema.StringAttribute{
				MarkdownDescription: "Business scenario of cluster, e.g. CLUSTER, ANALYTICAL_CLUSTER, KV_CLUSTER, SHARED_STORAGE_CLUSTER",
				Required:            true,
			},

			// optional fields to create cluster
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the cluster",
				Optional:            true,
				Computed:            true,
			},
			"sale_channel": schema.StringAttribute{
				MarkdownDescription: "Channel of the order to create the cluster",
				Optional:            true,
				Computed:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Id of the project, you can find project_id under your account",
				Optional:            true,
				Computed:            true,
			},

			// not available to be set at create time currently
			"node_num": schema.Int32Attribute{
				MarkdownDescription: "Total node number of the cluster, it's calculated by summing up the observer numbers in each zone",
				Optional:            true,
				Computed:            true,
			},
			"series": schema.StringAttribute{
				MarkdownDescription: "Series of the cluster, e.g. normal, normal_kv",
				Optional:            true,
				Computed:            true,
			},
			"cpu_arch": schema.StringAttribute{
				MarkdownDescription: "Cpu arch of the cluster",
				Optional:            true,
				Computed:            true,
			},
			"payment_type": schema.StringAttribute{
				MarkdownDescription: "Payment type of the cluster",
				Optional:            true,
				Computed:            true,
			},
			"is_stopped": schema.BoolAttribute{
				MarkdownDescription: "Whether the cluster is supposed to be stopped",
				Optional:            true,
				Computed:            true,
			},
			// computed fields, known after creation
			"id": schema.StringAttribute{
				MarkdownDescription: "Id of the cluster",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
			"create_time": schema.StringAttribute{
				MarkdownDescription: "Create time of cluster",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status of cluster, such as ONLINE, PENDING etc",
				Computed:            true,
			},
			"vpc": schema.StringAttribute{
				MarkdownDescription: "VPC where the cluster is created",
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
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_cpu": schema.Float64Attribute{
								MarkdownDescription: "Used cpu of the cluster",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
					"disk": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Disk resource info of the cluster",
						Attributes: map[string]schema.Attribute{
							"total_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Total disk size of the cluster",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_disk_size": schema.Float64Attribute{
								MarkdownDescription: "Used disk size of the cluster",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
					"memory": schema.SingleNestedAttribute{
						Computed:            true,
						MarkdownDescription: "Memory resource info of the cluster",
						Attributes: map[string]schema.Attribute{
							"total_memory": schema.Float64Attribute{
								MarkdownDescription: "Total memory of the cluster",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
							"used_memory": schema.Float64Attribute{
								MarkdownDescription: "Used memory of the cluster",
								PlanModifiers: []planmodifier.Float64{
									float64planmodifier.UseStateForUnknown(),
								},
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *clusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error",
			fmt.Sprintf("Error in Configure: expected %T but got %T", oceanbaseProvider{}, req.ProviderData))
	}
}

func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create cluster is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var plan model.Cluster

	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: add more checks
	if len(plan.Zones) == 0 {
		resp.Diagnostics.AddError("Zone is empty", "No zone is configured")
		return
	}
	if plan.Region.IsNull() {
		resp.Diagnostics.AddError("Unknown region", "Region is not allowed to be null")
		return
	}

	createInstanceBody := obcloudsdk.NewMultiCloudCreateInstanceRequestWithDefaults()
	createInstanceBody.SetCloudProvider(plan.CloudProvider.ValueString())
	createInstanceBody.SetDiskSize(plan.DiskSize.ValueInt32())
	createInstanceBody.SetRegion(plan.Region.ValueString())
	createInstanceBody.SetObVersion(plan.Version.ValueString())
	zoneList := make([]string, 0, len(plan.Zones))
	for _, zone := range plan.Zones {
		zoneList = append(zoneList, zone.ValueString())
	}
	zonesParam := strings.Join(zoneList, ",")
	createInstanceBody.SetZones(zonesParam)
	createInstanceBody.SetInstanceClass(plan.InstanceClass.ValueString())
	createInstanceBody.SetReplicaMode(plan.ReplicaMode.ValueString())
	createInstanceBody.SetInstanceType(plan.InstanceType.ValueString())
	// set optional values
	if !(plan.SaleChannel.IsNull() || plan.SaleChannel.IsUnknown()) {
		createInstanceBody.SetSaleChannel(plan.SaleChannel.ValueString())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		createInstanceBody.SetInstanceName(plan.Name.ValueString())
	}

	client := r.provider.clientGenerator.GetClient()
	result, response, err := client.MultiCloudOpenAPI.CreateInstance(ctx).RequestId("").Body(*createInstanceBody).XObProjectId(plan.ProjectId.ValueString()).Execute()
	api.LogAPIResult(ctx, result, response, err)

	// TODO: add retry policy according to result and response
	if err != nil || response.StatusCode != http.StatusOK || !(result.GetSuccess()) {
		resp.Diagnostics.AddError(
			"Error creating cluster",
			fmt.Sprintf("Could not create cluster: %v, response: %v", err, response),
		)
		return
	}

	state := plan
	createdInstance := result.GetData()
	state.Id = types.StringValue(createdInstance.GetInstanceId())
	diags = resp.State.Set(ctx, state)

	ctx = tflog.SetField(ctx, "instance_id", state.Id)
	tflog.Info(ctx, "Successfully created oceanbase cluster")

	timeoutDuration := DefaultClusterCreateTimeout
	envTimeoutSeconds := os.Getenv(EnvClusterCreateTimeoutSeconds)
	if envTimeoutSeconds != "" {
		timeoutSeconds, err := strconv.ParseInt(envTimeoutSeconds, 10, 64)
		if err == nil {
			timeoutDuration = time.Duration(timeoutSeconds) * time.Second
		}
	}
	err = retry.RetryContext(ctx, timeoutDuration,
		waitForClusterReadyFunc(ctx, &state, client, isClusterOnline))
	// Just add an error log here but consider the resource is created successfully
	// state is already refreshed while waiting, no need to refresh again here
	if err != nil {
		tflog.Error(ctx, "Wait for cluster online timed out, just save current state")
	}

	tflog.Debug(ctx, "Save obcluster to state")
	diags = resp.State.Set(ctx, state)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddWarning("Error after create the cluster", "Set the state error")
		return
	}
}

func (r *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Read cluster is called")
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}

	var cluster model.Cluster
	diags := req.State.Get(ctx, &cluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	client := r.provider.clientGenerator.GetClient()
	err := refreshCluster(ctx, &cluster, client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Refresh cluster info failed",
			fmt.Sprintf("Failed to refresh cluster info, err: %v", err))
		return
	}
	diags = resp.State.Set(ctx, cluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func getClusterOperation(ctx context.Context, state, plan *model.Cluster) (string, error) {
	// TODO: Is it necessary to check cluster status is ONLINE to perform operations
	modifiedSpecs := make([]string, 0)
	// modify node number
	if !(plan.NodeNum.IsUnknown() || plan.NodeNum.IsNull()) && !plan.NodeNum.Equal(state.NodeNum) {
		modifiedSpecs = append(modifiedSpecs, SpecNodeNum)
	}
	// modify cluster spec
	if !(plan.InstanceClass.IsUnknown() || plan.InstanceClass.IsNull()) && (!plan.InstanceClass.Equal(state.InstanceClass) || !plan.DiskSize.Equal(state.DiskSize)) {
		modifiedSpecs = append(modifiedSpecs, SpecInstanceSpec)
	}
	// start/stop cluster
	if !(plan.InstanceClass.IsUnknown() || plan.InstanceClass.IsNull()) && !plan.IsStopped.Equal(state.IsStopped) {
		modifiedSpecs = append(modifiedSpecs, SpecIsStopped)
	}
	// modify instance name
	if !(plan.Name.IsUnknown() || plan.Name.IsNull()) && !plan.Name.Equal(state.Name) {
		modifiedSpecs = append(modifiedSpecs, SpecInstanceName)
	}
	if len(modifiedSpecs) > 1 {
		return OperationUnsupported, errors.New(fmt.Sprintf("Modify %s at once is not supported", strings.Join(modifiedSpecs, ", ")))
	}
	if len(modifiedSpecs) == 1 {
		switch modifiedSpecs[0] {
		case SpecNodeNum:
			return OperationModifyNodeNum, nil
		case SpecInstanceSpec:
			return OperationModifyInstanceSpec, nil
		case SpecInstanceName:
			return OperationModifyInstanceName, nil
		case SpecIsStopped:
			if plan.IsStopped.ValueBool() {
				return OperationStopCluster, nil
			} else {
				return OperationStartCluster, nil
			}
		default:
			return OperationUnsupported, errors.New("Unsupported operation")
		}
	}
	return OperationNothing, nil
}

func (r *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update cluster is called")

	var plan model.Cluster
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state model.Cluster
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	operation, err := getClusterOperation(ctx, &state, &plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Update is not supported",
			fmt.Sprintf("Got error when get operation err: %v", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Got operation: %s", operation))

	client := r.provider.clientGenerator.GetClient()
	switch operation {
	case OperationModifyNodeNum:
		tflog.Info(ctx, "Modify cluster node_num")
		result, response, err := client.MultiCloudOpenAPI.ModifyInstanceNodeNum(ctx, plan.Id.ValueString()).Body(obcloudsdk.MultiCloudModifyInstanceNodeNumRequest{
			NodeNum: plan.NodeNum.ValueInt32(),
		}).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify cluster node_num failed")
			resp.Diagnostics.AddError(
				"Modify cluster node_num failed",
				fmt.Sprintf("Got error when update cluster node_num err: %v, response: %v", err, response))
			return
		}
	case OperationModifyInstanceSpec:
		tflog.Info(ctx, "Modify cluster instance spec")
		result, response, err := client.MultiCloudOpenAPI.ModifyInstanceSpec(ctx, plan.Id.ValueString()).Body(obcloudsdk.MultiCloudModifyInstanceSpecRequest{
			InstanceClass: plan.InstanceClass.ValueStringPointer(),
			DiskSize:      plan.DiskSize.ValueInt32Pointer(),
		}).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify cluster spec failed")
			resp.Diagnostics.AddError(
				"Modify cluster spec failed",
				fmt.Sprintf("Got error when update cluster spec err: %v, response: %v", err, response))
			return
		}
	case OperationModifyInstanceName:
		tflog.Info(ctx, "Modify cluster instance name")
		result, response, err := client.MultiCloudOpenAPI.ModifyInstanceName(ctx, plan.Id.ValueString()).Body(obcloudsdk.ModifyInstanceNameRequestV2{
			InstanceName: plan.Name.ValueStringPointer(),
		}).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Modify cluster name failed")
			resp.Diagnostics.AddError(
				"Modify cluster name failed",
				fmt.Sprintf("Got error when modify cluster name err: %v, response: %v", err, response))
			return
		}
		state.Name = plan.Name
	case OperationStartCluster:
		tflog.Info(ctx, "Start cluster")
		result, response, err := client.MultiCloudOpenAPI.StartCluster(ctx, plan.Id.ValueString()).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Start cluster spec failed")
			resp.Diagnostics.AddError(
				"Start cluster failed",
				fmt.Sprintf("Got error when start cluster err: %v, response: %v", err, response))
			return
		}
	case OperationStopCluster:
		tflog.Info(ctx, "Stop cluster")
		result, response, err := client.MultiCloudOpenAPI.StopCluster(ctx, plan.Id.ValueString()).RequestId("").Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
			tflog.Error(ctx, "Stop cluster spec failed")
			resp.Diagnostics.AddError(
				"Stop cluster failed",
				fmt.Sprintf("Got error when stop cluster err: %v, response: %v", err, response))
			return
		}
	case OperationNothing:
		tflog.Info(ctx, "Nothing to do")
	default:
		tflog.Info(ctx, "Unsupported operation")
	}

	if isAsyncOperation(operation) {
		timeoutDuration := DefaultClusterModifyTimeout
		envTimeoutSeconds := os.Getenv(EnvClusterModifyTimeoutSeconds)
		if envTimeoutSeconds != "" {
			timeoutSeconds, err := strconv.ParseInt(envTimeoutSeconds, 10, 64)
			if err == nil {
				timeoutDuration = time.Duration(timeoutSeconds) * time.Second
			}
		}

		tflog.Info(ctx, "Wait cluster turned into operating")
		err = retry.RetryContext(ctx, timeoutDuration,
			waitForClusterReadyFunc(ctx, &state, client, isClusterNotOnline))
		if err != nil {
			resp.Diagnostics.AddError(
				"Cluster operation failed after timeout",
				fmt.Sprintf("Cluster is not turned into operating after timed out: %v", err),
			)
			return
		}
		tflog.Info(ctx, "Cluster already turned into operating")

		tflog.Info(ctx, "Wait cluster turned into target status")
		validateFunction := isClusterOnline
		if operation == OperationStopCluster {
			validateFunction = isClusterStopped
		}
		err = retry.RetryContext(ctx, timeoutDuration,
			waitForClusterReadyFunc(ctx, &state, client, validateFunction))
		if err != nil {
			resp.Diagnostics.AddError(
				"Cluster operation failed after timeout",
				fmt.Sprintf("Cluster is not ready after timed out: %v", err),
			)
			return
		}
		tflog.Info(ctx, "Cluster already turned into online")
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func isAsyncOperation(operation string) bool {
	return (operation != OperationNothing && operation != OperationModifyInstanceName)
}

func (r *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete cluster is called")
	var cluster model.ClusterResource
	diags := req.State.Get(ctx, &cluster)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if cluster.Id.IsNull() || cluster.Id.IsUnknown() {
		resp.Diagnostics.AddError(
			"Cluster's id can't be null",
			"The id field is null, but it never should be. Please double check the value!",
		)
		return
	}

	client := r.provider.clientGenerator.GetClient()

	deleteInstanceRequest := obcloudsdk.NewDeleteInstanceOpenRequestWithDefaults()
	deleteInstanceRequest.SetInstanceId(cluster.Id.ValueString())
	result, response, err := client.MultiCloudOpenAPI.DeleteInstance(ctx).Body(*deleteInstanceRequest).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response.StatusCode != http.StatusOK || !(*result.Success) {
		resp.Diagnostics.AddError(
			"Failed to delete cluster",
			fmt.Sprintf("Unexpected error while delete cluster: %v, response: %v", err, response))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *clusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Import cluster is called")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func isClusterStopped(cluster *model.Cluster) bool {
	tflog.Debug(context.Background(), fmt.Sprintf("Cluster status when validate: %s", cluster.Status.ValueString()))
	return cluster.IsStopped.ValueBool()
}

func isClusterNotOnline(cluster *model.Cluster) bool {
	return !isClusterOnline(cluster)
}

func isClusterOnline(cluster *model.Cluster) bool {
	tflog.Debug(context.Background(), fmt.Sprintf("Cluster status when validate: %s", cluster.Status.ValueString()))
	return cluster.Status.ValueString() == model.InstanceStatusOnline
}

func waitForClusterReadyFunc(ctx context.Context, cluster *model.Cluster, client *obcloudsdk.APIClient, statusValidationFunction func(*model.Cluster) bool) retry.RetryFunc {
	return func() *retry.RetryError {
		err := refreshCluster(ctx, cluster, client)
		if err != nil {
			return retry.RetryableError(fmt.Errorf("Failed to refresh cluster info, err: %v", err))
		}
		if !statusValidationFunction(cluster) {
			return retry.RetryableError(fmt.Errorf("Cluster is still not ready, current status: %s", cluster.Status))
		}
		return nil
	}
}
