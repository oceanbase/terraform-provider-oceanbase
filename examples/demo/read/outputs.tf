output "cluster_status" { value = data.oceanbase_cluster.demo.status }
output "tenant_status" { value = data.oceanbase_tenant.demo.status }
output "database_status" { value = try(data.oceanbase_database.demo[0].status, null) }
output "database_user_status" { value = try(data.oceanbase_tenant_user.demo[0].status, null) }
output "security_group_cidrs" { value = try(data.oceanbase_tenant_security_group.demo[0].cidrs, null) }
output "internet_status" { value = try(data.oceanbase_tenant_connection.internet[0].status, null) }
