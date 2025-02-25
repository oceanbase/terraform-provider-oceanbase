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
