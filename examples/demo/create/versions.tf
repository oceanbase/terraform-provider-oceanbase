terraform {
  required_providers {
    oceanbase = {
      source = "oceanbase.github.io/oceanbase/oceanbase"
    }
  }
}

provider "oceanbase" {
  access_key = var.oceanbase_access_key
  secret_key = var.oceanbase_secret_key
  site       = var.site
}
