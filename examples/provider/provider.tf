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
