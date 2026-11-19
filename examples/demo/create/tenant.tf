resource "oceanbase_tenant" "demo" {
  cluster_id            = oceanbase_cluster.demo.id
  name                  = var.tenant_name
  mode                  = "MySQL"
  primary_zone          = var.primary_zone
  charset               = "utf8mb4"
  timezone              = "Asia/Shanghai"
  unit_num              = 1
  unit                  = { cpu = 1, memory = 4 }
  enable_public_address = false
}
