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
	"sync"
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
	body, err = clusterCreateRequest(&cluster)
	if err != nil {
		t.Fatalf("explicit node_num should be accepted: %v", err)
	}
	raw, err = json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "nodeNum") {
		t.Fatalf("create request unexpectedly contains nodeNum: %s", raw)
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
	for _, tt := range []struct {
		name         string
		nodeNum      types.Int32
		instanceType string
		createError  bool
		updateError  bool
	}{
		{"omitted", types.Int32Null(), "CLUSTER", false, false},
		{"unknown", types.Int32Unknown(), "CLUSTER", false, false},
		{"minimum", types.Int32Value(3), "CLUSTER", false, false},
		{"maximum", types.Int32Value(60), "CLUSTER", false, false},
		{"too small", types.Int32Value(0), "CLUSTER", true, true},
		{"not divisible", types.Int32Value(4), "CLUSTER", true, true},
		{"too large", types.Int32Value(63), "CLUSTER", true, true},
		{"unsupported type already at target", types.Int32Value(6), "SHARED_STORAGE_CLUSTER", true, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			state := tfsdk.State{Schema: schemaResp.Schema}
			state = clusterTestState(t, state, model.Cluster{ClusterSpec: model.ClusterSpec{Zones: []types.String{}, InstanceType: types.StringValue(tt.instanceType)}, ClusterInfo: model.ClusterInfo{NodeNum: tt.nodeNum}})
			beforeCreate := state
			beforeCreate.RemoveResource(context.Background())
			req := resource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw},
				Plan:   tfsdk.Plan{Schema: state.Schema, Raw: state.Raw},
				State:  beforeCreate,
			}
			for _, prior := range []struct {
				state     tfsdk.State
				wantError bool
			}{{beforeCreate, tt.createError}, {state, tt.updateError}} {
				req.State = prior.state
				var resp resource.ModifyPlanResponse
				r.ModifyPlan(context.Background(), req, &resp)
				if resp.Diagnostics.HasError() != prior.wantError {
					t.Fatalf("diagnostics = %v, want error = %v", resp.Diagnostics, prior.wantError)
				}
			}
		})
	}
}

func TestClusterCreateConvergesNodeNumber(t *testing.T) {
	for _, tt := range []struct {
		name         string
		target       types.Int32
		instanceType string
		pendingFirst bool
		holdScale    bool
		failScale    bool
		missingNodes bool
		wantCalls    string
		wantNodes    int32
		wantError    bool
	}{
		{name: "omitted uses remote", target: types.Int32Null(), wantCalls: "create", wantNodes: 3},
		{name: "already at target", target: types.Int32Value(3), wantCalls: "create", wantNodes: 3},
		{name: "scale to target after online", target: types.Int32Value(6), pendingFirst: true, wantCalls: "create,nodeNum", wantNodes: 6},
		{name: "accepted scale unconfirmed", target: types.Int32Value(6), holdScale: true, wantCalls: "create,nodeNum", wantNodes: 3, wantError: true},
		{name: "scale failed", target: types.Int32Value(6), failScale: true, wantCalls: "create,nodeNum", wantNodes: 3, wantError: true},
		{name: "missing remote node count", target: types.Int32Value(6), missingNodes: true, wantCalls: "create", wantError: true},
		{name: "omitted and missing remote node count", target: types.Int32Null(), missingNodes: true, wantCalls: "create", wantError: true},
		{name: "invalid before create", target: types.Int32Value(4), wantError: true},
		{name: "unsupported before create", target: types.Int32Value(6), instanceType: "SHARED_STORAGE_CLUSTER", wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var mu sync.Mutex
			var calls []string
			nodes := int32(3)
			onlineRead := false
			pending := tt.pendingFirst
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				mu.Lock()
				defer mu.Unlock()
				w.Header().Set("Content-Type", "application/json")
				switch req.Method {
				case http.MethodPost:
					calls = append(calls, "create")
					var body map[string]any
					if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
						t.Errorf("decode create request: %v", err)
					}
					if _, ok := body["nodeNum"]; ok {
						t.Error("create request must not contain nodeNum")
					}
					fmt.Fprint(w, `{"success":true,"data":{"instanceId":"new-cluster"}}`)
				case http.MethodPut:
					calls = append(calls, "nodeNum")
					if !onlineRead || req.URL.Path != "/api/v2/instances/new-cluster/nodeNum" {
						t.Errorf("scale before online read or wrong endpoint: %s", req.URL.Path)
					}
					var body struct {
						NodeNum int32 `json:"nodeNum"`
					}
					if err := json.NewDecoder(req.Body).Decode(&body); err != nil || body.NodeNum != 6 {
						t.Errorf("scale body = %+v, error = %v", body, err)
					}
					if tt.failScale {
						w.WriteHeader(http.StatusInternalServerError)
						fmt.Fprint(w, `{"success":false}`)
						return
					}
					if !tt.holdScale {
						nodes = body.NodeNum
					}
					fmt.Fprint(w, `{"success":true}`)
				case http.MethodGet:
					status := "ONLINE"
					if pending {
						status = "PENDING"
					}
					item := map[string]any{
						"instanceId": "new-cluster", "projectId": "project", "instanceType": "CLUSTER",
						"status": status, "diskSize": 100 * nodes / 3,
						"instanceName": "new", "instanceClass": "4C16G", "region": "cn-hangzhou",
						"availableZones": []string{"a"}, "cloudProvider": "ALIYUN", "version": "4.3", "replicaMode": "3",
					}
					if !tt.missingNodes {
						item["nodeNum"] = nodes
					}
					if req.URL.Path == "/api/v2/instances/new-cluster" {
						_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": item})
					} else {
						onlineRead = status == "ONLINE"
						pending = false
						_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": map[string]any{"total": 1, "dataList": []any{item}}})
					}
				default:
					t.Errorf("unexpected method %s", req.Method)
				}
			}))
			defer server.Close()
			r, state := clusterTestResource(t, server)
			instanceType := tt.instanceType
			if instanceType == "" {
				instanceType = "CLUSTER"
			}
			cluster := model.Cluster{ClusterSpec: model.ClusterSpec{
				CloudProvider: types.StringValue("ALIYUN"), DiskSize: types.Int32Value(100), Region: types.StringValue("cn-hangzhou"),
				Zones: []types.String{types.StringValue("a")}, InstanceClass: types.StringValue("4C16G"), ProjectId: types.StringValue("project"),
				Version: types.StringValue("4.3"), ReplicaMode: types.StringValue("3"), InstanceType: types.StringValue(instanceType),
			}, ClusterInfo: model.ClusterInfo{NodeNum: tt.target}}
			state = clusterTestState(t, state, cluster)
			resp := resource.CreateResponse{State: state}
			ctx := context.Background()
			if tt.holdScale {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, 200*time.Millisecond)
				defer cancel()
			}
			r.Create(ctx, resource.CreateRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}, &resp)
			mu.Lock()
			gotCalls := strings.Join(calls, ",")
			mu.Unlock()
			if gotCalls != tt.wantCalls || resp.Diagnostics.HasError() != tt.wantError {
				t.Fatalf("calls = %s, diagnostics = %v", gotCalls, resp.Diagnostics)
			}
			if tt.wantCalls != "" {
				var got model.Cluster
				if diags := resp.State.Get(context.Background(), &got); diags.HasError() || got.Id.ValueString() != "new-cluster" || got.NodeNum.ValueInt32() != tt.wantNodes || tt.missingNodes && !got.NodeNum.IsNull() {
					t.Fatalf("state = %+v, diagnostics = %v", got, diags)
				}
			}
		})
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
	}, ClusterInfo: model.ClusterInfo{NodeNum: types.Int32Value(3), Name: types.StringValue("old")}}
	for _, tt := range []struct {
		name  string
		edit  func(*model.Cluster)
		want  string
		error bool
	}{
		{"unknown stopped", func(p *model.Cluster) { p.IsStopped = types.BoolUnknown() }, "", false},
		{"stop", func(p *model.Cluster) { p.IsStopped = types.BoolValue(true) }, OperationStopCluster, false},
		{"node num", func(p *model.Cluster) { p.NodeNum = types.Int32Value(6) }, OperationModifyNodeNum, false},
		{"disk size", func(p *model.Cluster) { p.DiskSize = types.Int32Value(200) }, OperationModifyInstanceSpec, false},
		{"disk type", func(p *model.Cluster) { p.DiskType = types.StringValue("other") }, "", true},
		{"disk type case", func(p *model.Cluster) { p.DiskType = types.StringValue("CLOUD_ESSD_PL1") }, "", false},
		{"storage architecture", func(p *model.Cluster) { p.StorageArchitecture = types.StringValue("shareStorage") }, "", true},
		{"payment type", func(p *model.Cluster) { p.PaymentType = types.StringValue("PrePay") }, "", true},
		{"payment type case", func(p *model.Cluster) { p.PaymentType = types.StringValue("POSTPAY") }, "", false},
		{"version", func(p *model.Cluster) { p.Version = types.StringValue("4.4") }, "", true},
		{"zones reorder", func(p *model.Cluster) { p.Zones = []types.String{types.StringValue("b"), types.StringValue("a")} }, "", false},
		{"zones change", func(p *model.Cluster) { p.Zones = []types.String{types.StringValue("b"), types.StringValue("c")} }, "", true},
		{"unsupported checked first", func(p *model.Cluster) { p.Version = types.StringValue("4.4"); p.NodeNum = types.Int32Value(6) }, "", true},
		{"multiple operations", func(p *model.Cluster) {
			p.DiskSize = types.Int32Value(200)
			p.NodeNum = types.Int32Value(6)
			p.Name = types.StringValue("new")
			p.IsStopped = types.BoolValue(true)
		}, strings.Join([]string{OperationModifyInstanceSpec, OperationModifyNodeNum, OperationModifyInstanceName, OperationStopCluster}, ","), false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			plan := base
			tt.edit(&plan)
			got, err := getClusterOperations(&base, &plan)
			if strings.Join(got, ",") != tt.want || (err != nil) != tt.error {
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

type clusterUpdateFixture struct {
	mu            sync.Mutex
	t             *testing.T
	status        string
	name          string
	instanceClass string
	diskSize      int32
	nodeNum       int32
	calls         []string
	observed      bool
	fail          string
	failApplied   bool
	holdName      bool
}

func (f *clusterUpdateFixture) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	if req.Method == http.MethodGet {
		item := map[string]any{
			"instanceId": "target", "projectId": "project", "instanceType": "CLUSTER",
			"status": f.status, "instanceName": f.name, "instanceClass": f.instanceClass,
			"nodeNum": f.nodeNum, "diskSize": f.diskSize * f.nodeNum / 3,
			"region": "cn-hangzhou", "availableZones": []string{"a", "b"},
			"cloudProvider": "ALIYUN", "version": "4.3", "replicaMode": "3",
			"payType": "PostPay", "diskType": "cloud_essd_pl1", "storageArchitecture": "standardStorage",
		}
		if req.URL.Path == "/api/v2/instances/target" {
			_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": item})
		} else {
			f.observed = true
			_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": map[string]any{"total": 1, "dataList": []any{item}}})
		}
		return
	}
	operation := req.URL.Path[strings.LastIndex(req.URL.Path, "/")+1:]
	if len(f.calls) > 0 && !f.observed {
		f.t.Errorf("operation %s issued before the preceding operation was read", operation)
	}
	f.calls = append(f.calls, operation)
	f.observed = false
	apply := func() {
		switch operation {
		case "startCluster":
			f.status = "ONLINE"
		case "spec":
			f.instanceClass, f.diskSize = "8C32G", 200
		case "nodeNum":
			f.nodeNum = 6
		case "instanceName":
			if !f.holdName {
				f.name = "new"
			}
		case "stopCluster":
			if f.name != "new" {
				f.t.Errorf("stop issued before rename was confirmed")
			}
			f.status = "STOPPED"
		default:
			f.t.Errorf("unexpected operation %s", operation)
		}
	}
	if f.fail == operation {
		f.fail = ""
		if f.failApplied {
			apply()
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"success":false}`)
		return
	}
	apply()
	fmt.Fprint(w, `{"success":true}`)
}

func (f *clusterUpdateFixture) callList() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.calls...)
}

func clusterUpdateModel(stopped bool) model.Cluster {
	status := "ONLINE"
	if stopped {
		status = "STOPPED"
	}
	return model.Cluster{ClusterSpec: model.ClusterSpec{
		CloudProvider: types.StringValue("ALIYUN"), Region: types.StringValue("cn-hangzhou"),
		Zones:         []types.String{types.StringValue("cn-hangzhou-a"), types.StringValue("cn-hangzhou-b")},
		InstanceClass: types.StringValue("4C16G"), DiskSize: types.Int32Value(100),
		Version: types.StringValue("4.3"), ReplicaMode: types.StringValue("3"), InstanceType: types.StringValue("CLUSTER"),
		ProjectId: types.StringValue("project"), PaymentType: types.StringValue("PostPay"),
		DiskType: types.StringValue("cloud_essd_pl1"), StorageArchitecture: types.StringValue("standardStorage"),
		IsStopped: types.BoolValue(stopped),
	}, ClusterInfo: model.ClusterInfo{
		Id: types.StringValue("target"), Name: types.StringValue("old"), NodeNum: types.Int32Value(3), Status: types.StringValue(status),
	}}
}

func clusterUpdate(t *testing.T, ctx context.Context, r *clusterResource, state tfsdk.State, plan model.Cluster) resource.UpdateResponse {
	t.Helper()
	target := clusterTestState(t, state, plan)
	resp := resource.UpdateResponse{State: state}
	r.Update(ctx, resource.UpdateRequest{State: state, Plan: tfsdk.Plan{Schema: target.Schema, Raw: target.Raw}}, &resp)
	return resp
}

func TestClusterSequentialUpdates(t *testing.T) {
	for _, tt := range []struct {
		name    string
		stopped bool
		want    string
	}{
		{"start before changes", true, "startCluster,spec,nodeNum,instanceName"},
		{"stop after changes", false, "spec,nodeNum,instanceName,stopCluster"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			status := "ONLINE"
			if tt.stopped {
				status = "STOPPED"
			}
			fixture := &clusterUpdateFixture{t: t, status: status, name: "old", instanceClass: "4C16G", diskSize: 100, nodeNum: 3}
			server := httptest.NewServer(fixture)
			defer server.Close()
			r, state := clusterTestResource(t, server)
			base := clusterUpdateModel(tt.stopped)
			state = clusterTestState(t, state, base)
			plan := base
			plan.IsStopped = types.BoolValue(!tt.stopped)
			plan.InstanceClass, plan.DiskSize = types.StringValue("8C32G"), types.Int32Value(200)
			plan.NodeNum, plan.Name = types.Int32Value(6), types.StringValue("new")
			resp := clusterUpdate(t, context.Background(), r, state, plan)
			if resp.Diagnostics.HasError() {
				t.Fatalf("update failed: %v", resp.Diagnostics)
			}
			if got := strings.Join(fixture.callList(), ","); got != tt.want {
				t.Fatalf("calls = %s, want %s", got, tt.want)
			}
			var got model.Cluster
			if diags := resp.State.Get(context.Background(), &got); diags.HasError() || !got.IsStopped.Equal(plan.IsStopped) || !got.Name.Equal(plan.Name) || !got.InstanceClass.Equal(plan.InstanceClass) || !got.DiskSize.Equal(plan.DiskSize) || !got.NodeNum.Equal(plan.NodeNum) {
				t.Fatalf("updated state = %+v, diagnostics = %v", got, diags)
			}
		})
	}
}

func TestClusterUpdateFailureAndRetry(t *testing.T) {
	fixture := &clusterUpdateFixture{t: t, status: "ONLINE", name: "old", instanceClass: "4C16G", diskSize: 100, nodeNum: 3, fail: "nodeNum"}
	server := httptest.NewServer(fixture)
	defer server.Close()
	r, state := clusterTestResource(t, server)
	base := clusterUpdateModel(false)
	state = clusterTestState(t, state, base)
	plan := base
	plan.InstanceClass, plan.DiskSize = types.StringValue("8C32G"), types.Int32Value(200)
	plan.NodeNum, plan.Name, plan.IsStopped = types.Int32Value(6), types.StringValue("new"), types.BoolValue(true)
	failed := clusterUpdate(t, context.Background(), r, state, plan)
	var partial model.Cluster
	if diags := failed.State.Get(context.Background(), &partial); !failed.Diagnostics.HasError() || diags.HasError() || !partial.InstanceClass.Equal(plan.InstanceClass) || !partial.DiskSize.Equal(plan.DiskSize) || !partial.NodeNum.Equal(base.NodeNum) || partial.Id.ValueString() != "target" {
		t.Fatalf("partial state = %+v, update diagnostics = %v, state diagnostics = %v", partial, failed.Diagnostics, diags)
	}
	completed := clusterUpdate(t, context.Background(), r, failed.State, plan)
	if completed.Diagnostics.HasError() {
		t.Fatalf("retry failed: %v", completed.Diagnostics)
	}
	if got, want := strings.Join(fixture.callList(), ","), "spec,nodeNum,nodeNum,instanceName,stopCluster"; got != want {
		t.Fatalf("calls = %s, want %s", got, want)
	}
}

func TestClusterUpdateAmbiguousFailureRetainsConfirmedTarget(t *testing.T) {
	fixture := &clusterUpdateFixture{t: t, status: "ONLINE", name: "old", instanceClass: "4C16G", diskSize: 100, nodeNum: 3, fail: "instanceName", failApplied: true}
	server := httptest.NewServer(fixture)
	defer server.Close()
	r, state := clusterTestResource(t, server)
	base := clusterUpdateModel(false)
	state = clusterTestState(t, state, base)
	plan := base
	plan.Name = types.StringValue("new")
	failed := clusterUpdate(t, context.Background(), r, state, plan)
	var current model.Cluster
	if diags := failed.State.Get(context.Background(), &current); !failed.Diagnostics.HasError() || diags.HasError() || !current.Name.Equal(plan.Name) {
		t.Fatalf("confirmed state = %+v, update diagnostics = %v, state diagnostics = %v", current, failed.Diagnostics, diags)
	}
	completed := clusterUpdate(t, context.Background(), r, failed.State, plan)
	if calls := fixture.callList(); completed.Diagnostics.HasError() || len(calls) != 1 {
		t.Fatalf("confirmed rename should not be repeated: calls %v, diagnostics %v", calls, completed.Diagnostics)
	}
}

func TestClusterUpdateWaitsForRename(t *testing.T) {
	fixture := &clusterUpdateFixture{t: t, status: "ONLINE", name: "old", instanceClass: "4C16G", diskSize: 100, nodeNum: 3, holdName: true}
	server := httptest.NewServer(fixture)
	defer server.Close()
	r, state := clusterTestResource(t, server)
	base := clusterUpdateModel(false)
	state = clusterTestState(t, state, base)
	plan := base
	plan.Name, plan.IsStopped = types.StringValue("new"), types.BoolValue(true)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	resp := clusterUpdate(t, ctx, r, state, plan)
	var current model.Cluster
	if diags := resp.State.Get(context.Background(), &current); !resp.Diagnostics.HasError() || diags.HasError() || !current.Name.Equal(base.Name) || len(fixture.callList()) != 1 {
		t.Fatalf("unconfirmed rename must stop the sequence: state %+v, calls %v, diagnostics %v", current, fixture.callList(), resp.Diagnostics)
	}
}

func TestClusterUpdatePreflightAndPowerState(t *testing.T) {
	for _, tt := range []struct {
		name      string
		stopped   bool
		edit      func(*model.Cluster)
		wantCalls string
		wantError bool
	}{
		{"unsupported before mutation", false, func(p *model.Cluster) {
			p.Version = types.StringValue("4.4")
			p.DiskSize = types.Int32Value(200)
		}, "", true},
		{"invalid node count before mutation", false, func(p *model.Cluster) {
			p.NodeNum = types.Int32Value(4)
			p.DiskSize = types.Int32Value(200)
		}, "", true},
		{"unknown power state", false, func(p *model.Cluster) {
			p.IsStopped = types.BoolUnknown()
			p.InstanceClass, p.DiskSize = types.StringValue("8C32G"), types.Int32Value(200)
			p.NodeNum = types.Int32Value(6)
		}, "spec,nodeNum", false},
		{"stopped without explicit start", true, func(p *model.Cluster) {
			p.InstanceClass, p.DiskSize = types.StringValue("8C32G"), types.Int32Value(200)
		}, "spec", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			status := "ONLINE"
			if tt.stopped {
				status = "STOPPED"
			}
			fixture := &clusterUpdateFixture{t: t, status: status, name: "old", instanceClass: "4C16G", diskSize: 100, nodeNum: 3}
			server := httptest.NewServer(fixture)
			defer server.Close()
			r, state := clusterTestResource(t, server)
			base := clusterUpdateModel(tt.stopped)
			state = clusterTestState(t, state, base)
			plan := base
			tt.edit(&plan)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			if tt.stopped {
				cancel()
				ctx, cancel = context.WithTimeout(context.Background(), 200*time.Millisecond)
			}
			defer cancel()
			resp := clusterUpdate(t, ctx, r, state, plan)
			if calls := strings.Join(fixture.callList(), ","); calls != tt.wantCalls || resp.Diagnostics.HasError() != tt.wantError {
				t.Fatalf("calls = %s, diagnostics = %v", calls, resp.Diagnostics)
			}
		})
	}
}

func TestClusterRefreshTypes(t *testing.T) {
	for _, tt := range []struct {
		name, detailType, listType string
		wantError                  bool
	}{
		{"flagship", "FLAG_SHIP", "flag_ship", false},
		{"detail tenant", "MYSQL_TENANT", "FLAG_SHIP", true},
		{"list tenant", "FLAG_SHIP", "MYSQL_TENANT", true},
		{"detail unknown", "UNKNOWN_CLUSTER", "FLAG_SHIP", true},
		{"list unknown", "FLAG_SHIP", "UNKNOWN_CLUSTER", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if req.URL.Path == "/api/v2/instances/target" {
					fmt.Fprintf(w, `{"success":true,"data":{"instanceId":"target","instanceType":%q,"status":"ONLINE"}}`, tt.detailType)
				} else {
					fmt.Fprintf(w, `{"success":true,"data":{"total":1,"dataList":[{"instanceId":"target","instanceType":%q,"status":"ONLINE"}]}}`, tt.listType)
				}
			}))
			defer server.Close()
			cluster := model.Cluster{ClusterInfo: model.ClusterInfo{Id: types.StringValue("target")}}
			err := refreshCluster(context.Background(), &cluster, clusterTestClient(server))
			if (err != nil) != tt.wantError {
				t.Fatalf("refresh error = %v, want error = %t", err, tt.wantError)
			}
		})
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
