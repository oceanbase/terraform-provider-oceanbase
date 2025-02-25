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

resource "oceanbase_cluster" "obins" {
  cloud_provider = "ALIYUN"
  instance_class = "4C16G"
  version = "4.2.3"
  region = "cn-shanghai"
  zones = ["e", "f", "g"]
  instance_type = "CLUSTER"
  disk_size = 100
  replica_mode = 3
  node_num = 3
}

output "cluster_output" {
  value = resource.oceanbase_cluster.obins
}
