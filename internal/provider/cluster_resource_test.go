/*
 * Copyright 2025 OceanBase
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	resourceConfigInit = `
resource "oceanbase_cluster" "test" {
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
`
	resourceConfigUpdated = `
resource "oceanbase_cluster" "test" {
  cloud_provider = "ALIYUN"
  instance_class = "4C16G"
  version = "4.2.3"
  region = "cn-shanghai"
  zones = ["e", "f", "g"]
  instance_type = "CLUSTER"
  disk_size = 100
  replica_mode = 3
  node_num = 6
}
`
)

func TestAccOceanbaseResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + resourceConfigInit,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("resource.oceanbase_cluster.test", "node_num", "3"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + resourceConfigUpdated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("resource.oceanbase_cluster.test", "node_num", "6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
