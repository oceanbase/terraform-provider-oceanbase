terraform {
  required_providers {
    oceanbase = {
      source  = "oceanbase.github.io/oceanbase/oceanbase"
    }
  }
}

provider "oceanbase" {
  access_key  = "mock_access_key"
  secret_key = "mock_secret_key"
  site = "MOCK"
}
