# Terraform Provider for OceanBase
This provider allows user to manage OceanBase resources on cloud environment using terraform, it's based on OceanBase's [multi-cloud platform](https://console-cn.oceanbase.com) APIs.

## Requirements
Go 1.22 (to build the provider)

## Building The Provider
Run the following command to build the provider
```
make build
```
## Using the provider
First, configure the `.terraformrc` file under user's home directory, make sure to the real path where the binary located.
```
provider_installation {

  dev_overrides {
      "oceanbase.github.io/oceanbase/oceanbase" = "~/terraform-provider-oceanbase/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

You can refer to the following config to use the provider.
```
terraform {
  required_providers {
    oceanbase = {
      source  = "oceanbase.github.io/oceanbase/oceanbase"
    }
  }
}

provider "oceanbase" {
  access_key  = "access_key"
  secret_key = "secret_key"
  site = "CHINA" # CHINA or GLOBAL
}
```
For OceanBase related resources configurations, please refer to the files under example directory.

## Cluster and tenant configuration

| Resource | Configuration |
| --- | --- |
| `oceanbase_cluster` | Set `payment_type`, `disk_type`, and `storage_architecture` when creating the cluster. `node_num` may be set on the initial apply or later; it must be a multiple of 3 between 3 and 60. If needed, scaling starts after the new cluster is online. |
| `oceanbase_tenant` | `enable_public_address` and `whitelist` are creation settings. Use `oceanbase_tenant_connection` and `oceanbase_tenant_security_group` to manage individual addresses and IP groups after creation. |
| `oceanbase_tenant_connection` | To create a public address, set `type = "INTERNET"`. Omit `private_link`, `enable_ssl`, and `enable_rpc`. The provider reads SSL and RPC settings from the service. |
| `oceanbase_tenant_security_group` | A single IPv4 address can be written as either `192.0.2.1` or `192.0.2.1/32`. Terraform preserves the configured spelling on refresh. |

You can change a cluster's name and size in one apply, or change a tenant's name, unit configuration and primary zone together. The provider submits each change after the previous one reaches its target. If a cloud operation times out, check the resource by its recorded ID and run `terraform plan` to see which changes are still needed.
