variable "oceanbase_access_key" {
  description = "目标 OceanBase Cloud 环境的 AccessKey ID，建议使用 TF_VAR_oceanbase_access_key 输入。"
  type        = string
  sensitive   = true
}

variable "oceanbase_secret_key" {
  description = "目标 OceanBase Cloud 环境的 AccessKey Secret，建议使用 TF_VAR_oceanbase_secret_key 输入。"
  type        = string
  sensitive   = true
}

variable "site" {
  description = "Provider 站点：CHINA、GLOBAL；MOCK 仅供本机模拟服务使用。"
  type        = string
  default     = "CHINA"

  validation {
    condition     = contains(["CHINA", "GLOBAL", "MOCK"], var.site)
    error_message = "site 必须是 CHINA、GLOBAL 或 MOCK。"
  }
}

variable "project_id" {
  description = "目标环境中创建实例的 OceanBase 项目 ID。"
  type        = string
}

variable "instance_name" {
  description = "演示实例名称。"
  type        = string
  default     = "tf_demo_cluster"
}

variable "tenant_name" {
  description = "演示租户名称。"
  type        = string
  default     = "tf_demo_tenant"
}

variable "cloud_provider" {
  description = "实例所在的云厂商标识，例如 ALIYUN。"
  type        = string
  default     = "ALIYUN"
}

variable "region" {
  description = "云厂商地域 ID；必须与所选环境及可用区匹配。"
  type        = string
  default     = "cn-hangzhou"
}

variable "zones" {
  description = "实例使用的云厂商可用区 ID 列表。"
  type        = set(string)
  default     = ["cn-hangzhou-j", "cn-hangzhou-k"]
}

variable "primary_zone" {
  description = "租户主可用区，应与实例可用区匹配。"
  type        = string
  default     = "cn-hangzhou-j"
}

variable "instance_class" {
  description = "实例规格，例如 4C16G。"
  type        = string
  default     = "4C16G"
}

variable "ob_version" {
  description = "OceanBase 数据库版本，不是 Terraform Provider 版本。"
  type        = string
  default     = "4.4.2"
}

variable "disk_size" {
  description = "实例存储空间，单位 GB。"
  type        = number
  default     = 50
}

variable "disk_type" {
  description = "云盘类型；null 表示不指定，由目标环境选择默认值。"
  type        = string
  default     = null
}

variable "node_num" {
  description = "目标节点数；不适用的实例类型可设为 null。"
  type        = number
  default     = 3
}

variable "replica_mode" {
  description = "实例副本模式，具体取值由目标环境的产品规格决定。"
  type        = string
  default     = "2"
}

variable "instance_type" {
  description = "实例类型，例如 CLUSTER。"
  type        = string
  default     = "CLUSTER"
}

variable "storage_architecture" {
  description = "存储架构；null 表示不指定。"
  type        = string
  default     = "standardStorage"
}

variable "sale_channel" {
  description = "目标环境支持的售卖渠道；null 表示不指定。"
  type        = string
  default     = "domestic"
}

variable "tenant_mode" {
  description = "租户兼容模式；默认 MySQL。"
  type        = string
  default     = "MySQL"
}

variable "tenant_charset" {
  description = "租户字符集。"
  type        = string
  default     = "utf8mb4"
}

variable "tenant_timezone" {
  description = "租户时区。"
  type        = string
  default     = "Asia/Shanghai"
}

variable "tenant_unit_num" {
  description = "租户 Unit 数量。"
  type        = number
  default     = 1
}

variable "tenant_unit_cpu" {
  description = "每个租户 Unit 的 CPU 核数。"
  type        = number
  default     = 1
}

variable "tenant_unit_memory" {
  description = "每个租户 Unit 的内存，单位 GB。"
  type        = number
  default     = 4
}

variable "enable_internet" {
  description = "是否创建演示 IP 白名单组和公网地址。"
  type        = bool
  default     = false
}

variable "security_group_name" {
  description = "演示 IP 白名单组名称。"
  type        = string
  default     = "tf_demo_ips"
}

variable "allowed_ips" {
  description = "允许访问的客户端 IP 或 CIDR；开启公网演示时应填写。"
  type        = list(string)
  default     = []
}

variable "enable_database_user" {
  description = "是否创建演示数据库及普通账号。"
  type        = bool
  default     = false
}

variable "database_name" {
  description = "演示数据库名称。"
  type        = string
  default     = "tf_demo_db"
}

variable "database_charset" {
  description = "演示数据库字符集。"
  type        = string
  default     = "utf8mb4"
}

variable "database_user_name" {
  description = "演示数据库普通账号名称。"
  type        = string
  default     = "tf_demo_user"
}

variable "database_user_role" {
  description = "普通账号对演示数据库的权限角色，例如 readwrite 或 readonly。"
  type        = string
  default     = "readwrite"
}

variable "database_user_password" {
  description = "演示数据库账号密码；开启账号时通过 TF_VAR_database_user_password 输入。"
  type        = string
  sensitive   = true
  default     = null
}
