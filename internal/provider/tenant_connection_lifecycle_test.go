package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

func connectionModel() model.TenantConnection {
	return model.TenantConnection{
		Id: types.StringValue("managed"), ClusterId: types.StringValue("cluster"),
		TenantId: types.StringValue("tenant"), Type: types.StringValue("INTERNET"),
		ConnectionZones: types.ListNull(types.StringType),
		PrivateLink: types.ObjectNull(map[string]attr.Type{
			"user_account": types.StringType, "user_vpc": types.StringType, "user_vswitch": types.StringType,
		}),
		Intranet: types.ObjectNull(map[string]attr.Type{
			"domain": types.StringType, "status": types.StringType, "rpc_port": types.Int32Type, "peering_max_connections": types.Int32Type,
		}),
		Internet: types.ObjectNull(map[string]attr.Type{
			"domain": types.StringType, "status": types.StringType, "rpc_port": types.Int32Type,
			"max_connections": types.Int32Type, "address": types.StringType, "port": types.Int32Type,
		}),
	}
}

func connectionState(t *testing.T, r *tenantConnectionResource, value model.TenantConnection) tfsdk.State {
	t.Helper()
	var schemaResp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(context.Background(), &value); diags.HasError() {
		t.Fatalf("set state: %v", diags)
	}
	return state
}

type connectionRoundTripper func(*http.Request) (*http.Response, error)

func (f connectionRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func connectionResourceWithHandler(t *testing.T, handler http.HandlerFunc) *tenantConnectionResource {
	t.Helper()
	previous := http.DefaultTransport
	http.DefaultTransport = connectionRoundTripper(func(req *http.Request) (*http.Response, error) {
		recorder := httptest.NewRecorder()
		handler(recorder, req)
		return recorder.Result(), nil
	})
	t.Cleanup(func() { http.DefaultTransport = previous })
	return &tenantConnectionResource{provider: &oceanbaseProvider{
		configured: true, clientGenerator: &api.Generator{Site: api.SiteMock},
	}}
}

func TestConnectionRefreshIdentityAndMissing(t *testing.T) {
	for _, tt := range []struct {
		name, kind, response, publicResponse string
		status, publicStatus                 int
		missing, failure                     bool
	}{
		{name: "public exact id with decimal unit count", response: `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE","proxyClusterInfo":{"unitNum":1.0}}]}`},
		{name: "private exact id", kind: "INTRANET", response: `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`},
		{name: "empty list", response: `{"success":true,"data":[]}`, missing: true},
		{name: "known absent error", response: `{"success":false,"errorCode":"Tenant.Address.NotExist"}`, missing: true},
		{name: "different id", response: `{"success":true,"data":[{"addressId":"other","addressStatus":"ONLINE"}]}`, failure: true},
		{name: "different tenant", response: `{"success":true,"data":[{"addressId":"managed","tenantId":"other"}]}`, failure: true},
		{name: "different network", response: `{"success":true,"data":[{"addressId":"managed","networkType":"INTRANET"}]}`, failure: true},
		{name: "null entry and absent public address", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":{"useSSL":false}}`, missing: true},
		{name: "empty object entry", response: `{"success":true,"data":[{}]}`, failure: true},
		{name: "private null entry", kind: "INTRANET", response: `{"success":true,"data":[null]}`, failure: true},
		{name: "null entry and same public address", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":{"addressId":"managed"}}`, failure: true},
		{name: "null entry and other public address", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":{"addressId":"other"}}`, failure: true},
		{name: "null entry and no public data", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true}`, failure: true},
		{name: "null entry and null public data", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":null}`, failure: true},
		{name: "null entry and public access denied", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":false,"errorCode":"AccessDenied"}`, failure: true},
		{name: "null entry and public http forbidden", response: `{"success":true,"data":[null]}`, publicResponse: `{}`, publicStatus: http.StatusForbidden, failure: true},
		{name: "null entry and wrong public tenant", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":{"tenantId":"other"}}`, failure: true},
		{name: "null entry and public address without id", response: `{"success":true,"data":[null]}`, publicResponse: `{"success":true,"data":{"internetAddress":"192.0.2.1"}}`, failure: true},
		{name: "missing data", response: `{"success":true}`, failure: true},
		{name: "access denied", response: `{"success":false,"errorCode":"AccessDenied"}`, failure: true},
		{name: "http error", status: http.StatusForbidden, response: `{}`, failure: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			publicQueries := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if strings.HasSuffix(req.URL.Path, "/publicaddress") {
					publicQueries++
					if req.Method != http.MethodGet || tt.publicResponse == "" {
						t.Errorf("unexpected public address query: %s %s", req.Method, req.URL)
					}
					w.Header().Set("Content-Type", "application/json")
					if tt.publicStatus != 0 {
						w.WriteHeader(tt.publicStatus)
					}
					_, _ = fmt.Fprint(w, tt.publicResponse)
					return
				}
				wantPath := "/api/v2/instances/cluster/tenants/tenant/address"
				if tt.kind == "INTRANET" {
					wantPath = "/api/v2/instances/cluster/tenants/tenant/privatelink"
				}
				if req.Method != http.MethodGet || req.URL.Path != wantPath || req.URL.Query().Get("addressId") != "managed" {
					t.Errorf("unexpected query: %s %s", req.Method, req.URL)
				}
				w.Header().Set("Content-Type", "application/json")
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			cfg := obcloudsdk.NewConfiguration()
			cfg.Scheme, cfg.Host = "http", strings.TrimPrefix(server.URL, "http://")
			state := connectionModel()
			if tt.kind != "" {
				state.Type = types.StringValue(tt.kind)
			}
			err := refreshTenantConnection(context.Background(), &state, obcloudsdk.NewAPIClient(cfg))
			if errors.Is(err, errResourceNotFound) != tt.missing || (err != nil) != (tt.failure || tt.missing) {
				t.Fatalf("unexpected lookup result: %v", err)
			}
			if state.Id.ValueString() != "managed" {
				t.Fatalf("lookup changed managed identity: %q", state.Id.ValueString())
			}
			if err == nil && state.Status.ValueString() != "ONLINE" {
				t.Fatalf("unexpected status: %q", state.Status.ValueString())
			}
			if (publicQueries == 1) != (tt.publicResponse != "") {
				t.Fatalf("unexpected public address queries: %d", publicQueries)
			}
		})
	}
}

func TestConnectionCreateTimeoutRetainsIdentity(t *testing.T) {
	t.Setenv(EnvResourceCreateTimeoutSeconds, "1")
	created := 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch req.Method {
		case http.MethodPost:
			created++
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"managed"}}`)
		case http.MethodGet:
			if req.URL.Query().Get("addressId") != "managed" {
				t.Errorf("wrong address queried: %s", req.URL)
			}
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"CREATING"}]}`)
		default:
			t.Errorf("unexpected method %s", req.Method)
		}
	})
	plan := connectionModel()
	plan.Id = types.StringNull()
	state := connectionState(t, r, plan)
	response := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Config: tfsdk.Config{Raw: state.Raw, Schema: state.Schema}}, &response)
	if !response.Diagnostics.HasError() || created != 1 {
		t.Fatalf("expected one create and a timeout error, got %d and %v", created, response.Diagnostics)
	}
	var saved model.TenantConnection
	if diags := response.State.Get(context.Background(), &saved); diags.HasError() || saved.Id.ValueString() != "managed" {
		t.Fatalf("created identity was not retained: %q, %v", saved.Id.ValueString(), diags)
	}
}

func TestConnectionCreateWaitsForTenantAllocation(t *testing.T) {
	t.Setenv(EnvResourceCreateTimeoutSeconds, "3")
	var creates, addressQueries, tenantQueries atomic.Int32
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodPost:
			creates.Add(1)
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"managed"}}`)
		case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/address"):
			if req.URL.Query().Get("addressId") != "managed" {
				t.Errorf("wrong address queried: %s", req.URL)
			}
			addressQueries.Add(1)
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
		case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
			status := "ALLOCATING_INTERNET_ADDRESS"
			if tenantQueries.Add(1) >= 3 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
	})
	plan := connectionModel()
	plan.Id = types.StringNull()
	state := connectionState(t, r, plan)
	response := resource.CreateResponse{State: state}
	r.Create(context.Background(), resource.CreateRequest{Config: tfsdk.Config{Raw: state.Raw, Schema: state.Schema}}, &response)
	var saved model.TenantConnection
	diags := response.State.Get(context.Background(), &saved)
	if response.Diagnostics.HasError() || diags.HasError() || saved.Id.ValueString() != "managed" || creates.Load() != 1 || addressQueries.Load() != 1 || tenantQueries.Load() < 3 {
		t.Fatalf("creation did not wait for allocation: creates=%d addresses=%d tenants=%d id=%q diagnostics=%v", creates.Load(), addressQueries.Load(), tenantQueries.Load(), saved.Id.ValueString(), response.Diagnostics)
	}
}

func TestConnectionCreateAllocationFailureRetainsIdentity(t *testing.T) {
	for _, tt := range []struct {
		name, tenantResponse string
	}{
		{name: "timeout", tenantResponse: `{"success":true,"data":{"tenantId":"tenant","status":"ALLOCATING_INTERNET_ADDRESS"}}`},
		{name: "access denied", tenantResponse: `{"success":false,"errorCode":"AccessDenied"}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnvResourceCreateTimeoutSeconds, "1")
			var creates atomic.Int32
			r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case req.Method == http.MethodPost:
					creates.Add(1)
					_, _ = fmt.Fprint(w, `{"success":true,"data":{"addressId":"managed"}}`)
				case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/address"):
					_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
				case req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
					_, _ = fmt.Fprint(w, tt.tenantResponse)
				default:
					t.Errorf("unexpected request: %s %s", req.Method, req.URL)
				}
			})
			plan := connectionModel()
			plan.Id = types.StringNull()
			state := connectionState(t, r, plan)
			response := resource.CreateResponse{State: state}
			r.Create(context.Background(), resource.CreateRequest{Config: tfsdk.Config{Raw: state.Raw, Schema: state.Schema}}, &response)
			var saved model.TenantConnection
			diags := response.State.Get(context.Background(), &saved)
			if !response.Diagnostics.HasError() || diags.HasError() || saved.Id.ValueString() != "managed" || creates.Load() != 1 {
				t.Fatalf("allocation failure lost identity: creates=%d id=%q diagnostics=%v", creates.Load(), saved.Id.ValueString(), response.Diagnostics)
			}
		})
	}
}

func TestConnectionCreateRejectsUnsupportedOptionsBeforeAPICall(t *testing.T) {
	for _, tt := range []struct {
		name string
		set  func(*model.TenantConnection)
	}{
		{name: "enable_ssl", set: func(value *model.TenantConnection) { value.EnableSsl = types.BoolValue(false) }},
		{name: "enable_rpc", set: func(value *model.TenantConnection) { value.EnableRpc = types.BoolValue(true) }},
		{name: "private_link", set: func(value *model.TenantConnection) {
			value.PrivateLink = types.ObjectValueMust(value.PrivateLink.AttributeTypes(context.Background()), map[string]attr.Value{
				"user_account": types.StringValue("account"), "user_vpc": types.StringValue("vpc"), "user_vswitch": types.StringValue("vswitch"),
			})
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) { calls++ })
			value := connectionModel()
			value.Id = types.StringNull()
			tt.set(&value)
			configState := connectionState(t, r, value)
			config := tfsdk.Config{Raw: configState.Raw, Schema: configState.Schema}
			planned := tfsdk.Plan{Raw: configState.Raw, Schema: configState.Schema}
			prior := configState
			prior.RemoveResource(context.Background())
			var planResponse resource.ModifyPlanResponse
			r.ModifyPlan(context.Background(), resource.ModifyPlanRequest{Config: config, Plan: planned, State: prior}, &planResponse)
			if !planResponse.Diagnostics.HasError() {
				t.Fatalf("expected a plan error for %s", tt.name)
			}
			response := resource.CreateResponse{State: configState}
			r.Create(context.Background(), resource.CreateRequest{Config: config}, &response)
			if !response.Diagnostics.HasError() || calls != 0 {
				t.Fatalf("unsupported %s reached API: calls=%d diagnostics=%v", tt.name, calls, response.Diagnostics)
			}
		})
	}
}

func TestConnectionPlanRejectsUpdatesWithoutReplacement(t *testing.T) {
	for _, tt := range []struct {
		name string
		set  func(*model.TenantConnection)
	}{
		{name: "cluster_id", set: func(value *model.TenantConnection) { value.ClusterId = types.StringValue("other") }},
		{name: "tenant_id", set: func(value *model.TenantConnection) { value.TenantId = types.StringValue("other") }},
		{name: "type", set: func(value *model.TenantConnection) { value.Type = types.StringValue("INTRANET") }},
		{name: "enable_ssl", set: func(value *model.TenantConnection) { value.EnableSsl = types.BoolValue(true) }},
		{name: "enable_rpc", set: func(value *model.TenantConnection) { value.EnableRpc = types.BoolValue(true) }},
		{name: "private_link", set: func(value *model.TenantConnection) {
			value.PrivateLink = types.ObjectValueMust(value.PrivateLink.AttributeTypes(context.Background()), map[string]attr.Value{
				"user_account": types.StringValue("account"), "user_vpc": types.StringValue("vpc"), "user_vswitch": types.StringValue("vswitch"),
			})
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &tenantConnectionResource{}
			current := connectionModel()
			current.EnableSsl, current.EnableRpc = types.BoolValue(false), types.BoolValue(false)
			planned := current
			tt.set(&planned)
			state := connectionState(t, r, current)
			planState := connectionState(t, r, planned)
			var response resource.ModifyPlanResponse
			r.ModifyPlan(context.Background(), resource.ModifyPlanRequest{
				State: state, Plan: tfsdk.Plan{Raw: planState.Raw, Schema: planState.Schema},
			}, &response)
			if !response.Diagnostics.HasError() || len(response.RequiresReplace) != 0 || !strings.Contains(fmt.Sprint(response.Diagnostics), tt.name) {
				t.Fatalf("expected plan error for %s without replacement: %v", tt.name, response.Diagnostics)
			}
		})
	}
}

func TestConnectionPlanKeepsExistingOptions(t *testing.T) {
	r := &tenantConnectionResource{}
	current := connectionModel()
	current.EnableSsl, current.EnableRpc = types.BoolValue(true), types.BoolValue(false)
	state := connectionState(t, r, current)
	var response resource.ModifyPlanResponse
	r.ModifyPlan(context.Background(), resource.ModifyPlanRequest{
		State: state, Plan: tfsdk.Plan{Raw: state.Raw, Schema: state.Schema},
	}, &response)
	if response.Diagnostics.HasError() || len(response.RequiresReplace) != 0 {
		t.Fatalf("unchanged existing options rejected: %v", response.Diagnostics)
	}
}

func TestConnectionDeleteWaitsForExactAbsence(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "3")
	queries, deletes := 0, 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		if req.URL.Query().Get("addressId") != "managed" {
			t.Errorf("wrong address used: %s", req.URL)
		}
		switch req.Method {
		case http.MethodGet:
			queries++
			if queries < 3 {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
			} else {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
			}
		case http.MethodDelete:
			deletes++
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		default:
			t.Errorf("unexpected method %s", req.Method)
		}
	})
	state := connectionState(t, r, connectionModel())
	response := resource.DeleteResponse{State: state}
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &response)
	if response.Diagnostics.HasError() || deletes != 1 || queries < 3 || !response.State.Raw.IsNull() {
		t.Fatalf("deletion not confirmed: queries=%d deletes=%d state=%v diagnostics=%v", queries, deletes, response.State.Raw, response.Diagnostics)
	}
}

func TestConnectionReadAndDeleteDoNotTreatAccessFailureAsAbsence(t *testing.T) {
	queries := 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		if req.Method == http.MethodDelete {
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
			return
		}
		queries++
		if queries == 1 {
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed"}]}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":false,"errorCode":"AccessDenied"}`)
	})
	state := connectionState(t, r, connectionModel())
	deleteResponse := resource.DeleteResponse{State: state}
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &deleteResponse)
	if !deleteResponse.Diagnostics.HasError() || deleteResponse.State.Raw.IsNull() {
		t.Fatalf("delete lost state after lookup failure: %v", deleteResponse.Diagnostics)
	}
	readResponse := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &readResponse)
	if !readResponse.Diagnostics.HasError() || readResponse.State.Raw.IsNull() {
		t.Fatalf("read lost state after lookup failure: %v", readResponse.Diagnostics)
	}
}

func TestConnectionDeleteResumesPendingDeletion(t *testing.T) {
	queries := 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		if req.Method != http.MethodGet || req.URL.Query().Get("addressId") != "managed" {
			t.Errorf("unexpected request while resuming deletion: %s %s", req.Method, req.URL)
		}
		queries++
		if queries == 1 {
			_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"PENDING_DELETE"}]}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
	})
	state := connectionState(t, r, connectionModel())
	response := resource.DeleteResponse{State: state}
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &response)
	if response.Diagnostics.HasError() || queries != 2 || !response.State.Raw.IsNull() {
		t.Fatalf("pending deletion not reconciled: queries=%d diagnostics=%v", queries, response.Diagnostics)
	}
}

func TestConnectionReadRemovesOnlyConfirmedMissing(t *testing.T) {
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet && strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"ONLINE"}}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
	})
	state := connectionState(t, r, connectionModel())
	response := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
	if response.Diagnostics.HasError() || !response.State.Raw.IsNull() {
		t.Fatalf("confirmed missing connection remained in state: %v", response.Diagnostics)
	}
}

func TestConnectionReadOnlineDoesNotQueryTenant(t *testing.T) {
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodGet || !strings.HasSuffix(req.URL.Path, "/address") || req.URL.Query().Get("addressId") != "managed" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
	})
	state := connectionState(t, r, connectionModel())
	response := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
	if response.Diagnostics.HasError() || response.State.Raw.IsNull() {
		t.Fatalf("online read failed: %v", response.Diagnostics)
	}
}

func TestConnectionDeleteWaitsForTenantAfterNullAddress(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "3")
	addressQueries, tenantQueries, deletes := 0, 0, 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodDelete:
			deletes++
			_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
		case strings.HasSuffix(req.URL.Path, "/publicaddress"):
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"useSSL":false}}`)
		case strings.HasSuffix(req.URL.Path, "/address"):
			if req.URL.Query().Get("addressId") != "managed" {
				t.Errorf("wrong address used: %s", req.URL)
			}
			addressQueries++
			if addressQueries == 1 {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
			} else {
				_, _ = fmt.Fprint(w, `{"success":true,"data":[null]}`)
			}
		case strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
			tenantQueries++
			status := "PENDING_OFFLINE_INTERNET_ADDRESS"
			if tenantQueries == 1 || tenantQueries >= 4 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
	})
	state := connectionState(t, r, connectionModel())
	response := resource.DeleteResponse{State: state}
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &response)
	if response.Diagnostics.HasError() || !response.State.Raw.IsNull() || deletes != 1 || tenantQueries < 4 {
		t.Fatalf("release not completed: addresses=%d tenants=%d deletes=%d diagnostics=%v", addressQueries, tenantQueries, deletes, response.Diagnostics)
	}
}

func TestConnectionDeleteResumesTenantReleaseWithoutDeletingAgain(t *testing.T) {
	for _, initiallyMissing := range []bool{false, true} {
		t.Run(fmt.Sprintf("initiallyMissing=%t", initiallyMissing), func(t *testing.T) {
			t.Setenv(EnvResourceDeleteTimeoutSeconds, "3")
			addressQueries, tenantQueries, deletes := 0, 0, 0
			r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case req.Method == http.MethodDelete:
					deletes++
					_, _ = fmt.Fprint(w, `{"success":true,"data":true}`)
				case strings.HasSuffix(req.URL.Path, "/publicaddress"):
					_, _ = fmt.Fprint(w, `{"success":true,"data":{"useSSL":false}}`)
				case strings.HasSuffix(req.URL.Path, "/address"):
					addressQueries++
					if addressQueries == 1 && !initiallyMissing {
						_, _ = fmt.Fprint(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":"ONLINE"}]}`)
					} else {
						_, _ = fmt.Fprint(w, `{"success":true,"data":[null]}`)
					}
				case strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
					tenantQueries++
					status := "PENDING_OFFLINE_INTERNET_ADDRESS"
					if tenantQueries >= 3 {
						status = "ONLINE"
					}
					_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
				default:
					t.Errorf("unexpected request: %s %s", req.Method, req.URL)
				}
			})
			state := connectionState(t, r, connectionModel())
			response := resource.DeleteResponse{State: state}
			r.Delete(context.Background(), resource.DeleteRequest{State: state}, &response)
			if response.Diagnostics.HasError() || !response.State.Raw.IsNull() || deletes != 0 || tenantQueries < 3 {
				t.Fatalf("pending release not reconciled: addresses=%d tenants=%d deletes=%d diagnostics=%v", addressQueries, tenantQueries, deletes, response.Diagnostics)
			}
		})
	}
}

func TestConnectionReadWaitsForTenantRelease(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "3")
	tenantQueries, addressQueries := 0, 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method != http.MethodGet:
			t.Errorf("unexpected mutation during read: %s %s", req.Method, req.URL)
		case strings.HasSuffix(req.URL.Path, "/publicaddress"):
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"useSSL":false}}`)
		case strings.HasSuffix(req.URL.Path, "/address"):
			addressQueries++
			_, _ = fmt.Fprint(w, `{"success":true,"data":[null]}`)
		case strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
			tenantQueries++
			status := "PENDING_OFFLINE_INTERNET_ADDRESS"
			if tenantQueries >= 3 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
	})
	state := connectionState(t, r, connectionModel())
	response := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
	if response.Diagnostics.HasError() || !response.State.Raw.IsNull() || tenantQueries < 3 || addressQueries != 2 {
		t.Fatalf("read did not reconcile release: addresses=%d tenants=%d diagnostics=%v", addressQueries, tenantQueries, response.Diagnostics)
	}
}

func TestConnectionReadRechecksPendingAddressAfterRelease(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "3")
	addressQueries, tenantQueries := 0, 0
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasSuffix(req.URL.Path, "/address"):
			addressQueries++
			status := "PENDING_DELETE"
			if addressQueries > 1 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":[{"addressId":"managed","addressStatus":%q}]}`, status)
		case strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
			tenantQueries++
			status := "PENDING_OFFLINE_INTERNET_ADDRESS"
			if tenantQueries > 1 {
				status = "ONLINE"
			}
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant","status":%q}}`, status)
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
	})
	state := connectionState(t, r, connectionModel())
	response := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
	var saved model.TenantConnection
	diags := response.State.Get(context.Background(), &saved)
	if response.Diagnostics.HasError() || diags.HasError() || saved.Status.ValueString() != "ONLINE" || addressQueries != 2 || tenantQueries < 2 {
		t.Fatalf("read did not recheck address: addresses=%d tenants=%d state=%v diagnostics=%v", addressQueries, tenantQueries, saved.Status, response.Diagnostics)
	}
}

func TestConnectionReleaseStatusFailureRetainsState(t *testing.T) {
	for _, statusResponse := range []string{
		`{"success":false,"errorCode":"AccessDenied"}`,
		`{"success":true,"data":{"tenantId":"tenant"}}`,
		`{"success":true,"data":{"tenantId":"other","status":"ONLINE"}}`,
	} {
		t.Run(statusResponse, func(t *testing.T) {
			deletes := 0
			r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch {
				case req.Method == http.MethodDelete:
					deletes++
				case strings.HasSuffix(req.URL.Path, "/tenants/tenant"):
					_, _ = fmt.Fprint(w, statusResponse)
				default:
					_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
				}
			})
			state := connectionState(t, r, connectionModel())
			readResponse := resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, &readResponse)
			deleteResponse := resource.DeleteResponse{State: state}
			r.Delete(context.Background(), resource.DeleteRequest{State: state}, &deleteResponse)
			if !readResponse.Diagnostics.HasError() || readResponse.State.Raw.IsNull() || !deleteResponse.Diagnostics.HasError() || deleteResponse.State.Raw.IsNull() || deletes != 0 {
				t.Fatalf("status failure lost state: read=%v delete=%v deletes=%d", readResponse.Diagnostics, deleteResponse.Diagnostics, deletes)
			}
		})
	}
}

func TestConnectionReleaseTimeoutRetainsState(t *testing.T) {
	t.Setenv(EnvResourceDeleteTimeoutSeconds, "1")
	r := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodGet {
			t.Errorf("unexpected mutation: %s %s", req.Method, req.URL)
		}
		if strings.HasSuffix(req.URL.Path, "/tenants/tenant") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant","status":"PENDING_OFFLINE_INTERNET_ADDRESS"}}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":[]}`)
	})
	state := connectionState(t, r, connectionModel())
	readResponse := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &readResponse)
	deleteResponse := resource.DeleteResponse{State: state}
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &deleteResponse)
	if !readResponse.Diagnostics.HasError() || readResponse.State.Raw.IsNull() || !deleteResponse.Diagnostics.HasError() || deleteResponse.State.Raw.IsNull() {
		t.Fatalf("pending release lost state: read=%v delete=%v", readResponse.Diagnostics, deleteResponse.Diagnostics)
	}
}

func TestConnectionUpdateRejected(t *testing.T) {
	var response resource.UpdateResponse
	(&tenantConnectionResource{}).Update(context.Background(), resource.UpdateRequest{}, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("expected unsupported update to fail")
	}
}
