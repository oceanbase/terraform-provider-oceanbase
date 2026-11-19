package provider

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

var (
	_ resource.Resource                = &privateLinkServiceResource{}
	_ resource.ResourceWithConfigure   = &privateLinkServiceResource{}
	_ resource.ResourceWithImportState = &privateLinkServiceResource{}
)

type privateLinkServiceResource struct{ provider *oceanbaseProvider }

type privateLinkServiceModel struct {
	ID             types.String `tfsdk:"id"`
	ClusterID      types.String `tfsdk:"cluster_id"`
	TenantID       types.String `tfsdk:"tenant_id"`
	ProjectID      types.String `tfsdk:"project_id"`
	UserAccount    types.String `tfsdk:"user_account"`
	ServiceName    types.String `tfsdk:"service_name"`
	ServiceAccount types.String `tfsdk:"service_account"`
	Status         types.String `tfsdk:"status"`
}

func NewPrivateLinkServiceResource() resource.Resource { return &privateLinkServiceResource{} }

func (r *privateLinkServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_private_link_service"
}

func (r *privateLinkServiceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a tenant's primary PrivateLink service. The cloud endpoint is managed separately.",
		Attributes: map[string]schema.Attribute{
			"id":              schema.StringAttribute{Computed: true, MarkdownDescription: "OceanBase address ID."},
			"cluster_id":      schema.StringAttribute{Required: true, PlanModifiers: replace},
			"tenant_id":       schema.StringAttribute{Required: true, PlanModifiers: replace},
			"project_id":      schema.StringAttribute{Optional: true, PlanModifiers: replace},
			"user_account":    schema.StringAttribute{Required: true, PlanModifiers: replace, MarkdownDescription: "Cloud account allowed to connect to the service."},
			"service_name":    schema.StringAttribute{Computed: true, MarkdownDescription: "Cloud-facing service name or identifier."},
			"service_account": schema.StringAttribute{Computed: true, MarkdownDescription: "Cloud account that owns the service, when returned by the API."},
			"status":          schema.StringAttribute{Computed: true},
		},
	}
}

func (r *privateLinkServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error", fmt.Sprintf("Expected OceanBase provider data, got %T", req.ProviderData))
	}
}

func (m *privateLinkServiceModel) target() privateLinkTarget {
	return privateLinkTarget{clusterID: m.ClusterID.ValueString(), tenantID: m.TenantID.ValueString(), projectID: m.ProjectID.ValueString(), addressID: m.ID.ValueString()}
}

func observePrivateLinkService(entries []obcloudsdk.TenantPrivateLinkDTO, account string) (bool, string, string, string, error) {
	var found bool
	var status, serviceName, serviceAccount string
	for _, entry := range entries {
		entryStatus := entry.GetPrivateLinkStatus()
		if entryStatus == "" {
			return false, "", "", "", errors.New("the PrivateLink service response omitted its status")
		}
		if !privateLinkActive(entryStatus) {
			continue
		}
		if entry.GetPrivateLinkUserAccount() == "" {
			return false, "", "", "", errors.New("the PrivateLink service response omitted its cloud account")
		}
		if !strings.EqualFold(entry.GetPrivateLinkUserAccount(), account) {
			continue
		}
		found = true
		name := entry.GetPrivateLinkServiceId()
		if serviceName != "" && name != "" && serviceName != name {
			return false, "", "", "", errors.New("the account has multiple different PrivateLink services for this address")
		}
		owner := entry.GetPrivateLinkServiceAccount()
		if serviceAccount != "" && owner != "" && serviceAccount != owner {
			return false, "", "", "", errors.New("the PrivateLink service has multiple different cloud accounts")
		}
		if owner != "" {
			serviceAccount = owner
		}
		if name != "" {
			serviceName = name
			status = entryStatus
		} else if status == "" {
			status = entryStatus
		}
	}
	return found, status, serviceName, serviceAccount, nil
}

func (r *privateLinkServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var plan privateLinkServiceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if strings.TrimSpace(plan.UserAccount.ValueString()) == "" {
		resp.Diagnostics.AddError("Invalid cloud account", "user_account must not be empty.")
		return
	}
	target := plan.target()
	if err := target.validate(false); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	addresses, err := describePrivateLinks(ctx, r.provider, target)
	if err != nil && !errors.Is(err, errResourceNotFound) {
		resp.Diagnostics.AddError("Could not check PrivateLink service", err.Error())
		return
	}
	for _, entry := range privateLinkEntries(addresses) {
		if privateLinkActive(entry.GetPrivateLinkStatus()) {
			resp.Diagnostics.AddError("PrivateLink service already exists", "Import the existing service instead of creating another service for this tenant.")
			return
		}
	}
	addressID, err := createPrivateLinkService(ctx, r.provider, target, plan.UserAccount.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Could not create PrivateLink service", fmt.Sprintf("%v. Inspect the tenant before retrying; the request may have been accepted.", err))
		return
	}
	plan.ID = types.StringValue(addressID)
	plan.Status = types.StringValue("SERVICE_CREATING")
	plan.ServiceName = types.StringNull()
	plan.ServiceAccount = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	target = plan.target()
	timeout := operationTimeout(EnvResourceCreateTimeoutSeconds, DefaultResourceCreateTimeout)
	err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
		addresses, err := describePrivateLinks(ctx, r.provider, target)
		if errors.Is(err, errResourceNotFound) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		found, status, serviceName, serviceAccount, err := observePrivateLinkService(privateLinkEntries(addresses), plan.UserAccount.ValueString())
		if err != nil || !found {
			return false, err
		}
		plan.Status = types.StringValue(status)
		if strings.EqualFold(status, "SERVICE_CREATING") || serviceName == "" {
			return false, nil
		}
		switch strings.ToUpper(status) {
		case "SERVICE_CREATED", "CONNECTING", "WAIT_CONFIRMATION", "CONNECTED":
		default:
			return false, fmt.Errorf("unexpected PrivateLink service status %q", status)
		}
		plan.ServiceName = types.StringValue(serviceName)
		if serviceAccount != "" {
			plan.ServiceAccount = types.StringValue(serviceAccount)
		}
		return true, nil
	})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if err != nil {
		resp.Diagnostics.AddError("Waiting for PrivateLink service failed", fmt.Sprintf("Address %q was created but the service did not become ready: %v. Inspect the service before retrying.", addressID, err))
	}
}

func (r *privateLinkServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var state privateLinkServiceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := state.target().validate(true); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	addresses, err := describePrivateLinks(ctx, r.provider, state.target())
	if err != nil && !errors.Is(err, errResourceNotFound) {
		resp.Diagnostics.AddError("Could not read PrivateLink service", err.Error())
		return
	}
	found, status, serviceName, serviceAccount, err := observePrivateLinkService(privateLinkEntries(addresses), state.UserAccount.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Could not read PrivateLink service", err.Error())
		return
	}
	if !found {
		if strings.EqualFold(state.Status.ValueString(), "SERVICE_CREATING") {
			resp.Diagnostics.AddError("PrivateLink service is not yet observable", "The service was being created. Inspect it before retrying creation.")
			return
		}
		resp.State.RemoveResource(ctx)
		return
	}
	state.Status = types.StringValue(status)
	if serviceName != "" {
		state.ServiceName = types.StringValue(serviceName)
	} else {
		state.ServiceName = types.StringNull()
	}
	if serviceAccount != "" {
		state.ServiceAccount = types.StringValue(serviceAccount)
	} else {
		state.ServiceAccount = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *privateLinkServiceResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("PrivateLink service updates are not supported", "Changes to service inputs require replacement.")
}

func (r *privateLinkServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var state privateLinkServiceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	target := state.target()
	if err := target.validate(true); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	timeout := operationTimeout(EnvResourceDeleteTimeoutSeconds, DefaultResourceModifyTimeout)
	var shouldDelete bool
	err := waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
		addresses, err := describePrivateLinks(ctx, r.provider, target)
		if errors.Is(err, errResourceNotFound) {
			if strings.EqualFold(state.Status.ValueString(), "SERVICE_CREATING") {
				return false, errors.New("the service creation outcome is unknown; inspect it before removing state")
			}
			return true, nil
		}
		if err != nil {
			return false, err
		}
		entries := privateLinkEntries(addresses)
		found, status, _, _, err := observePrivateLinkService(entries, state.UserAccount.ValueString())
		if err != nil {
			return false, err
		}
		if !found {
			if strings.EqualFold(state.Status.ValueString(), "SERVICE_CREATING") {
				return false, errors.New("the service creation outcome is unknown; inspect it before removing state")
			}
			return true, nil
		}
		activeEntries := 0
		for _, entry := range entries {
			if privateLinkActive(entry.GetPrivateLinkStatus()) {
				activeEntries++
			}
			if privateLinkActive(entry.GetPrivateLinkStatus()) && entry.GetPrivateLinkEndpointId() != "" {
				return false, errors.New("the service still has a connected endpoint; remove its connections first")
			}
		}
		if activeEntries != 1 {
			return false, errors.New("the address has multiple PrivateLink entries; inspect them before deleting the service")
		}
		if strings.EqualFold(status, "SERVICE_CREATING") || strings.EqualFold(status, "DELETING") {
			return false, nil
		}
		if !strings.EqualFold(status, "SERVICE_CREATED") {
			return false, fmt.Errorf("cannot delete a PrivateLink service in status %q", status)
		}
		shouldDelete = true
		return true, nil
	})
	if err == nil && shouldDelete {
		err = deletePrivateLink(ctx, r.provider, target, "", "SERVICE_CREATED")
		if err == nil {
			err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
				addresses, err := describePrivateLinks(ctx, r.provider, target)
				if errors.Is(err, errResourceNotFound) {
					return true, nil
				}
				if err != nil {
					return false, err
				}
				found, _, _, _, err := observePrivateLinkService(privateLinkEntries(addresses), state.UserAccount.ValueString())
				return !found, err
			})
		}
		if err == nil {
			err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
				return privateLinkTenantReady(ctx, r.provider, target)
			})
		}
	}
	if err != nil {
		resp.Diagnostics.AddError("Could not delete PrivateLink service", err.Error())
	}
}

func privateLinkImportParts(id string, count int) ([]string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != count {
		return nil, fmt.Errorf("expected %d slash-separated values", count)
	}
	for i, part := range parts {
		value, err := url.PathUnescape(part)
		if err != nil || value == "" {
			return nil, errors.New("import values must be nonempty URL-escaped path segments")
		}
		parts[i] = value
	}
	return parts, nil
}

func (r *privateLinkServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts, err := privateLinkImportParts(req.ID, 4)
	if err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink service import ID", "Expected cluster_id/tenant_id/address_id/user_account: "+err.Error())
		return
	}
	for i, name := range []string{"cluster_id", "tenant_id", "id", "user_account"} {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(name), parts[i])...)
	}
}
