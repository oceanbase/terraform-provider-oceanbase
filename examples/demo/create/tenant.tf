resource "oceanbase_tenant" "demo" {
  cluster_id            = oceanbase_cluster.demo.id
  name                  = var.tenant_name
  mode                  = var.tenant_mode
  primary_zone          = var.primary_zone
  charset               = var.tenant_charset
  timezone              = var.tenant_timezone
  unit_num              = var.tenant_unit_num
  unit                  = { cpu = var.tenant_unit_cpu, memory = var.tenant_unit_memory }
  enable_public_address = false
}
