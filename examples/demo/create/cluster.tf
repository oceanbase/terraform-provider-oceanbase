resource "oceanbase_cluster" "demo" {
  project_id           = var.project_id
  name                 = var.instance_name
  cloud_provider       = var.cloud_provider
  region               = var.region
  zones                = var.zones
  instance_class       = var.instance_class
  version              = var.ob_version
  disk_size            = var.disk_size
  disk_type            = var.disk_type
  node_num             = var.node_num
  replica_mode         = var.replica_mode
  instance_type        = var.instance_type
  payment_type         = "PostPay"
  storage_architecture = var.storage_architecture
  sale_channel         = var.sale_channel
}
