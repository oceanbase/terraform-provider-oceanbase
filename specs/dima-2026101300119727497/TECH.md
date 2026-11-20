# PrivateLink support in the OceanBase provider

## Context

Manage a tenant's primary PrivateLink service and its connection to an endpoint
created by the cloud provider. The existing OceanBase OpenAPI supplies service
creation, address description, endpoint connection, confirmation and deletion.
The current generated SDK uses an outdated deletion path and lacks per-request
project headers for these calls.

Domestic production reported OCP commit `438fda67910bd213bd1a74fcdba9189ad03a2cf0`
from `GetObCloudBuildInfo`; that commit contains the PrivateLink OpenAPI
controllers and the last-endpoint deletion behavior described below.

## Proposed changes

- Add an OceanBase PrivateLink service resource that creates and observes the
  service and exports the cloud-facing service name and address ID.
- Add an OceanBase PrivateLink connection resource that accepts an existing
  endpoint ID, connects and confirms it, and observes or removes that exact
  endpoint. The cloud provider owns the endpoint itself.
- Use the existing OpenAPI and pass the configured project ID. Treat the backend
  service as absent if its final endpoint deletion removes it. Do not promise
  independent preservation of the service.
- Preserve known identities when an accepted asynchronous operation has not
  completed, and reconcile existing remote state before a deletion or retry.
- Limit this implementation to a primary service and one endpoint connection.
  Additional independently managed endpoints need a reliable way to identify
  the backend's per-account pending service entries.

## Testing and validation

| Behavior | Evidence |
| --- | --- |
| Create, describe, connect, confirm and delete | Focused mock-HTTP tests check the actual method, path, request body and project header. |
| Asynchronous and missing resources | Focused tests cover pending, connected, failed, absent and unknown results. |
| Shared or unrelated endpoints | Deletion is only attempted after the target identity is observed; test the final-endpoint behavior. |
| Real environment | Use the current account's project and cloud-side test endpoint; verify creation, refresh, a no-change plan and cleanup. |

## Risks and follow-ups

The OCP last-endpoint branch decides whether to remove the service by count;
it does not first compare the requested endpoint ID. Provider-side checks do
not eliminate concurrent remote changes. Validate the production implementation
and report whether a minimal OCP fix is needed. Alibaba Cloud also requires an
eligible account and compatible NLB configuration. Production support remains
unverified until the cloud-side endpoint and required account access are known.
