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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

func tenantTestClient(t *testing.T, handler http.HandlerFunc) *obcloudsdk.APIClient {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	cfg := obcloudsdk.NewConfiguration()
	cfg.Scheme, cfg.Host = "http", strings.TrimPrefix(server.URL, "http://")
	return obcloudsdk.NewAPIClient(cfg)
}

func testTenant() model.Tenant {
	return model.Tenant{TenantSpec: model.TenantSpec{ClusterId: types.StringValue("cluster"), EnablePublicAddress: types.BoolNull()}, TenantInfo: model.TenantInfo{Id: types.StringValue("tenant-id")}}
}

func TestTenantPublicAddressReadAndCreation(t *testing.T) {
	for _, tt := range []struct {
		name, publicResponse string
		initial, expected    types.Bool
		want                 types.Bool
		wait, wantError      bool
		publicCalls          int
	}{
		{name: "data source has no address despite capability", initial: types.BoolNull(), publicResponse: `{"success":true,"data":{}}`, want: types.BoolValue(false), publicCalls: 1},
		{name: "data source has address", initial: types.BoolNull(), publicResponse: `{"success":true,"data":{"addressId":"public-id"}}`, want: types.BoolValue(true), publicCalls: 1},
		{name: "existing creation value does not own independent address", initial: types.BoolValue(false), want: types.BoolValue(false)},
		{name: "create false verifies empty address", initial: types.BoolValue(false), expected: types.BoolValue(false), publicResponse: `{"success":true,"data":{}}`, want: types.BoolValue(false), publicCalls: 1, wait: true},
		{name: "create true verifies existing address", initial: types.BoolValue(true), expected: types.BoolValue(true), publicResponse: `{"success":true,"data":{"addressId":"public-id"}}`, want: types.BoolValue(true), publicCalls: 1, wait: true},
		{name: "mismatched address is retained and errors", initial: types.BoolValue(false), expected: types.BoolValue(false), publicResponse: `{"success":true,"data":{"addressId":"public-id"}}`, want: types.BoolValue(true), publicCalls: 1, wait: true, wantError: true},
		{name: "address lookup error leaves value unconfirmed", initial: types.BoolValue(false), expected: types.BoolValue(false), publicResponse: `{"success":false,"errorCode":"AccessDenied"}`, want: types.BoolNull(), publicCalls: 1, wait: true, wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var calls int
			client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if strings.HasSuffix(r.URL.Path, "/publicaddress") {
					calls++
					_, _ = fmt.Fprint(w, tt.publicResponse)
					return
				}
				_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant-id","status":"ONLINE","enableProxyPublicLink":true}}`)
			})
			tenant := testTenant()
			tenant.EnablePublicAddress = tt.initial
			var err error
			if tt.wait {
				err = waitForTenantCreation(context.Background(), &tenant, client, tt.expected, 25*time.Millisecond)
			} else {
				err = readTenant(context.Background(), &tenant, client)
			}
			if (err != nil) != tt.wantError {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tenant.EnablePublicAddress.Equal(tt.want) || tenant.Id.ValueString() != "tenant-id" {
				t.Fatalf("unexpected state: id=%s, public=%v", tenant.Id.ValueString(), tenant.EnablePublicAddress)
			}
			if calls != tt.publicCalls {
				t.Fatalf("public address calls = %d, want %d", calls, tt.publicCalls)
			}
		})
	}
}

func TestTenantReadConfirmsMissingWithoutMaskingErrors(t *testing.T) {
	for _, tt := range []struct {
		name, detail, list string
		wantErrorText      string
		status, listStatus int
		missing, failure   bool
		wantList           bool
	}{
		{name: "missing confirmed", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":true,"data":{"dataList":[],"total":0,"hasMore":false}}`, missing: true, wantList: true},
		{name: "detail missing but id exists", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id"}],"total":1,"hasMore":false}}`, failure: true, wantList: true},
		{name: "list denied", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":false,"errorCode":"AccessDenied"}`, failure: true, wantList: true},
		{name: "detail denied", detail: `{"success":false,"errorCode":"AccessDenied"}`, failure: true},
		{name: "resource state error and absent tenant", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"dataList":[],"total":0,"hasMore":false}}`, missing: true, wantList: true},
		{name: "resource state error and present tenant", detail: `{"success":false,"errorCode":"IllegalOperation.Resource","errorMessage":"not online"}`, list: `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id"}],"total":1,"hasMore":false}}`, failure: true, wantList: true, wantErrorText: "failed to describe tenant: <nil>, code: IllegalOperation.Resource, message: not online"},
		{name: "resource state error and list denied", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":false,"errorCode":"AccessDenied"}`, failure: true, wantList: true},
		{name: "resource state error and list http denied", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, listStatus: http.StatusForbidden, list: `{"success":false,"errorCode":"AccessDenied"}`, failure: true, wantList: true},
		{name: "resource state error and missing data list", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"total":0,"hasMore":false}}`, failure: true, wantList: true},
		{name: "resource state error and missing pagination", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"dataList":[]}}`, failure: true, wantList: true},
		{name: "resource state error and incomplete pagination", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"dataList":[],"total":1,"hasMore":false}}`, failure: true, wantList: true},
		{name: "resource state error and inconsistent total", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"dataList":[{"tenantId":"other"}],"total":0,"hasMore":false}}`, failure: true, wantList: true},
		{name: "resource state error and unfinished pagination", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{"success":true,"data":{"dataList":[],"hasMore":true}}`, failure: true, wantList: true},
		{name: "resource state error and malformed list", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, list: `{`, failure: true, wantList: true},
		{name: "detail http denied", status: http.StatusForbidden, detail: `forbidden`, failure: true},
		{name: "detail http denied with resource code", status: http.StatusForbidden, detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, failure: true},
		{name: "detail malformed", detail: `{`, failure: true},
		{name: "detail has no id", detail: `{"success":true,"data":{}}`, failure: true},
		{name: "detail has different id", detail: `{"success":true,"data":{"tenantId":"other"}}`, failure: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var listCalls int
			client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if strings.HasSuffix(r.URL.Path, "/tenants") {
					listCalls++
					if r.URL.Path != "/api/v2/instances/cluster/tenants" {
						t.Errorf("unexpected tenant list scope: %s", r.URL.Path)
					}
					if r.URL.Query().Has("tenantName") {
						t.Error("tenant list must not be filtered by name")
					}
					if tt.listStatus != 0 {
						w.WriteHeader(tt.listStatus)
					}
					_, _ = fmt.Fprint(w, tt.list)
					return
				}
				if r.URL.Path != "/api/v2/instances/cluster/tenants/tenant-id" {
					t.Errorf("unexpected tenant detail scope: %s", r.URL.Path)
				}
				if tt.status != 0 {
					w.WriteHeader(tt.status)
				}
				_, _ = fmt.Fprint(w, tt.detail)
			})
			tenant := testTenant()
			err := readTenant(context.Background(), &tenant, client)
			if errors.Is(err, errResourceNotFound) != tt.missing || (err != nil) != (tt.failure || tt.missing) {
				t.Fatalf("unexpected lookup result: %v", err)
			}
			if (listCalls > 0) != tt.wantList {
				t.Fatalf("unexpected list calls: %d", listCalls)
			}
			if tt.wantErrorText != "" && err.Error() != tt.wantErrorText {
				t.Fatalf("original detail error was not preserved: %v", err)
			}
		})
	}
}

func TestTenantReadNetworkErrorDoesNotQueryList(t *testing.T) {
	var listCalls int
	cfg := obcloudsdk.NewConfiguration()
	cfg.Scheme, cfg.Host = "http", "example.test"
	cfg.HTTPClient = &http.Client{Transport: connectionRoundTripper(func(req *http.Request) (*http.Response, error) {
		if strings.HasSuffix(req.URL.Path, "/tenants") {
			listCalls++
		}
		return nil, errors.New("network unavailable")
	})}
	tenant := testTenant()
	err := readTenant(context.Background(), &tenant, obcloudsdk.NewAPIClient(cfg))
	if err == nil || errors.Is(err, errResourceNotFound) || listCalls != 0 {
		t.Fatalf("network failure was not preserved: list calls=%d err=%v", listCalls, err)
	}
}

func TestTenantReadResourceStateErrorChecksEveryPage(t *testing.T) {
	for _, present := range []bool{false, true} {
		t.Run(fmt.Sprintf("present=%t", present), func(t *testing.T) {
			var pages int
			client := tenantTestClient(t, func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if !strings.HasSuffix(req.URL.Path, "/tenants") {
					_, _ = fmt.Fprint(w, `{"success":false,"errorCode":"IllegalOperation.Resource"}`)
					return
				}
				pages++
				if req.URL.Query().Get("pageNumber") != fmt.Sprint(pages) {
					t.Errorf("wrong page: %s", req.URL.RawQuery)
				}
				if pages == 1 {
					_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[{"tenantId":"other"}],"total":2,"hasMore":true}}`)
					return
				}
				id := "another"
				if present {
					id = "tenant-id"
				}
				_, _ = fmt.Fprintf(w, `{"success":true,"data":{"dataList":[{"tenantId":%q}],"total":2,"hasMore":false}}`, id)
			})
			tenant := testTenant()
			err := readTenant(context.Background(), &tenant, client)
			if pages != 2 || errors.Is(err, errResourceNotFound) == present || err == nil {
				t.Fatalf("pages=%d present=%t err=%v", pages, present, err)
			}
			if present && !strings.Contains(err.Error(), "IllegalOperation.Resource") {
				t.Fatalf("original detail error lost: %v", err)
			}
		})
	}
}

func TestTenantReadResourceAndDataSourceHandleConfirmedDeletion(t *testing.T) {
	provider := connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(req.URL.Path, "/tenants") {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[],"total":0,"hasMore":false}}`)
		} else {
			_, _ = fmt.Fprint(w, `{"success":false,"errorCode":"IllegalOperation.Resource"}`)
		}
	}).provider
	ctx := context.Background()
	state := tenantUpdateState(t)
	resourceResponse := resource.ReadResponse{State: state}
	(&tenantResource{provider: provider}).Read(ctx, resource.ReadRequest{State: state}, &resourceResponse)
	if resourceResponse.Diagnostics.HasError() || !resourceResponse.State.Raw.IsNull() {
		t.Fatalf("resource did not clear confirmed deleted state: %v", resourceResponse.Diagnostics)
	}

	d := &tenantDataSource{provider: provider}
	var schemaResponse datasource.SchemaResponse
	d.Schema(ctx, datasource.SchemaRequest{}, &schemaResponse)
	raw, err := (&tfprotov6.RawState{JSON: []byte(`{"cluster_id":"cluster","id":"tenant-id"}`)}).Unmarshal(schemaResponse.Schema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatal(err)
	}
	dataResponse := datasource.ReadResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	d.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: schemaResponse.Schema, Raw: raw}}, &dataResponse)
	if !dataResponse.Diagnostics.HasError() {
		t.Fatal("data source succeeded for a confirmed deleted tenant")
	}
}

func TestTenantListChecksAllPagesAndExactID(t *testing.T) {
	var calls int
	client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Has("tenantName") {
			t.Error("tenant list must not be filtered by name")
		}
		if r.URL.Query().Get("pageNumber") == "1" {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id-other"}],"total":2,"hasMore":true}}`)
		} else {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id"}],"total":2,"hasMore":false}}`)
		}
	})
	tenant := testTenant()
	if err := findTenant(context.Background(), &tenant, client); err != nil || calls != 2 {
		t.Fatalf("expected exact id on second page, calls=%d err=%v", calls, err)
	}
}

func TestTenantCreationOnlyChanges(t *testing.T) {
	state := testTenant()
	state.EnablePublicAddress = types.BoolValue(false)
	state.Whitelist = types.StringValue("192.0.2.1")
	plan := state
	if err := validateTenantCreationInputs(&state, &plan); err != nil {
		t.Fatal(err)
	}
	plan.EnablePublicAddress = types.BoolValue(true)
	if err := validateTenantCreationInputs(&state, &plan); err == nil || !strings.Contains(err.Error(), "oceanbase_tenant_connection") {
		t.Fatalf("expected public address update diagnostic, got %v", err)
	}
	plan = state
	plan.Whitelist = types.StringValue("192.0.2.2")
	if err := validateTenantCreationInputs(&state, &plan); err == nil || !strings.Contains(err.Error(), "oceanbase_tenant_security_group") {
		t.Fatalf("expected whitelist update diagnostic, got %v", err)
	}
}

func TestTenantUpgradeLegacyPublicAddressState(t *testing.T) {
	ctx := context.Background()
	server, err := testAccProtoV6ProviderFactories["oceanbase"]()
	if err != nil {
		t.Fatal(err)
	}
	legacy := &tfprotov6.RawState{JSON: []byte(`{"id":"tenant-id","cluster_id":"cluster","name":"tenant","enable_public_address":true,"whitelist":"192.0.2.1","unit":{"cpu":2,"memory":8},"tags":[{"key":"team","value":"db"}]}`)}
	upgraded, err := server.UpgradeResourceState(ctx, &tfprotov6.UpgradeResourceStateRequest{
		TypeName: "oceanbase_tenant", Version: 0, RawState: legacy,
	})
	if err != nil || len(upgraded.Diagnostics) != 0 {
		t.Fatalf("state upgrade failed: %v, diagnostics: %v", err, upgraded.Diagnostics)
	}
	var schemaResponse resource.SchemaResponse
	(&tenantResource{}).Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	if schemaResponse.Schema.Version != 1 {
		t.Fatalf("unexpected schema version: %d", schemaResponse.Schema.Version)
	}
	schemaType := schemaResponse.Schema.Type().TerraformType(ctx)
	previousValue, err := legacy.Unmarshal(schemaType)
	if err != nil {
		t.Fatal(err)
	}
	upgradedValue, err := upgraded.UpgradedState.Unmarshal(schemaType)
	if err != nil {
		t.Fatal(err)
	}
	var previousFields, upgradedFields map[string]tftypes.Value
	if err := previousValue.As(&previousFields); err != nil {
		t.Fatal(err)
	}
	if err := upgradedValue.As(&upgradedFields); err != nil {
		t.Fatal(err)
	}
	for name, before := range previousFields {
		if name != "enable_public_address" && !before.Equal(upgradedFields[name]) {
			t.Fatalf("upgrade changed %s: %v to %v", name, before, upgradedFields[name])
		}
	}
	if !upgradedFields["enable_public_address"].IsNull() {
		t.Fatal("legacy capability flag was not cleared")
	}

	state := tfsdk.State{Schema: schemaResponse.Schema, Raw: upgradedValue}
	var tenant model.Tenant
	if diags := state.Get(ctx, &tenant); diags.HasError() {
		t.Fatalf("could not decode upgraded state: %v", diags)
	}
	var publicCalls int
	client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/publicaddress") {
			publicCalls++
			_, _ = fmt.Fprint(w, `{"success":true,"data":{}}`)
			return
		}
		_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant-id","status":"ONLINE","enableProxyPublicLink":true}}`)
	})
	for i := 0; i < 2; i++ {
		if err := readTenant(ctx, &tenant, client); err != nil {
			t.Fatalf("refresh after upgrade failed: %v", err)
		}
		plan := tenant
		plan.EnablePublicAddress = types.BoolValue(false)
		if err := validateTenantCreationInputs(&tenant, &plan); err != nil {
			t.Fatalf("unchanged false configuration was rejected after upgrade: %v", err)
		}
	}
	if tenant.Id.ValueString() != "tenant-id" || tenant.EnablePublicAddress.ValueBool() || publicCalls != 1 {
		t.Fatalf("unexpected refreshed state: id=%s public=%v public calls=%d", tenant.Id.ValueString(), tenant.EnablePublicAddress, publicCalls)
	}
}

func TestTenantUpdateWaitsForRequestedState(t *testing.T) {
	for _, tt := range []struct {
		name, actualZone string
		wantError        bool
	}{
		{name: "online with old value times out", actualZone: "zone-a", wantError: true},
		{name: "online with requested value completes", actualZone: "zone-b"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant-id","tenantName":"tenant","status":"ONLINE","primaryZone":%q,"unitNum":1,"unitCpu":2,"unitMem":8}}`, tt.actualZone)
			})
			tenant := testTenant()
			tenant.EnablePublicAddress = types.BoolValue(false)
			plan := tenant
			plan.PrimaryZone = types.StringValue("zone-b")
			err := waitForTenantUpdate(context.Background(), &tenant, &plan, OperationModifyTenantPrimaryZone, client, 25*time.Millisecond)
			if (err != nil) != tt.wantError || tenant.Id.ValueString() != "tenant-id" || tenant.PrimaryZone.ValueString() != tt.actualZone {
				t.Fatalf("unexpected update result: id=%s zone=%s err=%v", tenant.Id.ValueString(), tenant.PrimaryZone.ValueString(), err)
			}
		})
	}
}

func tenantUpdateState(t *testing.T) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	var schemaResponse resource.SchemaResponse
	(&tenantResource{}).Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	raw, err := (&tfprotov6.RawState{JSON: []byte(`{"id":"tenant-id","cluster_id":"cluster","name":"old-name","mode":"NORMAL","primary_zone":"zone-a","charset":"utf8mb4","timezone":"UTC","unit_num":1,"unit":{"cpu":2,"memory":8},"enable_public_address":false,"whitelist":"192.0.2.1","compatible_mode":"MYSQL","description":"description","lower_case_table_names":1}`)}).Unmarshal(schemaResponse.Schema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatal(err)
	}
	return tfsdk.State{Schema: schemaResponse.Schema, Raw: raw}
}

func tenantUpdatePlan(t *testing.T, state tfsdk.State) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}
	for name, value := range map[string]attr.Value{"name": types.StringValue("new-name"), "primary_zone": types.StringValue("zone-b")} {
		if diags := plan.SetAttribute(context.Background(), path.Root(name), value); diags.HasError() {
			t.Fatalf("set %s: %v", name, diags)
		}
	}
	return plan
}

func TestTenantUpdateRejectsEveryUnsupportedFieldBeforeAPICalls(t *testing.T) {
	var calls int
	r := &tenantResource{provider: connectionResourceWithHandler(t, func(http.ResponseWriter, *http.Request) { calls++ }).provider}
	for _, tt := range []struct {
		name  string
		value attr.Value
	}{
		{"cluster_id", types.StringValue("other")},
		{"mode", types.StringValue("READONLY")},
		{"charset", types.StringValue("other")},
		{"timezone", types.StringValue("other")},
		{"compatible_mode", types.StringValue("other")},
		{"description", types.StringValue("other")},
		{"description", types.StringNull()},
		{"lower_case_table_names", types.Int32Value(0)},
		{"enable_public_address", types.BoolValue(true)},
		{"whitelist", types.StringValue("other")},
		{"whitelist", types.StringNull()},
	} {
		t.Run(fmt.Sprintf("%s=%v", tt.name, tt.value), func(t *testing.T) {
			state := tenantUpdateState(t)
			plan := tenantUpdatePlan(t, state)
			if diags := plan.SetAttribute(context.Background(), path.Root(tt.name), tt.value); diags.HasError() {
				t.Fatal(diags)
			}
			var planResponse resource.ModifyPlanResponse
			r.ModifyPlan(context.Background(), resource.ModifyPlanRequest{State: state, Plan: plan}, &planResponse)
			if !planResponse.Diagnostics.HasError() || !strings.Contains(fmt.Sprint(planResponse.Diagnostics), tt.name) {
				t.Fatalf("plan accepted %s: %v", tt.name, planResponse.Diagnostics)
			}
			response := resource.UpdateResponse{State: state}
			r.Update(context.Background(), resource.UpdateRequest{State: state, Plan: plan}, &response)
			if !response.Diagnostics.HasError() || calls != 0 {
				t.Fatalf("update accepted %s, calls=%d diagnostics=%v", tt.name, calls, response.Diagnostics)
			}
		})
	}
}

func TestTenantUpdateRetriesOnlyUnfinishedOperation(t *testing.T) {
	name, zone := "old-name", "zone-a"
	var events []string
	zoneCalls := 0
	r := &tenantResource{provider: connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodGet:
			events = append(events, "get:"+name+":"+zone)
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant-id","tenantName":%q,"tenantMode":"NORMAL","status":"ONLINE","primaryZone":%q,"unitNum":1,"unitCpu":2,"unitMem":8,"charset":"utf8mb4","timeZone":"UTC","tenantCompatibilityMode":"MYSQL","description":"description","lowerCaseTableNamesParam":1}}`, name, zone)
		case strings.HasSuffix(req.URL.Path, "/tenantName"):
			events = append(events, "name")
			name = "new-name"
			_, _ = fmt.Fprint(w, `{"success":true}`)
		case strings.HasSuffix(req.URL.Path, "/primaryZone"):
			events = append(events, "zone")
			zoneCalls++
			if zoneCalls == 1 {
				_, _ = fmt.Fprint(w, `{"success":false}`)
			} else {
				zone = "zone-b"
				_, _ = fmt.Fprint(w, `{"success":true}`)
			}
		default:
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
	}).provider}
	state := tenantUpdateState(t)
	plan := tenantUpdatePlan(t, state)
	first := resource.UpdateResponse{State: state}
	r.Update(context.Background(), resource.UpdateRequest{State: state, Plan: plan}, &first)
	var observed model.Tenant
	if diags := first.State.Get(context.Background(), &observed); diags.HasError() {
		t.Fatal(diags)
	}
	if !first.Diagnostics.HasError() || observed.Name.ValueString() != "new-name" || observed.PrimaryZone.ValueString() != "zone-a" {
		t.Fatalf("partial state was not preserved: name=%s zone=%s diagnostics=%v", observed.Name, observed.PrimaryZone, first.Diagnostics)
	}
	second := resource.UpdateResponse{State: first.State}
	r.Update(context.Background(), resource.UpdateRequest{State: first.State, Plan: plan}, &second)
	if diags := second.State.Get(context.Background(), &observed); diags.HasError() {
		t.Fatal(diags)
	}
	if second.Diagnostics.HasError() || observed.Name.ValueString() != "new-name" || observed.PrimaryZone.ValueString() != "zone-b" {
		t.Fatalf("retry did not complete: name=%s zone=%s diagnostics=%v", observed.Name, observed.PrimaryZone, second.Diagnostics)
	}
	want := "get:old-name:zone-a,name,get:new-name:zone-a,zone,get:new-name:zone-a,get:new-name:zone-a,zone,get:new-name:zone-b"
	if got := strings.Join(events, ","); got != want {
		t.Fatalf("unexpected operation order: %s", got)
	}
}

func TestTenantUpdateAppliesSupportedChangesInOrder(t *testing.T) {
	name, zone := "old-name", "zone-a"
	cpu, memory, unitNum := int32(2), int32(8), int32(1)
	var events []string
	r := &tenantResource{provider: connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if req.Method == http.MethodGet {
			events = append(events, "get")
			_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant-id","tenantName":%q,"tenantMode":"NORMAL","status":"ONLINE","primaryZone":%q,"unitNum":%d,"unitCpu":%d,"unitMem":%d,"charset":"utf8mb4","timeZone":"UTC","tenantCompatibilityMode":"MYSQL","description":"description","lowerCaseTableNamesParam":1}}`, name, zone, unitNum, cpu, memory)
			return
		}
		var body map[string]interface{}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Errorf("invalid request body: %v", err)
		}
		switch {
		case strings.HasSuffix(req.URL.Path, "/tenantName"):
			events = append(events, "name")
			if body["tenantName"] != "new-name" {
				t.Errorf("wrong name request: %v", body)
			}
			name = "new-name"
		case strings.HasSuffix(req.URL.Path, "/primaryZone"):
			events = append(events, "zone")
			if body["primaryZone"] != "zone-b" {
				t.Errorf("wrong zone request: %v", body)
			}
			zone = "zone-b"
		default:
			events = append(events, "unit")
			if body["cpu"] != float64(4) || body["memory"] != float64(16) || body["unitNum"] != float64(2) {
				t.Errorf("wrong unit request: %v", body)
			}
			cpu, memory, unitNum = 4, 16, 2
		}
		_, _ = fmt.Fprint(w, `{"success":true}`)
	}).provider}
	state := tenantUpdateState(t)
	plan := tenantUpdatePlan(t, state)
	unit := types.ObjectValueMust(map[string]attr.Type{"cpu": types.Int32Type, "memory": types.Int32Type}, map[string]attr.Value{"cpu": types.Int32Value(4), "memory": types.Int32Value(16)})
	for name, value := range map[string]attr.Value{"unit": unit, "unit_num": types.Int32Value(2)} {
		if diags := plan.SetAttribute(context.Background(), path.Root(name), value); diags.HasError() {
			t.Fatal(diags)
		}
	}
	response := resource.UpdateResponse{State: state}
	r.Update(context.Background(), resource.UpdateRequest{State: state, Plan: plan}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	var observed model.Tenant
	if diags := response.State.Get(context.Background(), &observed); diags.HasError() {
		t.Fatal(diags)
	}
	if !observed.Unit.Equal(unit) || observed.UnitNum.ValueInt32() != 2 || observed.Name.ValueString() != "new-name" || observed.PrimaryZone.ValueString() != "zone-b" {
		t.Fatalf("unexpected final state: %+v", observed)
	}
	if got, want := strings.Join(events, ","), "get,unit,get,name,get,zone,get"; got != want {
		t.Fatalf("unexpected operation order: %s", got)
	}
}

func TestTenantUpdateDoesNotContinueBeforeNameIsObserved(t *testing.T) {
	var nameCalls, zoneCalls int
	r := &tenantResource{provider: connectionResourceWithHandler(t, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case req.Method == http.MethodGet:
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant-id","tenantName":"old-name","tenantMode":"NORMAL","status":"ONLINE","primaryZone":"zone-a","unitNum":1,"unitCpu":2,"unitMem":8,"charset":"utf8mb4","timeZone":"UTC","tenantCompatibilityMode":"MYSQL","description":"description","lowerCaseTableNamesParam":1}}`)
		case strings.HasSuffix(req.URL.Path, "/tenantName"):
			nameCalls++
			_, _ = fmt.Fprint(w, `{"success":true}`)
		case strings.HasSuffix(req.URL.Path, "/primaryZone"):
			zoneCalls++
		}
	}).provider}
	state := tenantUpdateState(t)
	plan := tenantUpdatePlan(t, state)
	response := resource.UpdateResponse{State: state}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	r.Update(ctx, resource.UpdateRequest{State: state, Plan: plan}, &response)
	var observed model.Tenant
	if diags := response.State.Get(context.Background(), &observed); diags.HasError() {
		t.Fatal(diags)
	}
	if !response.Diagnostics.HasError() || nameCalls != 1 || zoneCalls != 0 || observed.Name.ValueString() != "old-name" || observed.PrimaryZone.ValueString() != "zone-a" {
		t.Fatalf("unobserved name continued: name calls=%d zone calls=%d state=%s/%s diagnostics=%v", nameCalls, zoneCalls, observed.Name, observed.PrimaryZone, response.Diagnostics)
	}
}

func TestTenantUpdateChecksNameAndIncompleteDetails(t *testing.T) {
	for _, tt := range []struct {
		name, response string
		wantError      bool
	}{
		{"old name", `"tenantName":"old-name"`, true},
		{"requested name", `"tenantName":"new-name"`, false},
		{"missing name", ``, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if tt.response != "" {
					_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant-id","status":"ONLINE","primaryZone":"zone-a","unitNum":1,"unitCpu":2,"unitMem":8,%s}}`, tt.response)
				} else {
					_, _ = fmt.Fprint(w, `{"success":true,"data":{"tenantId":"tenant-id","status":"ONLINE","primaryZone":"zone-a","unitNum":1,"unitCpu":2,"unitMem":8}}`)
				}
			})
			tenant := testTenant()
			tenant.Name = types.StringValue("old-name")
			tenant.EnablePublicAddress = types.BoolValue(false)
			plan := tenant
			plan.Name = types.StringValue("new-name")
			err := waitForTenantUpdate(context.Background(), &tenant, &plan, OperationModifyTenantName, client, 25*time.Millisecond)
			if (err != nil) != tt.wantError || (tt.wantError && tenant.Name.ValueString() != "old-name") {
				t.Fatalf("unexpected name=%s err=%v", tenant.Name, err)
			}
		})
	}
}

func TestTenantDeletionWaitsForConfirmedDisappearance(t *testing.T) {
	var calls int
	client := tenantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		if calls == 1 {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id","status":"PENDING_OFFLINE"}],"total":1,"hasMore":false}}`)
		} else {
			_, _ = fmt.Fprint(w, `{"success":true,"data":{"dataList":[],"total":0,"hasMore":false}}`)
		}
	})
	tenant := testTenant()
	err := waitForRemoval(context.Background(), time.Second, func(ctx context.Context) error {
		return findTenant(ctx, &tenant, client)
	})
	if err != nil || calls != 2 {
		t.Fatalf("expected pending tenant to disappear, calls=%d err=%v", calls, err)
	}
}
