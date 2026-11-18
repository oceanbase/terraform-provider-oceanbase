package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type clusterTestTransport struct {
	base   http.RoundTripper
	target *url.URL
}

func (t clusterTestTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	clone := request.Clone(request.Context())
	clone.URL.Scheme, clone.URL.Host, clone.Host = t.target.Scheme, t.target.Host, t.target.Host
	return t.base.RoundTrip(clone)
}

func clusterTestResource(t *testing.T, server *httptest.Server) (*clusterResource, tfsdk.State) {
	t.Helper()
	target, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	previous := http.DefaultTransport
	http.DefaultTransport = clusterTestTransport{base: server.Client().Transport, target: target}
	t.Cleanup(func() { http.DefaultTransport = previous })
	r := &clusterResource{provider: &oceanbaseProvider{configured: true, clientGenerator: &api.Generator{Site: api.SiteMock}}}
	var schemaResp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	return r, state
}

func clusterTestState(t *testing.T, state tfsdk.State, cluster model.Cluster) tfsdk.State {
	t.Helper()
	resourceType := state.Schema.GetAttributes()["resource"].GetType().(basetypes.ObjectType)
	cluster.Resource = types.ObjectNull(resourceType.AttrTypes)
	if diags := state.Set(context.Background(), cluster); diags.HasError() {
		t.Fatalf("failed to construct cluster state: %v", diags)
	}
	return state
}

func clusterTestClient(server *httptest.Server) *obcloudsdk.APIClient {
	cfg := obcloudsdk.NewConfiguration()
	cfg.Scheme, cfg.Host = "http", strings.TrimPrefix(server.URL, "http://")
	return obcloudsdk.NewAPIClient(cfg)
}

func TestClusterCreateRequest(t *testing.T) {
	cluster := model.Cluster{ClusterSpec: model.ClusterSpec{
		CloudProvider: types.StringValue("ALIYUN"), DiskSize: types.Int32Value(100),
		Region: types.StringValue("cn-hangzhou"), Zones: []types.String{types.StringValue("h")},
		Version: types.StringValue("4.3.5"), InstanceClass: types.StringValue("4C16G"),
		InstanceType: types.StringValue("CLUSTER"), ReplicaMode: types.StringValue("3"),
		PaymentType: types.StringValue("PostPay"), DiskType: types.StringValue("cloud_essd_pl1"),
		StorageArchitecture: types.StringValue("standardStorage"),
	}}
	body, err := clusterCreateRequest(&cluster)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]string{"chargeType": "PostPay", "diskType": "cloud_essd_pl1", "storageArchitecture": "standardStorage"} {
		if payload[key] != want {
			t.Errorf("%s = %v, want %s", key, payload[key], want)
		}
	}
	if _, ok := payload["nodeNum"]; ok {
		t.Fatal("create request unexpectedly contains nodeNum")
	}
	cluster.NodeNum = types.Int32Value(3)
	if _, err := clusterCreateRequest(&cluster); err == nil || !strings.Contains(err.Error(), "node_num") {
		t.Fatalf("explicit node_num should be rejected: %v", err)
	}
}

func TestClusterStorageFieldsPopulateWhenUnset(t *testing.T) {
	detail := obcloudsdk.NewDescribeInstanceResponseV2OpenAPI()
	detail.SetDiskType("cloud_essd_pl1")
	detail.SetStorageArchitecture("standardStorage")
	detail.SetPayType("POSTPAY")
	var cluster model.Cluster
	cluster.RefreshClusterFromInstance(detail)
	if cluster.DiskType.ValueString() != "cloud_essd_pl1" || cluster.StorageArchitecture.ValueString() != "standardStorage" || cluster.PaymentType.ValueString() != "POSTPAY" {
		t.Fatalf("detail values were not populated: %+v", cluster)
	}
}

func TestClusterCreateWaitFailureRetainsID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodPost:
			fmt.Fprint(w, `{"success":true,"data":{"instanceId":"new-cluster"}}`)
		case req.URL.Path == "/api/v2/instances/new-cluster":
			fmt.Fprint(w, `{"success":true,"data":{"instanceId":"new-cluster","projectId":"project","instanceType":"CLUSTER","status":"PENDING"}}`)
		default:
			fmt.Fprint(w, `{"success":true,"data":{"total":1,"dataList":[{"instanceId":"new-cluster","instanceType":"CLUSTER","status":"PENDING"}]}}`)
		}
	}))
	defer server.Close()
	r, state := clusterTestResource(t, server)
	cluster := model.Cluster{ClusterSpec: model.ClusterSpec{
		CloudProvider: types.StringValue("ALIYUN"), DiskSize: types.Int32Value(100), Region: types.StringValue("cn-hangzhou"),
		Zones: []types.String{types.StringValue("a")}, InstanceClass: types.StringValue("4C16G"),
		Version: types.StringValue("4.3"), ReplicaMode: types.StringValue("3"), InstanceType: types.StringValue("CLUSTER"),
	}}
	state = clusterTestState(t, state, cluster)
	req := resource.CreateRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}
	resp := resource.CreateResponse{State: state}
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	r.Create(ctx, req, &resp)
	if !resp.Diagnostics.HasError() {
		t.Fatal("creation should report a readiness timeout")
	}
	var got model.Cluster
	if diags := resp.State.Get(context.Background(), &got); diags.HasError() || got.Id.ValueString() != "new-cluster" {
		t.Fatalf("created ID was not retained: %s, diagnostics: %v", got.Id.ValueString(), diags)
	}
}

func TestClusterNodeNumberPlan(t *testing.T) {
	r := &clusterResource{}
	var schemaResp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	state = clusterTestState(t, state, model.Cluster{ClusterSpec: model.ClusterSpec{Zones: []types.String{}}, ClusterInfo: model.ClusterInfo{NodeNum: types.Int32Value(3)}})
	beforeCreate := state
	beforeCreate.RemoveResource(context.Background())
	req := resource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw},
		Plan:   tfsdk.Plan{Schema: state.Schema, Raw: state.Raw},
		State:  beforeCreate,
	}
	var createResp resource.ModifyPlanResponse
	r.ModifyPlan(context.Background(), req, &createResp)
	if !createResp.Diagnostics.HasError() {
		t.Fatal("a node count configured at creation must be rejected in the plan")
	}
	req.State = state
	var updateResp resource.ModifyPlanResponse
	r.ModifyPlan(context.Background(), req, &updateResp)
	if updateResp.Diagnostics.HasError() {
		t.Fatalf("node count must remain configurable for updates: %v", updateResp.Diagnostics)
	}
}

func TestClusterDeleteWaitsForInventory(t *testing.T) {
	for _, tt := range []struct {
		name, afterDelete string
		wantError         bool
	}{
		{"confirmed absent", `{"success":true,"data":{"total":0,"dataList":[]}}`, false},
		{"inventory denied", `{"success":false,"errorCode":"PRI_ACCESS_DENIED_EXCEPTION"}`, true},
		{"still present", `{"success":true,"data":{"total":1,"dataList":[{"instanceId":"target"}]}}`, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			accepted, checked := false, false
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case req.Method == http.MethodDelete:
					accepted = true
					fmt.Fprint(w, `{"success":true}`)
				case req.URL.Path == "/api/v2/instances/target":
					fmt.Fprint(w, `{"success":true,"data":{"instanceId":"target","projectId":"project","instanceType":"CLUSTER","status":"ONLINE"}}`)
				case accepted:
					checked = true
					if req.Header.Get("X-Ob-Project-Id") != "project" || req.URL.Query().Has("instanceId") {
						t.Errorf("delete confirmation must enumerate the full project: %s", req.URL)
					}
					fmt.Fprint(w, tt.afterDelete)
				default:
					fmt.Fprint(w, `{"success":true,"data":{"total":1,"dataList":[{"instanceId":"target","instanceType":"CLUSTER","status":"ONLINE"}]}}`)
				}
			}))
			defer server.Close()
			r, state := clusterTestResource(t, server)
			state = clusterTestState(t, state, model.Cluster{ClusterSpec: model.ClusterSpec{ProjectId: types.StringValue("project"), Zones: []types.String{}}, ClusterInfo: model.ClusterInfo{Id: types.StringValue("target")}})
			resp := resource.DeleteResponse{State: state}
			ctx := context.Background()
			if tt.name == "still present" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, 200*time.Millisecond)
				defer cancel()
			}
			r.Delete(ctx, resource.DeleteRequest{State: state}, &resp)
			if !accepted || !checked || resp.Diagnostics.HasError() != tt.wantError || resp.State.Raw.IsNull() == tt.wantError {
				t.Fatalf("accepted=%v checked=%v state null=%v diagnostics=%v", accepted, checked, resp.State.Raw.IsNull(), resp.Diagnostics)
			}
		})
	}
}

func TestClusterOperations(t *testing.T) {
	base := model.Cluster{ClusterSpec: model.ClusterSpec{
		CloudProvider: types.StringValue("ALIYUN"), Region: types.StringValue("cn-hangzhou"),
		Zones:         []types.String{types.StringValue("a"), types.StringValue("b")},
		InstanceClass: types.StringValue("4C16G"), DiskSize: types.Int32Value(100),
		Version: types.StringValue("4.3"), PaymentType: types.StringValue("PostPay"),
		StorageArchitecture: types.StringValue("standardStorage"), DiskType: types.StringValue("cloud_essd_pl1"),
		IsStopped: types.BoolValue(false),
	}, ClusterInfo: model.ClusterInfo{NodeNum: types.Int32Value(3)}}
	for _, tt := range []struct {
		name  string
		edit  func(*model.Cluster)
		want  string
		error bool
	}{
		{"unknown stopped", func(p *model.Cluster) { p.IsStopped = types.BoolUnknown() }, OperationNothing, false},
		{"stop", func(p *model.Cluster) { p.IsStopped = types.BoolValue(true) }, OperationStopCluster, false},
		{"node num", func(p *model.Cluster) { p.NodeNum = types.Int32Value(6) }, OperationModifyNodeNum, false},
		{"disk size", func(p *model.Cluster) { p.DiskSize = types.Int32Value(200) }, OperationModifyInstanceSpec, false},
		{"disk type", func(p *model.Cluster) { p.DiskType = types.StringValue("other") }, OperationUnsupported, true},
		{"disk type case", func(p *model.Cluster) { p.DiskType = types.StringValue("CLOUD_ESSD_PL1") }, OperationNothing, false},
		{"storage architecture", func(p *model.Cluster) { p.StorageArchitecture = types.StringValue("shareStorage") }, OperationUnsupported, true},
		{"payment type", func(p *model.Cluster) { p.PaymentType = types.StringValue("PrePay") }, OperationUnsupported, true},
		{"payment type case", func(p *model.Cluster) { p.PaymentType = types.StringValue("POSTPAY") }, OperationNothing, false},
		{"version", func(p *model.Cluster) { p.Version = types.StringValue("4.4") }, OperationUnsupported, true},
		{"zones reorder", func(p *model.Cluster) { p.Zones = []types.String{types.StringValue("b"), types.StringValue("a")} }, OperationNothing, false},
		{"zones change", func(p *model.Cluster) { p.Zones = []types.String{types.StringValue("b"), types.StringValue("c")} }, OperationUnsupported, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			plan := base
			tt.edit(&plan)
			got, err := getClusterOperation(context.Background(), &base, &plan)
			if got != tt.want || (err != nil) != tt.error {
				t.Fatalf("operation = %s, error = %v", got, err)
			}
		})
	}
	plan := base
	plan.NodeNum = types.Int32Value(6)
	current := base
	current.Status = types.StringValue(model.InstanceStatusOnline)
	if clusterUpdateComplete(OperationModifyNodeNum, &current, &plan) {
		t.Fatal("online status alone must not complete a node update")
	}
	current.NodeNum = plan.NodeNum
	if !clusterUpdateComplete(OperationModifyNodeNum, &current, &plan) {
		t.Fatal("target node count and online status should complete the update")
	}
}

func TestClusterInventoryExactPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Ob-Project-Id") != "project" || r.URL.Query().Has("instanceId") {
			t.Errorf("inventory must use project scope without a fuzzy ID filter: %s", r.URL)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("pageNumber") == "1" {
			fmt.Fprint(w, `{"success":true,"data":{"total":2,"dataList":[{"instanceId":"target-other"}]}}`)
		} else {
			fmt.Fprint(w, `{"success":true,"data":{"total":2,"dataList":[{"instanceId":"target"}]}}`)
		}
	}))
	defer server.Close()
	got, err := findClusterInProject(context.Background(), clusterTestClient(server), "project", "target")
	if err != nil || got.GetInstanceId() != "target" {
		t.Fatalf("exact lookup = %v, %v", got, err)
	}
	_, err = findClusterInProject(context.Background(), clusterTestClient(server), "project", "absent")
	if !errors.Is(err, errResourceNotFound) {
		t.Fatalf("complete inventory should establish absence: %v", err)
	}
}

func TestClusterFilteredRefreshWithoutProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/v2/instances/target" {
			fmt.Fprint(w, `{"success":true,"data":{"instanceId":"target","instanceType":"CLUSTER","status":"ONLINE"}}`)
			return
		}
		if r.URL.Query().Get("instanceId") != "target" || r.Header.Get("X-Ob-Project-Id") != "" {
			t.Errorf("expected an ID-filtered lookup without a project header: %s", r.URL)
		}
		fmt.Fprint(w, `{"success":true,"data":{"total":2,"dataList":[{"instanceId":"target-other","instanceType":"CLUSTER"},{"instanceId":"target","instanceType":"CLUSTER","status":"ONLINE"}]}}`)
	}))
	defer server.Close()
	cluster := model.Cluster{ClusterInfo: model.ClusterInfo{Id: types.StringValue("target")}}
	if err := refreshCluster(context.Background(), &cluster, clusterTestClient(server)); err != nil {
		t.Fatalf("ordinary refresh without project ID failed: %v", err)
	}
}

func TestClusterRefreshEvidence(t *testing.T) {
	for _, tt := range []struct {
		name              string
		detailStatus      int
		detail, inventory string
		missing, failure  bool
	}{
		{"forbidden", http.StatusForbidden, ``, `{"success":true,"data":{"total":0,"dataList":[]}}`, false, true},
		{"business denied", http.StatusOK, `{"success":false,"errorCode":"PRI_ACCESS_DENIED_EXCEPTION"}`, `{"success":true,"data":{"total":0,"dataList":[]}}`, false, true},
		{"missing confirmed", http.StatusNotFound, ``, `{"success":true,"data":{"total":0,"dataList":[]}}`, true, false},
		{"business missing confirmed", http.StatusOK, `{"success":false,"errorCode":"Instance.NotExists"}`, `{"success":true,"data":{"total":0,"dataList":[]}}`, true, false},
		{"list denied", http.StatusNotFound, ``, `{"success":false,"errorCode":"PRI_ACCESS_DENIED_EXCEPTION"}`, false, true},
		{"incomplete inventory", http.StatusNotFound, ``, `{"success":true,"data":{"total":2,"dataList":[]}}`, false, true},
		{"present", http.StatusOK, `{"success":true,"data":{"instanceId":"target","projectId":"project","instanceType":"CLUSTER","status":"ONLINE","payType":"POSTPAY","diskType":"cloud_essd_pl1","storageArchitecture":"standardStorage","nodeNum":3}}`, `{"success":true,"data":{"total":1,"dataList":[{"instanceId":"target","instanceType":"CLUSTER","status":"ONLINE","payType":"POSTPAY","diskType":"cloud_essd_pl1","storageArchitecture":"standardStorage","diskSize":100}]}}`, false, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.URL.Path == "/api/v2/instances/target" {
					w.WriteHeader(tt.detailStatus)
					fmt.Fprint(w, tt.detail)
				} else {
					fmt.Fprint(w, tt.inventory)
				}
			}))
			defer server.Close()
			cluster := model.Cluster{ClusterSpec: model.ClusterSpec{ProjectId: types.StringValue("project"), PaymentType: types.StringValue("PostPay"), DiskType: types.StringValue("CLOUD_ESSD_PL1"), StorageArchitecture: types.StringValue("STANDARDSTORAGE")}, ClusterInfo: model.ClusterInfo{Id: types.StringValue("target")}}
			err := refreshCluster(context.Background(), &cluster, clusterTestClient(server))
			if errors.Is(err, errResourceNotFound) != tt.missing || (err != nil) != (tt.missing || tt.failure) {
				t.Fatalf("unexpected refresh error: %v", err)
			}
			if tt.name == "present" && (cluster.DiskType.ValueString() != "CLOUD_ESSD_PL1" || cluster.StorageArchitecture.ValueString() != "STANDARDSTORAGE") {
				t.Fatalf("storage fields were not refreshed: %+v", cluster)
			}
			if tt.name == "present" && cluster.PaymentType.ValueString() != "PostPay" {
				t.Fatalf("payment type lost configured casing: %s", cluster.PaymentType.ValueString())
			}
		})
	}
}
