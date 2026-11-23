terraform {
  required_providers {
    oceanbase = {
      source  = "oceanbase.github.io/oceanbase/oceanbase"
      version = ">= 1.1.0, < 2.0.0"
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

variable "site" {
  type    = string
  default = "CHINA"
}

provider "oceanbase" {
  access_key = var.oceanbase_access_key
  secret_key = var.oceanbase_secret_key
  site       = var.site
}

variable "project_id" { type = string }
variable "instance_name" { type = string }
variable "region" { type = string }
variable "zones" { type = set(string) }
variable "instance_class" { type = string }
variable "ob_version" { type = string }
variable "disk_size" { type = number }
variable "replica_mode" { type = string }
variable "node_num" { type = number }
variable "payment_type" { type = string }

# Match these settings to the existing instance before importing it with:
# terraform import oceanbase_cluster.example <instance-id>

resource "oceanbase_cluster" "example" {
  project_id     = var.project_id
  name           = var.instance_name
  cloud_provider = "ALIYUN"
  region         = var.region
  zones          = var.zones
  instance_class = var.instance_class
  version        = var.ob_version
  disk_size      = var.disk_size
  replica_mode   = var.replica_mode
  instance_type  = "CLUSTER"
  node_num       = var.node_num
  payment_type   = var.payment_type
}
