# Production Environment Variables
# Production-grade defaults with high availability and security

variable "project" {
  type = string
}

variable "environment" {
  type    = string
  default = "prd"
}

variable "aws_region" {
  type    = string
  default = "ap-northeast-1"
}

# VPC
variable "vpc_cidr" {
  type    = string
  default = "10.2.0.0/16" # Different CIDR for prd
}

variable "az_count" {
  type    = number
  default = 3 # 3 AZs for production
}

variable "enable_nat_gateway" {
  type    = bool
  default = true
}

variable "single_nat_gateway" {
  type    = bool
  default = false # Multi-NAT for production HA
}

variable "enable_flow_logs" {
  type    = bool
  default = true
}

variable "enable_vpc_endpoints" {
  type    = bool
  default = true
}

# Application
variable "app_port" {
  type    = number
  default = 8080
}

variable "health_check_path" {
  type    = string
  default = "/health"
}

# Database
variable "database_name" {
  type = string
}

variable "database_master_username" {
  type    = string
  default = "postgres"
}

variable "database_app_username" {
  type    = string
  default = "app_user"
}

variable "aurora_engine_version" {
  type    = string
  default = "16.4"
}

variable "aurora_engine_version_major" {
  type    = string
  default = "16"
}

variable "aurora_instance_count" {
  type    = number
  default = 3 # 3 instances for production (multi-AZ)
}

variable "aurora_min_capacity" {
  type    = number
  default = 2 # Higher minimum for production
}

variable "aurora_max_capacity" {
  type    = number
  default = 64 # Higher maximum for production
}

variable "aurora_backup_retention_period" {
  type    = number
  default = 35 # Maximum retention for production
}

variable "aurora_deletion_protection" {
  type    = bool
  default = true
}

variable "aurora_skip_final_snapshot" {
  type    = bool
  default = false
}

# ECS
variable "ecs_task_cpu" {
  type    = number
  default = 1024 # Higher for production
}

variable "ecs_task_memory" {
  type    = number
  default = 2048 # Higher for production
}

variable "ecs_desired_count" {
  type    = number
  default = 3 # 3 tasks for production
}

variable "ecs_enable_autoscaling" {
  type    = bool
  default = true
}

variable "ecs_min_capacity" {
  type    = number
  default = 3
}

variable "ecs_max_capacity" {
  type    = number
  default = 20 # Higher for production
}

variable "ecs_enable_execute_command" {
  type    = bool
  default = false # Disabled for production security
}

# ALB
variable "alb_deletion_protection" {
  type    = bool
  default = true
}

# CloudFront
variable "cloudfront_price_class" {
  type    = string
  default = "PriceClass_All" # All edge locations for production
}

# WAF
variable "enable_waf" {
  type    = bool
  default = true
}

variable "waf_rate_limit" {
  type    = number
  default = 5000 # Higher limit for production
}

# Domain
variable "domain_name" {
  type = string
}

variable "create_hosted_zone" {
  type    = bool
  default = false
}

variable "hosted_zone_id" {
  type    = string
  default = null
}

# Cognito
variable "cognito_mfa_configuration" {
  type    = string
  default = "ON" # Required MFA for production
}

variable "cognito_deletion_protection" {
  type    = bool
  default = true
}

variable "cognito_advanced_security_mode" {
  type    = string
  default = "ENFORCED" # Full security for production
}

variable "cognito_callback_urls" {
  type    = list(string)
  default = []
}

variable "cognito_logout_urls" {
  type    = list(string)
  default = []
}

# GitHub Actions
variable "create_github_oidc_provider" {
  type    = bool
  default = false
}

variable "github_oidc_provider_arn" {
  type    = string
  default = null
}

variable "github_repository" {
  type = string
}
