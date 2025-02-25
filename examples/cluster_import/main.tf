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

resource "oceanbase_cluster" "test" {
  cloud_provider = "ALIYUN"
  instance_class = "8C32G"
  version = "4.2.1"
  region = "cn-hangzhou"
  zones = ["h", "j", "k"]
  disk_size = 100
  replica_mode = 3
  instance_type = "CLUSTER"
  node_num = 6
  sale_channel = "normandy_domestic"
}
