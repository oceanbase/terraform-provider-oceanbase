package provider

import (
	"context"
	"errors"
	"fmt"
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
	_ resource.Resource                = &privateLinkConnectionResource{}
	_ resource.ResourceWithConfigure   = &privateLinkConnectionResource{}
	_ resource.ResourceWithImportState = &privateLinkConnectionResource{}
)

type privateLinkConnectionResource struct{ provider *oceanbaseProvider }

type privateLinkConnectionModel struct {
	ID                  types.String `tfsdk:"id"`
	ClusterID           types.String `tfsdk:"cluster_id"`
	TenantID            types.String `tfsdk:"tenant_id"`
	ProjectID           types.String `tfsdk:"project_id"`
	AddressID           types.String `tfsdk:"address_id"`
	EndpointID          types.String `tfsdk:"endpoint_id"`
	UserEndpointAddress types.String `tfsdk:"user_endpoint_address"`
	EndpointAddress     types.String `tfsdk:"endpoint_address"`
	Domain              types.String `tfsdk:"domain"`
	Status              types.String `tfsdk:"status"`
}

func NewPrivateLinkConnectionResource() resource.Resource { return &privateLinkConnectionResource{} }

func (r *privateLinkConnectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant_private_link_connection"
}

func (r *privateLinkConnectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connects an existing cloud endpoint to a tenant's PrivateLink service.",
		Attributes: map[string]schema.Attribute{
			"id":                    schema.StringAttribute{Computed: true, MarkdownDescription: "Cloud endpoint ID."},
			"cluster_id":            schema.StringAttribute{Required: true, PlanModifiers: replace},
			"tenant_id":             schema.StringAttribute{Required: true, PlanModifiers: replace},
			"project_id":            schema.StringAttribute{Optional: true, PlanModifiers: replace},
			"address_id":            schema.StringAttribute{Required: true, PlanModifiers: replace, MarkdownDescription: "OceanBase PrivateLink service address ID."},
			"endpoint_id":           schema.StringAttribute{Required: true, PlanModifiers: replace, MarkdownDescription: "Endpoint ID created by the cloud provider."},
			"user_endpoint_address": schema.StringAttribute{Optional: true, PlanModifiers: replace, MarkdownDescription: "Optional endpoint address to send when connecting."},
			"endpoint_address":      schema.StringAttribute{Computed: true},
			"domain":                schema.StringAttribute{Computed: true},
			"status":                schema.StringAttribute{Computed: true},
		},
	}
}

func (r *privateLinkConnectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	var ok bool
	if r.provider, ok = req.ProviderData.(*oceanbaseProvider); !ok {
		resp.Diagnostics.AddError("Internal provider error", fmt.Sprintf("Expected OceanBase provider data, got %T", req.ProviderData))
	}
}

func (m *privateLinkConnectionModel) target() privateLinkTarget {
	return privateLinkTarget{clusterID: m.ClusterID.ValueString(), tenantID: m.TenantID.ValueString(), projectID: m.ProjectID.ValueString(), addressID: m.AddressID.ValueString()}
}

func (m *privateLinkConnectionModel) validateState() error {
	if err := m.target().validate(true); err != nil {
		return err
	}
	if strings.TrimSpace(m.EndpointID.ValueString()) == "" || m.ID.ValueString() != m.EndpointID.ValueString() {
		return errors.New("the connection ID must match a nonempty endpoint_id")
	}
	return nil
}

func findPrivateLinkEndpoint(entries []obcloudsdk.TenantPrivateLinkDTO, endpointID string) (*obcloudsdk.TenantPrivateLinkDTO, error) {
	var match *obcloudsdk.TenantPrivateLinkDTO
	for i := range entries {
		if entries[i].GetPrivateLinkEndpointId() != endpointID {
			continue
		}
		if entries[i].GetPrivateLinkStatus() == "" {
			return nil, fmt.Errorf("the endpoint %q has no PrivateLink status", endpointID)
		}
		if !privateLinkActive(entries[i].GetPrivateLinkStatus()) {
			continue
		}
		if match != nil {
			return nil, fmt.Errorf("the endpoint %q appeared more than once in the PrivateLink response", endpointID)
		}
		match = &entries[i]
	}
	return match, nil
}

func setPrivateLinkConnectionObservation(state *privateLinkConnectionModel, entry *obcloudsdk.TenantPrivateLinkDTO) {
	state.Status = types.StringValue(entry.GetPrivateLinkStatus())
	if value := entry.GetIntranetAddress(); value != "" {
		state.EndpointAddress = types.StringValue(value)
	} else {
		state.EndpointAddress = types.StringNull()
	}
	if value := entry.GetPrivateLinkDomain(); value != "" {
		state.Domain = types.StringValue(value)
	} else {
		state.Domain = types.StringNull()
	}
}

func (r *privateLinkConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var plan privateLinkConnectionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if strings.TrimSpace(plan.EndpointID.ValueString()) == "" || strings.TrimSpace(plan.AddressID.ValueString()) == "" {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", "address_id and endpoint_id must not be empty.")
		return
	}
	target := plan.target()
	if err := target.validate(true); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	addresses, err := describePrivateLinks(ctx, r.provider, target)
	if err != nil {
		resp.Diagnostics.AddError("Could not check PrivateLink endpoint", err.Error())
		return
	}
	entries := privateLinkEntries(addresses)
	entry, err := findPrivateLinkEndpoint(entries, plan.EndpointID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Could not check PrivateLink endpoint", err.Error())
		return
	}
	if entry != nil {
		resp.Diagnostics.AddError("PrivateLink endpoint already exists", "Import the existing endpoint connection instead of connecting it again.")
		return
	}
	serviceReady := false
	for _, entry := range entries {
		if strings.EqualFold(entry.GetPrivateLinkStatus(), "SERVICE_CREATED") && entry.GetPrivateLinkServiceId() != "" {
			serviceReady = true
			break
		}
	}
	if !serviceReady {
		resp.Diagnostics.AddError("PrivateLink service is not ready", "The address must have a created service before connecting an endpoint.")
		return
	}
	plan.ID = plan.EndpointID
	plan.Status = types.StringValue("CONNECTING")
	plan.EndpointAddress = types.StringNull()
	plan.Domain = types.StringNull()
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := connectPrivateLink(ctx, r.provider, target, plan.EndpointID.ValueString(), plan.UserEndpointAddress.ValueString()); err != nil {
		resp.Diagnostics.AddError("Could not connect PrivateLink endpoint", fmt.Sprintf("%v. Inspect the endpoint before retrying; the request may have been accepted.", err))
		return
	}
	timeout := operationTimeout(EnvResourceCreateTimeoutSeconds, DefaultResourceCreateTimeout)
	confirmed := false
	err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
		addresses, err := describePrivateLinks(ctx, r.provider, target)
		if errors.Is(err, errResourceNotFound) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		entry, err := findPrivateLinkEndpoint(privateLinkEntries(addresses), plan.EndpointID.ValueString())
		if err != nil || entry == nil {
			return false, err
		}
		setPrivateLinkConnectionObservation(&plan, entry)
		switch strings.ToUpper(entry.GetPrivateLinkStatus()) {
		case "CONNECTED":
			return true, nil
		case "WAIT_CONFIRMATION":
			if !confirmed {
				confirmed = true
				if err := confirmPrivateLink(ctx, r.provider, target, plan.EndpointID.ValueString()); err != nil {
					return false, err
				}
			}
			return false, nil
		case "CONNECTING", "SERVICE_CREATED":
			return false, nil
		case "CONNECT_FAIL":
			return false, errors.New("the endpoint connection failed")
		default:
			return false, fmt.Errorf("unexpected PrivateLink endpoint status %q", entry.GetPrivateLinkStatus())
		}
	})
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if err != nil {
		resp.Diagnostics.AddError("Waiting for PrivateLink endpoint failed", fmt.Sprintf("Endpoint %q was submitted but did not become connected: %v. Inspect it before retrying.", plan.EndpointID.ValueString(), err))
	}
}

func (r *privateLinkConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var state privateLinkConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := state.validateState(); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	addresses, err := describePrivateLinks(ctx, r.provider, state.target())
	if err != nil && !errors.Is(err, errResourceNotFound) {
		resp.Diagnostics.AddError("Could not read PrivateLink endpoint", err.Error())
		return
	}
	entry, err := findPrivateLinkEndpoint(privateLinkEntries(addresses), state.EndpointID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Could not read PrivateLink endpoint", err.Error())
		return
	}
	if entry == nil {
		if !strings.EqualFold(state.Status.ValueString(), "CONNECTED") {
			resp.Diagnostics.AddError("PrivateLink endpoint is not yet observable", "The connection was not previously confirmed. Inspect it before retrying.")
			return
		}
		resp.State.RemoveResource(ctx)
		return
	}
	setPrivateLinkConnectionObservation(&state, entry)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *privateLinkConnectionResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("PrivateLink connection updates are not supported", "Changes to connection inputs require replacement.")
}

func (r *privateLinkConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	if r.provider == nil || !r.provider.configured {
		addConfigureProviderErr(&resp.Diagnostics)
		return
	}
	var state privateLinkConnectionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := state.validateState(); err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink identity", err.Error())
		return
	}
	target := state.target()
	timeout := operationTimeout(EnvResourceDeleteTimeoutSeconds, DefaultResourceModifyTimeout)
	var entryToDelete *obcloudsdk.TenantPrivateLinkDTO
	var deletionPending bool
	err := waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
		addresses, err := describePrivateLinks(ctx, r.provider, target)
		if errors.Is(err, errResourceNotFound) {
			if !strings.EqualFold(state.Status.ValueString(), "CONNECTED") {
				return false, errors.New("the endpoint was not previously confirmed and its outcome is unknown; inspect it before removing state")
			}
			return true, nil
		}
		if err != nil {
			return false, err
		}
		entry, err := findPrivateLinkEndpoint(privateLinkEntries(addresses), state.EndpointID.ValueString())
		if err != nil {
			return false, err
		}
		if entry == nil {
			if !strings.EqualFold(state.Status.ValueString(), "CONNECTED") {
				return false, errors.New("the endpoint was not previously confirmed and its outcome is unknown; inspect it before removing state")
			}
			return true, nil
		}
		switch strings.ToUpper(entry.GetPrivateLinkStatus()) {
		case "CONNECTING":
			return false, nil
		case "DELETING":
			deletionPending = true
			return false, nil
		case "CONNECTED", "WAIT_CONFIRMATION", "CONNECT_FAIL", "SERVICE_CREATED":
			entryToDelete = entry
			return true, nil
		default:
			return false, fmt.Errorf("cannot delete an endpoint in status %q", entry.GetPrivateLinkStatus())
		}
	})
	if err == nil && entryToDelete != nil {
		err = deletePrivateLink(ctx, r.provider, target, state.EndpointID.ValueString(), entryToDelete.GetPrivateLinkStatus())
		if err == nil {
			deletionPending = true
			err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
				addresses, err := describePrivateLinks(ctx, r.provider, target)
				if errors.Is(err, errResourceNotFound) {
					return true, nil
				}
				if err != nil {
					return false, err
				}
				entry, err := findPrivateLinkEndpoint(privateLinkEntries(addresses), state.EndpointID.ValueString())
				return entry == nil, err
			})
		}
	}
	if err == nil && deletionPending {
		err = waitPrivateLink(ctx, timeout, func(ctx context.Context) (bool, error) {
			return privateLinkTenantReady(ctx, r.provider, target)
		})
	}
	if err != nil {
		resp.Diagnostics.AddError("Could not delete PrivateLink endpoint", err.Error())
	}
}

func (r *privateLinkConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts, err := privateLinkImportParts(req.ID, 4)
	if err != nil {
		resp.Diagnostics.AddError("Invalid PrivateLink connection import ID", "Expected cluster_id/tenant_id/address_id/endpoint_id: "+err.Error())
		return
	}
	for i, name := range []string{"cluster_id", "tenant_id", "address_id", "endpoint_id"} {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(name), parts[i])...)
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[3])...)
}
