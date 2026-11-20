data "oceanbase_cluster" "demo" {
  id = var.cluster_id
}

data "oceanbase_tenant" "demo" {
  cluster_id = var.cluster_id
  id         = var.tenant_id
}

data "oceanbase_database" "demo" {
  count = var.database_name == "" ? 0 : 1

  cluster_id = var.cluster_id
  tenant_id  = var.tenant_id
  name       = var.database_name
}

data "oceanbase_tenant_user" "demo" {
  count = var.database_user == "" ? 0 : 1

  cluster_id = var.cluster_id
  tenant_id  = var.tenant_id
  name       = var.database_user
}

data "oceanbase_tenant_security_group" "demo" {
  count = var.security_group_name == "" ? 0 : 1

  cluster_id = var.cluster_id
  tenant_id  = var.tenant_id
  name       = var.security_group_name
}

data "oceanbase_tenant_connection" "internet" {
  count = var.internet_address_id == "" ? 0 : 1

  cluster_id = var.cluster_id
  tenant_id  = var.tenant_id
  type       = "INTERNET"
  id         = var.internet_address_id
}
