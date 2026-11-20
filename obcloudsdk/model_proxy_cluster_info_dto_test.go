package obcloudsdk

import (
	"encoding/json"
	"testing"
)

func TestConnectionProxyUnitNumDecimal(t *testing.T) {
	for _, tt := range []struct {
		jsonValue string
		want      float64
	}{
		{jsonValue: "1", want: 1},
		{jsonValue: "1.0", want: 1},
		{jsonValue: "1.25", want: 1.25},
	} {
		t.Run(tt.jsonValue, func(t *testing.T) {
			var result OBCloudResultListTenantConnectionDTO
			body := `{"success":true,"data":[{"addressId":"managed","proxyClusterInfo":{"unitNum":` + tt.jsonValue + `}}]}`
			if err := json.Unmarshal([]byte(body), &result); err != nil {
				t.Fatal(err)
			}
			info := result.GetData()[0].GetProxyClusterInfo()
			if info.GetUnitNum() != tt.want {
				t.Fatalf("got unitNum %v, want %v", info.GetUnitNum(), tt.want)
			}
		})
	}
}
