// Atlas configuration for database migrations
// See: https://atlasgo.io/atlas-schema/hcl

// Define environment variables from .devcontainer/compose.override.yaml
variable "db_host" {
  type    = string
  default = getenv("DB_HOST")
}

variable "db_port" {
  type    = string
  default = getenv("DB_PORT")
}

variable "db_name" {
  type    = string
  default = getenv("DB_DBNAME")
}

variable "db_user" {
  type    = string
  default = getenv("DB_USERNAME")
}

variable "db_password" {
  type    = string
  default = getenv("DB_PASSWORD")
}

variable "db_uri" {
  type    = string
  default = getenv("DB_URI")
}

variable "dev_db_uri" {
  type    = string
  default = getenv("DEV_DB_URI")
}

// Construct database URL from environment variables
locals {
  db_url = var.db_uri != "" ? var.db_uri : "postgres://${var.db_user}:${var.db_password}@${var.db_host}:${var.db_port}/${var.db_name}?sslmode=disable"
  dev_db_url = var.dev_db_uri != "" ? var.dev_db_uri : "postgres://${var.db_user}:${var.db_password}@${var.db_host}:${var.db_port}/${var.db_name}_dev?sslmode=disable"
}

// Environment configuration
env "local" {
  // Source schema files - Atlas will use these as the desired state
  src = "file://schema"

  // Target database URL - where migrations will be applied
  url = local.db_url

  // Dev database URL - used for diffing and validation
  // This should be a separate database that Atlas can freely manipulate
  dev = local.dev_db_url

  // Migration directory configuration
  migration {
    // Directory where migration files are stored
    dir = "file://migrations"
  }
}

// Docker Compose environment (same as local, for compatibility)
env "docker" {
  src = "file://schema"
  url = local.db_url
  dev = local.dev_db_url

  migration {
    dir = "file://migrations"
  }
}

// CI/CD environment
env "ci" {
  src = "file://schema"
  url = local.db_url

  migration {
    dir = "file://migrations"

    // Baseline version (if needed for existing databases)
    // baseline = "20240101000000"
  }

  // Lint configuration for CI
  lint {
    // Review policy
    review = ERROR

    // Detect destructive changes
    destructive {
      error = true
    }

    // Detect data-dependent changes
    data_depend {
      error = true
    }
  }
}
