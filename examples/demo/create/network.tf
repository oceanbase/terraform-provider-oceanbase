resource "oceanbase_tenant_security_group" "demo" {
  count = var.enable_internet ? 1 : 0

  cluster_id = oceanbase_cluster.demo.id
  tenant_id  = oceanbase_tenant.demo.id
  name       = var.security_group_name
  cidrs      = var.allowed_ips
}

resource "oceanbase_tenant_connection" "internet" {
  count = var.enable_internet ? 1 : 0

  cluster_id = oceanbase_cluster.demo.id
  tenant_id  = oceanbase_tenant.demo.id
  type       = "INTERNET"

  depends_on = [oceanbase_tenant_security_group.demo]
}
