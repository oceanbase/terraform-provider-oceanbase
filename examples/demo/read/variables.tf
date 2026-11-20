variable "oceanbase_access_key" {
  description = "目标环境的 OceanBase Cloud AccessKey ID。"
  type        = string
  sensitive   = true
}

variable "oceanbase_secret_key" {
  description = "目标环境的 OceanBase Cloud AccessKey Secret。"
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

variable "cluster_id" {
  description = "目标环境中已有的实例 ID。"
  type        = string
}

variable "tenant_id" {
  description = "该实例下已有的租户 ID。"
  type        = string
}

variable "database_name" {
  description = "需要查询的已有数据库名；空字符串表示跳过。"
  type        = string
  default     = ""
}

variable "database_user" {
  description = "需要查询的已有数据库账号；空字符串表示跳过。"
  type        = string
  default     = ""
}

variable "security_group_name" {
  description = "需要查询的已有 IP 白名单组名；空字符串表示跳过。"
  type        = string
  default     = ""
}

variable "internet_address_id" {
  description = "需要查询的已有公网地址 ID；空字符串表示跳过。"
  type        = string
  default     = ""
}
