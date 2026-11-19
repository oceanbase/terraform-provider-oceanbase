output "cluster_id" { value = oceanbase_cluster.demo.id }
output "tenant_id" { value = oceanbase_tenant.demo.id }
output "database_name" { value = try(oceanbase_database.demo[0].name, null) }
output "database_user" { value = try(oceanbase_tenant_user.demo[0].name, null) }
output "internet_address_id" { value = try(oceanbase_tenant_connection.internet[0].id, null) }
output "internet_address" { value = try(oceanbase_tenant_connection.internet[0].internet.address, null) }
output "internet_port" { value = try(oceanbase_tenant_connection.internet[0].internet.port, null) }
