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

import "time"

// cloud providers
const (
	CloudProviderAliyun string = "ALIYUN"
	CloudProviderQCloud string = "QCLOUD"
	CloudProviderHuawei string = "HUAWEI"
	CloudProviderAWS    string = "AWS"
	CloudProviderGCP    string = "GCP"
)

const (
	EnvNameAccessKey                string = "OCEANBASE_ACCESS_KEY"
	EnvNameSecretKey                string = "OCEANBASE_SECRET_KEY"
	EnvNameSiteURL                  string = "OCEANBASE_SITE_URL"
	EnvClusterCreateTimeoutSeconds  string = "TF_OCEANBASE_CLUSTER_CREATE_TIMEOUT_SECONDS"
	EnvClusterModifyTimeoutSeconds  string = "TF_OCEANBASE_CLUSTER_MODIFY_TIMEOUT_SECONDS"
	EnvResourceCreateTimeoutSeconds string = "TF_OCEANBASE_RESOURCE_CREATE_TIMEOUT_SECONDS"
	EnvResourceModifyTimeoutSeconds string = "TF_OCEANBASE_RESOURCE_MODIFY_TIMEOUT_SECONDS"
)

const (
	DefaultClusterCreateTimeout  time.Duration = 90 * time.Minute
	DefaultClusterModifyTimeout  time.Duration = 90 * time.Minute
	DefaultResourceCreateTimeout time.Duration = 90 * time.Minute
	DefaultResourceModifyTimeout time.Duration = 90 * time.Minute
)

const (
	InstanceTypeCluster string = "CLUSTER"
)

const (
	// for cluster instance
	SpecNodeNum      string = "node_num"
	SpecInstanceSpec string = "instance_spec(instance_class or disk_size)"
	SpecIsStopped    string = "is_stopped"
	SpecInstanceName string = "instance_name"
	// for tenant
	SpecTenantUnit        string = "unit(unit spec and unit num)"
	SpecTenantName        string = "name"
	SpecTenantPrimaryZone string = "primary zone"
	SpecTenantTags        string = "tags"
	// for tenant user
	SpecTenantUserDescription  string = "description"
	SpecTenantUserPassword     string = "password"
	SpecTenantUserPrivileges   string = "privileges"
	SpecTenantUserLockedStatus string = "locked-status"
	// for database
	SpecDatabaseDescription string = "description"
)

const (
	// common
	OperationNothing     string = "nothing"
	OperationUnsupported string = "unsupported"
	// for cluster instance
	OperationModifyNodeNum      string = "modify_node_num"
	OperationModifyInstanceSpec string = "modify_instance_spec"
	OperationModifyInstanceName string = "modify_instance_name"
	OperationStartCluster       string = "start_cluster"
	OperationStopCluster        string = "stop_cluster"
	// for tenant
	OperationModifyTenantUnit        string = "modify_tenant_unit"
	OperationModifyTenantName        string = "modify_tenant_name"
	OperationModifyTenantPrimaryZone string = "modify_tenant_primary_zone"
	OperationModifyTenantTags        string = "modify_tenant_tags"
	// for tenant user
	OperationModifyTenantUserDescription  string = "modify_tenant_user_description"
	OperationModifyTenantUserPassword     string = "modify_tenant_user_password"
	OperationModifyTenantUserPrivileges   string = "modify_tenant_user_privileges"
	OperationModifyTenantUserLockedStatus string = "modify_tenant_user_status"
	// for tenant security group
	OperationModifyTenantSecurityGroupCidrs string = "modify_tenant_security_group_cidrs"
	// for database
	OperationModifyDatabaseDescription string = "modify_database_description"
)
