package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
)

func privateLinkTestProvider(t *testing.T, handler http.HandlerFunc) *oceanbaseProvider {
	t.Helper()
	previous := http.DefaultTransport
	http.DefaultTransport = connectionRoundTripper(func(req *http.Request) (*http.Response, error) {
		recorder := httptest.NewRecorder()
		handler(recorder, req)
		return recorder.Result(), nil
	})
	t.Cleanup(func() { http.DefaultTransport = previous })
	return &oceanbaseProvider{configured: true, clientGenerator: &api.Generator{Site: api.SiteMock}}
}

func privateLinkServiceState(t *testing.T, r *privateLinkServiceResource, value privateLinkServiceModel) tfsdk.State {
	t.Helper()
	var schemaResp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(context.Background(), &value); diags.HasError() {
		t.Fatalf("set service state: %v", diags)
	}
	return state
}

func privateLinkConnectionState(t *testing.T, r *privateLinkConnectionResource, value privateLinkConnectionModel) tfsdk.State {
	t.Helper()
	var schemaResp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(context.Background(), &value); diags.HasError() {
		t.Fatalf("set connection state: %v", diags)
	}
	return state
}

func serviceTestModel() privateLinkServiceModel {
	return privateLinkServiceModel{
		ID: types.StringNull(), ClusterID: types.StringValue("cluster"), TenantID: types.StringValue("tenant"),
		ProjectID: types.StringValue("project"), UserAccount: types.StringValue("account"),
		ProxySetting: types.ObjectNull(privateLinkProxySettingTypes()),
		ServiceName:  types.StringNull(), ServiceAccount: types.StringNull(), Status: types.StringNull(),
	}
}

func privateLinkProxySettingTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mode": types.StringType, "deployment_architecture": types.StringType,
		"zone": types.ListType{ElemType: types.StringType}, "slave_zone": types.ListType{ElemType: types.StringType},
		"node_cpu": types.Float64Type,
	}
}

func TestPrivateLinkServiceForwardsProxySetting(t *testing.T) {
	t.Setenv(EnvResourceCreateTimeoutSeconds, "5")
	created := false
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodGet:
			if !created {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
				return
			}
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"SERVICE_CREATED","privateLinkServiceId":"service-name","privateLinkUserAccount":"account"}]}]}`)
		case req.Method == http.MethodPost && strings.HasSuffix(req.URL.Path, "/privatelink/service"):
			var body struct {
				UserAccount  string `json:"userAccount"`
				ProxySetting struct {
					Mode                   string   `json:"mode"`
					DeploymentArchitecture string   `json:"deploymentArchitecture"`
					Zone                   []string `json:"zone"`
					SlaveZone              []string `json:"slaveZone"`
					NodeCPU                float64  `json:"nodeCpu"`
				} `json:"proxySetting"`
			}
			if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			setting := body.ProxySetting
			if body.UserAccount != "account" || setting.Mode != "k8s_shared" || setting.DeploymentArchitecture != "dual_az_active_active_deployment" || len(setting.Zone) != 2 || setting.Zone[0] != "zone-a" || setting.Zone[1] != "zone-b" || setting.SlaveZone == nil || len(setting.SlaveZone) != 0 || setting.NodeCPU != 1 {
				t.Errorf("unexpected service body: %+v", body)
			}
			created = true
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"address"}}`)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})
	model := serviceTestModel()
	model.ProxySetting = types.ObjectValueMust(privateLinkProxySettingTypes(), map[string]attr.Value{
		"mode": types.StringValue("k8s_shared"), "deployment_architecture": types.StringValue("dual_az_active_active_deployment"),
		"zone":       types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zone-a"), types.StringValue("zone-b")}),
		"slave_zone": types.ListNull(types.StringType), "node_cpu": types.Float64Value(1),
	})
	r := &privateLinkServiceResource{provider: p}
	state := privateLinkServiceState(t, r, model)
	resp := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Raw: state.Raw, Schema: state.Schema}}, &resp)
	if resp.Diagnostics.HasError() || !created {
		t.Fatalf("service creation failed: %v", resp.Diagnostics)
	}
}

func connectionTestModel() privateLinkConnectionModel {
	return privateLinkConnectionModel{
		ID: types.StringNull(), ClusterID: types.StringValue("cluster"), TenantID: types.StringValue("tenant"),
		ProjectID: types.StringValue("project"), AddressID: types.StringValue("address"), EndpointID: types.StringValue("endpoint"),
		UserEndpointAddress: types.StringNull(), EndpointAddress: types.StringNull(), Domain: types.StringNull(), Status: types.StringNull(),
	}
}

func TestPrivateLinkLifecycleUsesOpenAPI(t *testing.T) {
	t.Setenv(EnvResourceCreateTimeoutSeconds, "5")
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "5")
	status, endpoint := "UNCREATED", ""
	var calls []string
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if got := req.Header.Get("X-Ob-Project-Id"); got != "project" {
			t.Errorf("project header = %q", got)
		}
		base := "/api/v2/instances/cluster/tenants/tenant/privatelink"
		calls = append(calls, req.Method+" "+strings.TrimPrefix(req.URL.Path, base))
		switch {
		case req.Method == http.MethodGet && req.URL.Path == strings.TrimSuffix(base, "/privatelink"):
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
		case req.Method == http.MethodGet && req.URL.Path == base:
			if status == "UNCREATED" {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"UNCREATED"}]}]}`)
				return
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":%q,"privateLinkServiceId":"service-name","privateLinkServiceAccount":"service-owner","privateLinkUserAccount":"account","privateLinkEndpointId":%q,"privateLinkDomain":"private.example"}]}]}`, status, endpoint)
		case req.Method == http.MethodPost && req.URL.Path == base+"/service":
			var body map[string]string
			_ = json.NewDecoder(req.Body).Decode(&body)
			if body["userAccount"] != "account" {
				t.Errorf("service account = %q", body["userAccount"])
			}
			status = "SERVICE_CREATED"
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"address"}}`)
		case req.Method == http.MethodPost && req.URL.Path == base+"/connection":
			var body map[string]string
			_ = json.NewDecoder(req.Body).Decode(&body)
			if body["addressId"] != "address" || body["endpointId"] != "endpoint" {
				t.Errorf("incorrect connection body: %v", body)
			}
			endpoint, status = "endpoint", "WAIT_CONFIRMATION"
			_, _ = fmt.Fprint(w, `{"success":true,"data":{}}`)
		case req.Method == http.MethodPost && req.URL.Path == base+"/confirmation":
			status = "CONNECTED"
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		case req.Method == http.MethodDelete && req.URL.Path == base+"/endpoint":
			if req.URL.Query().Get("addressId") != "address" || req.URL.Query().Get("endpointId") != "endpoint" {
				t.Errorf("incorrect endpoint deletion query: %s", req.URL.RawQuery)
			}
			status, endpoint = "UNCREATED", ""
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})

	svc := &privateLinkServiceResource{provider: p}
	svcState := privateLinkServiceState(t, svc, serviceTestModel())
	svcCreate := resource.CreateResponse{State: svcState}
	svc.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Raw: svcState.Raw, Schema: svcState.Schema}}, &svcCreate)
	var service privateLinkServiceModel
	if svcCreate.Diagnostics.HasError() || svcCreate.State.Get(context.Background(), &service).HasError() || service.ID.ValueString() != "address" || service.ServiceName.ValueString() != "service-name" || service.ServiceAccount.ValueString() != "service-owner" {
		t.Fatalf("service creation failed: state=%+v diagnostics=%v", service, svcCreate.Diagnostics)
	}

	connection := &privateLinkConnectionResource{provider: p}
	connState := privateLinkConnectionState(t, connection, connectionTestModel())
	connCreate := resource.CreateResponse{State: connState}
	connection.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Raw: connState.Raw, Schema: connState.Schema}}, &connCreate)
	var connected privateLinkConnectionModel
	if connCreate.Diagnostics.HasError() || connCreate.State.Get(context.Background(), &connected).HasError() || connected.ID.ValueString() != "endpoint" || connected.Status.ValueString() != "CONNECTED" {
		t.Fatalf("connection creation failed: state=%+v diagnostics=%v", connected, connCreate.Diagnostics)
	}

	read := resource.ReadResponse{State: connCreate.State}
	connection.Read(context.Background(), resource.ReadRequest{State: connCreate.State}, &read)
	if read.Diagnostics.HasError() {
		t.Fatalf("connection refresh failed: %v", read.Diagnostics)
	}
	del := resource.DeleteResponse{}
	connection.Delete(context.Background(), resource.DeleteRequest{State: read.State}, &del)
	if del.Diagnostics.HasError() {
		t.Fatalf("connection deletion failed: %v", del.Diagnostics)
	}
	svcDelete := resource.DeleteResponse{}
	svc.Delete(context.Background(), resource.DeleteRequest{State: svcCreate.State}, &svcDelete)
	if svcDelete.Diagnostics.HasError() {
		t.Fatalf("service reconciliation failed: %v", svcDelete.Diagnostics)
	}
	confirmCount, deleteCount := 0, 0
	for _, call := range calls {
		if call == "POST /confirmation" {
			confirmCount++
		}
		if call == "DELETE /endpoint" {
			deleteCount++
		}
	}
	if confirmCount != 1 || deleteCount != 1 {
		t.Fatalf("confirmation or deletion repeated: %v", calls)
	}
}

func TestPrivateLinkDeleteDoesNotDeleteAnotherEndpoint(t *testing.T) {
	deletes := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		if req.Method == http.MethodDelete {
			deletes++
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"CONNECTED","privateLinkEndpointId":"another","privateLinkUserAccount":"account","privateLinkServiceId":"service-name"}]}]}`)
	})
	r := &privateLinkConnectionResource{provider: p}
	state := connectionTestModel()
	state.ID, state.Status = types.StringValue("endpoint"), types.StringValue("CONNECTED")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkConnectionState(t, r, state)}, &resp)
	if resp.Diagnostics.HasError() || deletes != 0 {
		t.Fatalf("missing endpoint must not cause a delete: deletes=%d diagnostics=%v", deletes, resp.Diagnostics)
	}
}

func TestPrivateLinkDeleteWaitsForTenant(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "5")
	deleted, tenantReads := false, 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
			tenantReads++
			status := "DELETING_MASTER_ADDRESS"
			if tenantReads > 1 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
		case req.Method == http.MethodDelete:
			deleted = true
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		case deleted:
			_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
		default:
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"CONNECTED","privateLinkEndpointId":"endpoint"}]}]}`)
		}
	})
	r := &privateLinkConnectionResource{provider: p}
	state := connectionTestModel()
	state.ID, state.Status = types.StringValue("endpoint"), types.StringValue("CONNECTED")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkConnectionState(t, r, state)}, &resp)
	if resp.Diagnostics.HasError() || !deleted || tenantReads < 2 {
		t.Fatalf("deletion did not wait for tenant: deleted=%v tenant reads=%d diagnostics=%v", deleted, tenantReads, resp.Diagnostics)
	}
}

func TestPrivateLinkServiceDoesNotDeleteOtherEndpoints(t *testing.T) {
	deletes := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodDelete {
			deletes++
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"CONNECTED","privateLinkEndpointId":"another","privateLinkUserAccount":"account","privateLinkServiceId":"service-name"}]}]}`)
	})
	r := &privateLinkServiceResource{provider: p}
	state := serviceTestModel()
	state.ID, state.Status = types.StringValue("address"), types.StringValue("CONNECTED")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkServiceState(t, r, state)}, &resp)
	if !resp.Diagnostics.HasError() || deletes != 0 {
		t.Fatalf("service with an endpoint must not be deleted: deletes=%d diagnostics=%v", deletes, resp.Diagnostics)
	}
}

func TestPrivateLinkServiceDeletesUnconnectedService(t *testing.T) {
	status, deletes := "SERVICE_CREATED", 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		if req.Method == http.MethodDelete {
			deletes++
			if !strings.HasSuffix(req.URL.Path, "/privatelink/endpoint") || req.URL.Query().Get("status") != "SERVICE_CREATED" || req.URL.Query().Get("endpointId") != "" {
				t.Errorf("incorrect service cleanup request: %s", req.URL)
			}
			status = "UNCREATED"
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
			return
		}
		_, _ = fmt.Fprintf(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":%q,"privateLinkUserAccount":"account","privateLinkServiceId":"service-name"}]}]}`, status)
	})
	r := &privateLinkServiceResource{provider: p}
	state := serviceTestModel()
	state.ID, state.Status = types.StringValue("address"), types.StringValue("SERVICE_CREATED")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkServiceState(t, r, state)}, &resp)
	if resp.Diagnostics.HasError() || deletes != 1 {
		t.Fatalf("unconnected service cleanup failed: deletes=%d diagnostics=%v", deletes, resp.Diagnostics)
	}
}

func TestPrivateLinkCreateTimeoutRetainsServiceIdentity(t *testing.T) {
	t.Setenv(EnvResourceCreateTimeoutSeconds, "1")
	created := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodPost {
			created++
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"address"}}`)
			return
		}
		if created == 0 {
			_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"SERVICE_CREATING","privateLinkUserAccount":"account"}]}]}`)
	})
	r := &privateLinkServiceResource{provider: p}
	state := privateLinkServiceState(t, r, serviceTestModel())
	resp := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Raw: state.Raw, Schema: state.Schema}}, &resp)
	var saved privateLinkServiceModel
	if resp.State.Get(context.Background(), &saved).HasError() || !resp.Diagnostics.HasError() || saved.ID.ValueString() != "address" || created != 1 {
		t.Fatalf("timeout did not preserve service identity: id=%q creates=%d diagnostics=%v", saved.ID.ValueString(), created, resp.Diagnostics)
	}
}

func TestPrivateLinkDeleteUnconfirmedEndpointKeepsState(t *testing.T) {
	deletes := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodDelete {
			deletes++
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
	})
	r := &privateLinkConnectionResource{provider: p}
	state := connectionTestModel()
	state.ID, state.Status = types.StringValue("endpoint"), types.StringValue("CONNECTING")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkConnectionState(t, r, state)}, &resp)
	if !resp.Diagnostics.HasError() || deletes != 0 {
		t.Fatalf("unknown endpoint outcome must not be removed: deletes=%d diagnostics=%v", deletes, resp.Diagnostics)
	}
}

func TestPrivateLinkDeleteUnconfirmedServiceKeepsState(t *testing.T) {
	deletes := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodDelete {
			deletes++
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
	})
	r := &privateLinkServiceResource{provider: p}
	state := serviceTestModel()
	state.ID, state.Status = types.StringValue("address"), types.StringValue("SERVICE_CREATING")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkServiceState(t, r, state)}, &resp)
	if !resp.Diagnostics.HasError() || deletes != 0 {
		t.Fatalf("unknown service outcome must not be removed: deletes=%d diagnostics=%v", deletes, resp.Diagnostics)
	}
}

func TestPrivateLinkDeleteRejectsMissingEndpointIdentity(t *testing.T) {
	requests := 0
	p := privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		requests++
	})
	r := &privateLinkConnectionResource{provider: p}
	state := connectionTestModel()
	state.ID, state.EndpointID, state.Status = types.StringValue("endpoint"), types.StringNull(), types.StringValue("CONNECTED")
	resp := resource.DeleteResponse{}
	r.Delete(context.Background(), resource.DeleteRequest{State: privateLinkConnectionState(t, r, state)}, &resp)
	if !resp.Diagnostics.HasError() || requests != 0 {
		t.Fatalf("invalid endpoint identity must not call the API: requests=%d diagnostics=%v", requests, resp.Diagnostics)
	}
}

func TestPrivateLinkTerraformLifecycle(t *testing.T) {
	var lock sync.Mutex
	status, endpoint := "UNCREATED", ""
	privateLinkTestProvider(t, func(w http.ResponseWriter, req *http.Request) {
		lock.Lock()
		defer lock.Unlock()
		w.Header().Set("Content-Type", "application/json")
		base := "/api/v2/instances/cluster/tenants/tenant/privatelink"
		switch {
		case req.Method == http.MethodGet && req.URL.Path == strings.TrimSuffix(base, "/privatelink"):
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
		case req.Method == http.MethodGet && req.URL.Path == base:
			if status == "UNCREATED" {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":"UNCREATED"}]}]}`)
				return
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":[{"addressId":"address","privateLinkList":[{"privateLinkStatus":%q,"privateLinkServiceId":"service-name","privateLinkServiceAccount":"service-owner","privateLinkUserAccount":"account","privateLinkEndpointId":%q}]}]}`, status, endpoint)
		case req.Method == http.MethodPost && req.URL.Path == base+"/service":
			status = "SERVICE_CREATED"
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"address"}}`)
		case req.Method == http.MethodPost && req.URL.Path == base+"/connection":
			status, endpoint = "WAIT_CONFIRMATION", "endpoint"
			_, _ = fmt.Fprint(w, `{"success":true,"data":{}}`)
		case req.Method == http.MethodPost && req.URL.Path == base+"/confirmation":
			status = "CONNECTED"
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		case req.Method == http.MethodDelete && req.URL.Path == base+"/endpoint":
			status, endpoint = "UNCREATED", ""
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})
	config := providerConfig + `
resource "oceanbase_tenant_private_link_service" "test" {
  cluster_id   = "cluster"
  tenant_id    = "tenant"
  user_account = "account"
}
resource "oceanbase_tenant_private_link_connection" "test" {
  cluster_id  = "cluster"
  tenant_id   = "tenant"
  address_id  = oceanbase_tenant_private_link_service.test.id
  endpoint_id = "endpoint"
}`
	tfresource.UnitTest(t, tfresource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []tfresource.TestStep{
			{Config: config, Check: tfresource.ComposeAggregateTestCheckFunc(
				tfresource.TestCheckResourceAttr("oceanbase_tenant_private_link_service.test", "service_name", "service-name"),
				tfresource.TestCheckResourceAttr("oceanbase_tenant_private_link_service.test", "service_account", "service-owner"),
				tfresource.TestCheckResourceAttr("oceanbase_tenant_private_link_connection.test", "status", "CONNECTED"),
			)},
			{Config: config, PlanOnly: true},
			{
				ResourceName:            "oceanbase_tenant_private_link_service.test",
				ImportState:             true,
				ImportStateId:           "cluster/tenant/address/account",
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id", "status"},
			},
			{
				ResourceName:            "oceanbase_tenant_private_link_connection.test",
				ImportState:             true,
				ImportStateId:           "cluster/tenant/address/endpoint",
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id", "user_endpoint_address"},
			},
		},
	})
}
