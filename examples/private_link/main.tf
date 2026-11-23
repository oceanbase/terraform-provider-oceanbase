terraform {
  required_version = ">= 1.5.0"
  required_providers {
    oceanbase = {
      source  = "oceanbase.github.io/oceanbase/oceanbase"
      version = ">= 1.1.0, < 2.0.0"
    }
    alicloud = {
      source  = "aliyun/alicloud"
      version = ">= 1.111.0, < 3.0.0"
    }
  }
}

variable "oceanbase_access_key" {
  type      = string
  sensitive = true
}

variable "oceanbase_secret_key" {
  type      = string
  sensitive = true
}

variable "oceanbase_site" {
  type    = string
  default = "CHINA"
}

variable "region" { type = string }

provider "oceanbase" {
  access_key = var.oceanbase_access_key
  secret_key = var.oceanbase_secret_key
  site       = var.oceanbase_site
}

provider "alicloud" {
  region = var.region
}

# The VPC and vSwitch belong to the cloud account identified by user_account.
variable "project_id" { type = string }
variable "cluster_id" { type = string }
variable "tenant_id" { type = string }
variable "user_account" { type = string }
variable "vpc_id" { type = string }
variable "vswitch_id" { type = string }
variable "zone_id" { type = string }
variable "security_group_ids" { type = set(string) }

variable "proxy_setting" {
  description = "ODP settings when service creation requires a new commercialized proxy; otherwise null."
  type = object({
    mode                    = string
    deployment_architecture = string
    zone                    = list(string)
    slave_zone              = optional(list(string), [])
    node_cpu                = number
  })
  default = null
}

resource "oceanbase_tenant_private_link_service" "primary" {
  project_id    = var.project_id
  cluster_id    = var.cluster_id
  tenant_id     = var.tenant_id
  user_account  = var.user_account
  proxy_setting = var.proxy_setting
}

resource "alicloud_privatelink_vpc_endpoint" "primary" {
  vpc_id             = var.vpc_id
  service_name       = oceanbase_tenant_private_link_service.primary.service_name
  security_group_ids = var.security_group_ids
}

resource "alicloud_privatelink_vpc_endpoint_zone" "primary" {
  endpoint_id = alicloud_privatelink_vpc_endpoint.primary.id
  vswitch_id  = var.vswitch_id
  zone_id     = var.zone_id
}

resource "oceanbase_tenant_private_link_connection" "primary" {
  project_id  = var.project_id
  cluster_id  = var.cluster_id
  tenant_id   = var.tenant_id
  address_id  = oceanbase_tenant_private_link_service.primary.id
  endpoint_id = alicloud_privatelink_vpc_endpoint.primary.id

  depends_on = [alicloud_privatelink_vpc_endpoint_zone.primary]
}

output "endpoint_id" {
  value = alicloud_privatelink_vpc_endpoint.primary.id
}

output "service_name" {
  value = oceanbase_tenant_private_link_service.primary.service_name
}

output "private_link_domain" {
  value = oceanbase_tenant_private_link_connection.primary.domain
}
