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
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/api"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/internal/model"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

// DescribeTenantPublicAddress returns an empty object when no address exists.
func tenantHasPublicAddress(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) (bool, error) {
	result, response, err := client.MultiCloudOpenAPI.DescribeTenantPublicAddress(ctx, tenant.ClusterId.ValueString(), tenant.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		return false, fmt.Errorf("failed to describe tenant public address: %v", err)
	}
	return result.Data.GetAddressId() != "", nil
}

// Compare IDs across every page; only a complete successful list is sufficient
// evidence that a tenant is gone.
func findTenant(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) error {
	_, err := tenantListStatus(ctx, tenant, client)
	return err
}

// This detail error also occurs for a deleted tenant; a list must distinguish it.
type tenantDetailResourceStateError struct{ error }

func tenantListStatus(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) (string, error) {
	if tenant.ClusterId.IsNull() || tenant.ClusterId.IsUnknown() || tenant.ClusterId.ValueString() == "" ||
		tenant.Id.IsNull() || tenant.Id.IsUnknown() || tenant.Id.ValueString() == "" {
		return "", errors.New("tenant cluster id and id are required")
	}
	const pageSize int32 = 100
	var seen int64
	for page := int32(1); ; page++ {
		result, response, err := client.MultiCloudOpenAPI.QueryTenants(ctx, tenant.ClusterId.ValueString()).RequestId("").
			PageNumber(page).PageSize(pageSize).Execute()
		api.LogAPIResult(ctx, result, response, err)
		if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
			return "", fmt.Errorf("failed to list tenants: %v", err)
		}
		data := result.GetData()
		if !data.HasDataList() {
			return "", errors.New("tenant list omitted its data list")
		}
		for _, item := range data.GetDataList() {
			if item.GetTenantId() == "" {
				return "", errors.New("tenant list returned an item without an id")
			}
			if item.GetTenantId() == tenant.Id.ValueString() {
				return item.GetStatus(), nil
			}
		}
		seen += int64(len(data.GetDataList()))
		if !data.HasHasMore() && !data.HasTotal() {
			return "", errors.New("tenant list omitted pagination information")
		}
		more := data.GetHasMore()
		if !data.HasHasMore() {
			more = seen < data.GetTotal()
		}
		if !more {
			if data.HasTotal() && seen != data.GetTotal() {
				return "", errors.New("tenant list pagination is incomplete")
			}
			return "", fmt.Errorf("tenant %s: %w", tenant.Id.ValueString(), errResourceNotFound)
		}
		if len(data.GetDataList()) == 0 {
			return "", errors.New("tenant list has an empty page before the end")
		}
	}
}

func readTenant(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) error {
	if err := refreshTenant(ctx, tenant, client); err != nil {
		var resourceStateErr *tenantDetailResourceStateError
		if !errors.Is(err, errResourceNotFound) && !errors.As(err, &resourceStateErr) {
			return err
		}
		listErr := findTenant(ctx, tenant, client)
		if errors.Is(listErr, errResourceNotFound) {
			return listErr
		}
		if listErr != nil {
			return fmt.Errorf("%v; could not confirm tenant existence: %w", err, listErr)
		}
		if resourceStateErr != nil {
			return err
		}
		return errors.New("tenant detail reported it missing, but it is still present in the tenant list")
	}
	return nil
}

func waitForTenantCreation(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient, expected types.Bool, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	verifyAddress := !expected.IsNull() && !expected.IsUnknown()
	var mu sync.Mutex
	current := *tenant
	var observedPublicAddress *bool
	err := retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		mu.Lock()
		candidate := current
		mu.Unlock()
		if err := refreshTenant(ctx, &candidate, client); err != nil {
			return retry.RetryableError(fmt.Errorf("failed to refresh tenant: %w", err))
		}
		mu.Lock()
		current = candidate
		mu.Unlock()
		if !isTenantOnline(&candidate) {
			return retry.RetryableError(fmt.Errorf("tenant is not online: %s", candidate.Status.ValueString()))
		}
		if verifyAddress {
			enabled, err := tenantHasPublicAddress(ctx, &candidate, client)
			if err != nil {
				return retry.RetryableError(err)
			}
			mu.Lock()
			observedPublicAddress = &enabled
			mu.Unlock()
			if enabled != expected.ValueBool() {
				return retry.RetryableError(fmt.Errorf("public address presence is %t, expected %t", enabled, expected.ValueBool()))
			}
		}
		return nil
	})
	mu.Lock()
	defer mu.Unlock()
	*tenant = current
	if verifyAddress {
		if observedPublicAddress == nil {
			tenant.EnablePublicAddress = types.BoolNull()
		} else {
			tenant.EnablePublicAddress = types.BoolValue(*observedPublicAddress)
		}
	}
	return err
}

// Update reconciliation needs actual values for every supported operation. The
// SDK's getters return zero values for omitted fields, which are not evidence
// that a requested value has been reached.
func refreshTenantForUpdate(ctx context.Context, tenant *model.Tenant, client *obcloudsdk.APIClient) error {
	if tenant.Id.IsNull() || tenant.Id.IsUnknown() || tenant.Id.ValueString() == "" ||
		tenant.ClusterId.IsNull() || tenant.ClusterId.IsUnknown() || tenant.ClusterId.ValueString() == "" {
		return errors.New("tenant cluster id and id are required")
	}
	result, response, err := client.MultiCloudOpenAPI.DescribeTenant(ctx, tenant.ClusterId.ValueString(), tenant.Id.ValueString()).RequestId("").Execute()
	api.LogAPIResult(ctx, result, response, err)
	if err != nil || response == nil || response.StatusCode != http.StatusOK || result == nil || !result.GetSuccess() || !result.HasData() {
		return fmt.Errorf("failed to describe tenant: %v", err)
	}
	data := result.Data
	if data.GetTenantId() != tenant.Id.ValueString() {
		return errors.New("describe tenant returned a different or missing tenant id")
	}
	if !data.HasStatus() || !data.HasTenantName() || !data.HasPrimaryZone() || !data.HasUnitNum() || !data.HasUnitCpu() || !data.HasUnitMem() {
		return errors.New("describe tenant omitted fields needed to reconcile its update")
	}
	tenant.RefreshTenant(data)
	if tenant.EnablePublicAddress.IsNull() || tenant.EnablePublicAddress.IsUnknown() {
		enabled, err := tenantHasPublicAddress(ctx, tenant, client)
		if err != nil {
			return err
		}
		tenant.EnablePublicAddress = types.BoolValue(enabled)
	}
	return nil
}

func waitForTenantUpdate(ctx context.Context, tenant *model.Tenant, plan *model.Tenant, operation string, client *obcloudsdk.APIClient, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var mu sync.Mutex
	current := *tenant
	err := retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		mu.Lock()
		candidate := current
		mu.Unlock()
		if err := refreshTenantForUpdate(ctx, &candidate, client); err != nil {
			return retry.RetryableError(fmt.Errorf("failed to refresh tenant: %w", err))
		}
		mu.Lock()
		current = candidate
		mu.Unlock()
		matched := isTenantOnline(&candidate)
		switch operation {
		case OperationModifyTenantUnit:
			matched = matched && candidate.UnitNum.Equal(plan.UnitNum) && candidate.Unit.Equal(plan.Unit)
		case OperationModifyTenantName:
			matched = matched && candidate.Name.Equal(plan.Name)
		case OperationModifyTenantPrimaryZone:
			matched = matched && candidate.PrimaryZone.Equal(plan.PrimaryZone)
		}
		if !matched {
			return retry.RetryableError(fmt.Errorf("tenant has not reached the requested state, current status: %s", candidate.Status.ValueString()))
		}
		return nil
	})
	mu.Lock()
	*tenant = current
	mu.Unlock()
	return err
}
