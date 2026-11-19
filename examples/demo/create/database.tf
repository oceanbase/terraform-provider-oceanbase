resource "oceanbase_database" "demo" {
  count = var.enable_database_user ? 1 : 0

  cluster_id  = oceanbase_cluster.demo.id
  tenant_id   = oceanbase_tenant.demo.id
  name        = "tf_demo_db"
  charset     = "utf8mb4"
  description = "Terraform demo database"
}

resource "oceanbase_tenant_user" "demo" {
  count = var.enable_database_user ? 1 : 0

  cluster_id = oceanbase_cluster.demo.id
  tenant_id  = oceanbase_tenant.demo.id
  name       = "tf_demo_user"
  password   = var.database_user_password
  type       = "Normal"

  privileges = [{
    database = oceanbase_database.demo[0].name
    role     = "readwrite"
  }]
}
