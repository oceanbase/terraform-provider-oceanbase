resource "oceanbase_cluster" "demo" {
  project_id           = var.project_id
  name                 = var.instance_name
  cloud_provider       = "ALIYUN"
  region               = var.region
  zones                = var.zones
  instance_class       = var.instance_class
  version              = "4.4.2"
  disk_size            = var.disk_size
  node_num             = 3
  replica_mode         = "2"
  instance_type        = "CLUSTER"
  payment_type         = "PostPay"
  storage_architecture = "standardStorage"
  sale_channel         = "domestic"
}
