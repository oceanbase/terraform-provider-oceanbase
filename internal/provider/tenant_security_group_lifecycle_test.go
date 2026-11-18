package provider

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

func TestSecurityGroupRefresh(t *testing.T) {
	for _, tt := range []struct {
		name, response   string
		missing, failure bool
		want             string
	}{
		{name: "keep host prefix spelling", response: `{"success":true,"data":{"dataList":[{"securityIpGroupName":"example","securityIps":"192.0.2.1"}]}}`, want: "192.0.2.1/32"},
		{name: "detect drift", response: `{"success":true,"data":{"dataList":[{"securityIpGroupName":"example","securityIps":"192.0.2.3"}]}}`, want: "192.0.2.3"},
		{name: "deleted", response: `{"success":true,"data":{"total":0,"dataList":[]}}`, missing: true},
		{name: "another group", response: `{"success":true,"data":{"dataList":[{"securityIpGroupName":"other"}]}}`, missing: true},
		{name: "forbidden", response: `{"success":false,"errorCode":"PRI_ACCESS_DENIED_EXCEPTION"}`, failure: true},
		{name: "incomplete", response: `{"success":true,"data":{"total":2,"dataList":[]}}`, failure: true},
		{name: "no data", response: `{"success":true}`, failure: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/api/v2/instances/cluster/tenants/tenant/securityIpGroups" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.response))
			}))
			defer server.Close()
			cfg := obcloudsdk.NewConfiguration()
			cfg.Scheme, cfg.Host = "http", strings.TrimPrefix(server.URL, "http://")
			group := model.TenantSecurityGroup{
				ClusterId: types.StringValue("cluster"), TenantId: types.StringValue("tenant"),
				Name: types.StringValue("example"), Cidrs: []types.String{types.StringValue("192.0.2.1/32")},
			}
			err := refreshTenantSecurityGroup(context.Background(), &group, obcloudsdk.NewAPIClient(cfg))
			if errors.Is(err, errResourceNotFound) != tt.missing || (err != nil) != (tt.failure || tt.missing) {
				t.Fatalf("unexpected lookup result: %v", err)
			}
			if tt.want != "" && (len(group.Cidrs) != 1 || group.Cidrs[0].ValueString() != tt.want) {
				t.Errorf("unexpected CIDRs: %v", group.Cidrs)
			}
		})
	}
}

func TestSecurityGroupRequestHostPrefix(t *testing.T) {
	group := model.TenantSecurityGroup{
		Name:  types.StringValue("example"),
		Cidrs: []types.String{types.StringValue("192.0.2.1/32"), types.StringValue("198.51.100.0/24")},
	}
	body, err := securityGroupRequest(&group)
	if err != nil {
		t.Fatal(err)
	}
	var ips []string
	if err := json.Unmarshal([]byte(body.GetSecurityIps()), &ips); err != nil ||
		len(ips) != 2 || ips[0] != "192.0.2.1" || ips[1] != "198.51.100.0/24" {
		t.Fatalf("unexpected API securityIps: %q", body.GetSecurityIps())
	}
}
