# The OceanBase and cloud providers must be configured separately. The VPC and
# vSwitch below belong to the cloud account identified by user_account.
variable "project_id" { type = string }
variable "cluster_id" { type = string }
variable "tenant_id" { type = string }
variable "user_account" { type = string }
variable "vpc_id" { type = string }
variable "vswitch_id" { type = string }
variable "zone_id" { type = string }

resource "oceanbase_tenant_private_link_service" "primary" {
  project_id   = var.project_id
  cluster_id   = var.cluster_id
  tenant_id    = var.tenant_id
  user_account = var.user_account
}

resource "alicloud_privatelink_vpc_endpoint" "primary" {
  vpc_id       = var.vpc_id
  service_name = oceanbase_tenant_private_link_service.primary.service_name
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
