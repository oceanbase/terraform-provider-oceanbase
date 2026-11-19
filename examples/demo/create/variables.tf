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

variable "project_id" { type = string }

variable "instance_name" {
  type    = string
  default = "tf_demo_cluster"
}

variable "tenant_name" {
  type    = string
  default = "tf_demo_tenant"
}

variable "region" {
  type    = string
  default = "cn-hangzhou"
}

variable "zones" {
  type    = set(string)
  default = ["cn-hangzhou-j", "cn-hangzhou-k"]
}

variable "primary_zone" {
  type    = string
  default = "cn-hangzhou-j"
}

variable "instance_class" {
  type    = string
  default = "4C16G"
}

variable "disk_size" {
  type    = number
  default = 50
}

variable "enable_internet" {
  type    = bool
  default = false
}

variable "allowed_ips" {
  description = "Client IPs or CIDRs for the demo security group."
  type        = list(string)
  default     = []
}

variable "enable_database_user" {
  type    = bool
  default = false
}

variable "database_user_password" {
  type      = string
  sensitive = true
  default   = null
}
