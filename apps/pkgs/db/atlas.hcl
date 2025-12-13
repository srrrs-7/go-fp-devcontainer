// Atlas configuration for database migrations
// See: https://atlasgo.io/atlas-schema/hcl

// Define environment variables from .devcontainer/compose.override.yaml
variable "db_uri" {
  type    = string
  default = getenv("DB_MIGRATE_URI")
}

// Construct database URL from environment variables
locals {
  url = var.db_uri
}

// Environment configuration
env "local" {
  // Source schema files - Atlas will use these as the desired state
  src = "file://schema"

  // Target database URL - where migrations will be applied
  url = local.url

  // Migration directory configuration
  migration {
    // Directory where migration files are stored
    dir = "file://migrations"
  }
}

// Docker Compose environment (same as local, for compatibility)
env "docker" {
  src = "file://schema"
  url = local.url

  migration {
    dir = "file://migrations"
  }
}

// CI/CD environment
env "ci" {
  src = "file://schema"
  url = local.url

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
