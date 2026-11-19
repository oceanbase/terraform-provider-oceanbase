# oceanbase_tenant_private_link_service

Creates the primary PrivateLink service for a tenant using the OceanBase
OpenAPI. The cloud account in `user_account` is allowed to connect to the
service. Use `service_name` and, where needed, `service_account` to create an
endpoint with the appropriate cloud provider. Pass this resource's `id` as the
connection's `address_id`.

See [the Alibaba Cloud example](../../examples/private_link/main.tf). Cloud
provider, region, account permissions and availability are subject to the
OceanBase service and the cloud provider. Alibaba Cloud PrivateLink also
requires an eligible account and compatible network load balancer.

Arguments: `cluster_id`, `tenant_id` and `user_account` are required;
`project_id` is optional. Changing an input replaces the resource.

The provider does not adopt an existing service during creation. Import an
existing service with `cluster_id/tenant_id/address_id/user_account`; escape
each component as a URL path segment. If connections still exist, delete them
before deleting the service.

OceanBase may remove the service when its last endpoint connection is deleted.
If this resource remains configured, a later plan can propose to recreate it.
The current implementation covers a primary service and one endpoint connection;
it does not provide independent service retention or manage cloud endpoints.
