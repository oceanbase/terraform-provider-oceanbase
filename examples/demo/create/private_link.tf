# OceanBase 侧先创建服务；云厂商侧的 Endpoint 可稍后创建并传入。
resource "oceanbase_tenant_private_link_service" "demo" {
  count = var.enable_private_link ? 1 : 0

  project_id   = var.project_id
  cluster_id   = oceanbase_cluster.demo.id
  tenant_id    = oceanbase_tenant.demo.id
  user_account = var.private_link_user_account

  lifecycle {
    precondition {
      condition     = length(trimspace(var.private_link_user_account)) > 0
      error_message = "启用 PrivateLink 时，请填写允许接入的云厂商主账号 ID。"
    }
  }
}

resource "oceanbase_tenant_private_link_connection" "demo" {
  count = var.enable_private_link && var.private_link_endpoint_id != "" ? 1 : 0

  project_id  = var.project_id
  cluster_id  = oceanbase_cluster.demo.id
  tenant_id   = oceanbase_tenant.demo.id
  address_id  = oceanbase_tenant_private_link_service.demo[0].id
  endpoint_id = var.private_link_endpoint_id
}
