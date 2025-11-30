output "cluster_id" {
  description = "Aurora cluster identifier"
  value       = aws_rds_cluster.main.id
}

output "cluster_arn" {
  description = "Aurora cluster ARN"
  value       = aws_rds_cluster.main.arn
}

output "cluster_endpoint" {
  description = "Aurora cluster endpoint (writer)"
  value       = aws_rds_cluster.main.endpoint
}

output "cluster_reader_endpoint" {
  description = "Aurora cluster reader endpoint"
  value       = aws_rds_cluster.main.reader_endpoint
}

output "cluster_port" {
  description = "Aurora cluster port"
  value       = aws_rds_cluster.main.port
}

output "database_name" {
  description = "Database name"
  value       = aws_rds_cluster.main.database_name
}

output "master_username" {
  description = "Master username"
  value       = aws_rds_cluster.main.master_username
}

output "master_password_secret_arn" {
  description = "Secrets Manager secret ARN for master password"
  value       = aws_secretsmanager_secret.master_password.arn
}

output "instance_identifiers" {
  description = "Aurora instance identifiers"
  value       = aws_rds_cluster_instance.main[*].identifier
}

output "cluster_resource_id" {
  description = "Aurora cluster resource ID (for IAM auth)"
  value       = aws_rds_cluster.main.cluster_resource_id
}
