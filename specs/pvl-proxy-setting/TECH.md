# PrivateLink ODP proxy settings

## Context

The OceanBase `CreatePrivatelinkService` OpenAPI accepts `proxySetting` when creating a new ODP and NLB. In domestic PRE, an ODP-commercialized tenant without a private main address needs this value. The provider's older request model cannot send it.

## Proposed changes

- Add an optional `proxy_setting` object to `oceanbase_tenant_private_link_service`; include `mode`, `deployment_architecture`, `zone`, `slave_zone` and `node_cpu`.
- Pass the object as `proxySetting` in the existing OpenAPI request. Omit the object when not configured; serialize an omitted `slave_zone` as an empty list when the object is present.
- Treat a change to this create-time input as replacement. Keep it in Terraform state because the service response does not return it; document that an imported service cannot recover the setting.
- Expose the optional setting in the existing demo without changing defaults.

## Testing and validation

| Requirement | Verification |
| --- | --- |
| Existing users omit the setting | Existing PrivateLink lifecycle test |
| Configured values reach OpenAPI unchanged, including an empty `slaveZone` | Focused request serialization and resource test |
| No unrelated instance or tenant changes | PRE Terraform plan before apply |
| Create path works in PRE | Reconcile existing service first, apply once, then read back service and plan |

## Risks and mitigations

Cloud/region and ODP availability, quota, deployment architecture, and billing are validated by OceanBase. Do not assume a particular setting is portable across environments. A failed or timed-out create must be reconciled before another request.

## Follow-ups

Cloud-provider endpoint creation and end-to-end database connectivity remain separate steps.
