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
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
		status             int
		missing, failure   bool
		wantList           bool
	}{
		{name: "missing confirmed", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":true,"data":{"dataList":[],"total":0,"hasMore":false}}`, missing: true, wantList: true},
		{name: "detail missing but id exists", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":true,"data":{"dataList":[{"tenantId":"tenant-id"}],"total":1,"hasMore":false}}`, failure: true, wantList: true},
		{name: "list denied", detail: `{"success":false,"errorCode":"Tenant.NotExists"}`, list: `{"success":false,"errorCode":"AccessDenied"}`, failure: true, wantList: true},
		{name: "detail denied", detail: `{"success":false,"errorCode":"AccessDenied"}`, failure: true},
		{name: "detail resource state error", detail: `{"success":false,"errorCode":"IllegalOperation.Resource"}`, failure: true},
		{name: "detail http denied", status: http.StatusForbidden, detail: `forbidden`, failure: true},
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
					if r.URL.Query().Has("tenantName") {
						t.Error("tenant list must not be filtered by name")
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
		})
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
				_, _ = fmt.Fprintf(w, `{"success":true,"data":{"tenantId":"tenant-id","status":"ONLINE","primaryZone":%q}}`, tt.actualZone)
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
