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

variable "cluster_id" { type = string }
variable "tenant_id" { type = string }

variable "database_name" {
  type    = string
  default = ""
}

variable "database_user" {
  type    = string
  default = ""
}

variable "security_group_name" {
  type    = string
  default = ""
}

variable "internet_address_id" {
  type    = string
  default = ""
}
