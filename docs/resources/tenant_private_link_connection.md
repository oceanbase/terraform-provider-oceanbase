# oceanbase_tenant_private_link_connection

Connects an existing cloud PrivateLink endpoint to an OceanBase tenant's
primary PrivateLink service. The endpoint itself remains under the cloud
provider's control. The provider waits for the connection and confirms it when
the OceanBase API reports that confirmation is needed.

Arguments: `cluster_id`, `tenant_id`, `address_id` and `endpoint_id` are
required. `project_id` and `user_endpoint_address` are optional. Changing an
input replaces the resource. The observed `status`, `endpoint_address` and
`domain` are computed.

Import an existing connection with `cluster_id/tenant_id/address_id/endpoint_id`;
escape each component as a URL path segment. The configured project ID and
`user_endpoint_address` cannot be recovered from the connection response.

OceanBase may delete the associated service when this is the last connection.
The provider checks that the endpoint exists before requesting deletion; do not
change the same PrivateLink concurrently from another client while deleting it.
