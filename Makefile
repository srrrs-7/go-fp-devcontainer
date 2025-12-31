PHONY: help
help:
	cat ./Makefile

#############
# Variables #
#############
API_MOD = ./apps/api
PKGS_MOD = ./apps/pkgs
WEB_MOD = ./apps/web
DB_MOD = ./apps/pkgs/db
IAC_DIR = ./apps/iac
MODS = $(API_MOD) $(PKGS_MOD) $(WEB_MOD)
ATLAS_ENV ?= local
ATLAS_DIR = $(DB_MOD)

#################
# Dev Container #
#################
.PHONY: dev-cp run-api run-web run-migrate run-all

dev-cp:
	cp .devcontainer/compose.override.yaml.example .devcontainer/compose.override.yaml

# Run API server (port 8080)
run-api:
	cd ${API_MOD}/src && go run ./cmd

# Run Web server (port 3000)
# Requires: make templ-gen (run once after templ file changes)
run-web:
	cd ${WEB_MOD}/src && go run ./cmd

# Run database migrations (Atlas)
run-migrate:
	@echo "Running database migrations..."
	cd ${DB_MOD} && atlas migrate apply --env local

# Run all services (migrate -> api -> web in background)
run-all: run-migrate
	@echo "Starting API server..."
	cd ${API_MOD}/src && go run ./cmd &
	@echo "Starting Web server..."
	cd ${WEB_MOD}/src && go run ./cmd

#############
# Container #
#############
.PHONY: cp fmt vet test tidy graph env

cp:
	cp compose.override.yaml.example compose.override.yaml

test:
	for mod in $(MODS); do \
		(cd $$mod && go test ./...); \
	done

tidy:
	for mod in $(MODS); do \
		(cd $$mod && go mod tidy); \
	done

fmt:
	for mod in $(MODS); do \
		(cd $$mod && go fmt ./...); \
	done	

vet:
	for mod in $(MODS); do \
		(cd $$mod && go vet ./...); \
	done

graph:
	for mod in $(MODS); do \
		(cd $$mod && go mod graph); \
	done

env:
	for mod in $(MODS); do \
		(cd $$mod && go env); \
	done

#########
# templ #
#########
.PHONY: templ-gen templ-watch templ-fmt

templ-gen:
	cd ${WEB_MOD}/src && templ generate

templ-watch:
	cd ${WEB_MOD}/src && templ generate --watch

templ-fmt:
	cd ${WEB_MOD}/src && templ fmt .

########
# wasm #
########
# NOTE: WASM targets are deprecated. Use templ for web frontend.
.PHONY: wasm wasm-clean

wasm:
	@echo "DEPRECATED: Use 'make templ-gen' instead"
	@echo "Building WebAssembly..."
	cd ${WEB_MOD} && GOOS=js GOARCH=wasm go build -o main.wasm .
	@echo "Done: ${WEB_MOD}/main.wasm"

wasm-clean:
	rm -f ${WEB_MOD}/main.wasm

########
# sqlc #
########
.PHONY: sqlc-gen sqlc-gen sqlc-compile sqlc-verify sqlc-help

sqlc-gen:
	cd ${DB_MOD} && sqlc generate

sqlc-compile:
	cd ${DB_MOD} && sqlc compile

sqlc-verify:
	cd ${DB_MOD} && sqlc verify

###################
# atlas migration #
###################
.PHONY: atlas-new atlas-diff atlas-apply atlas-status atlas-hash atlas-lint atlas-inspect atlas-clean atlas-validate

# Create a new migration file with a given name
# Usage: make atlas-new NAME=create_users_table
atlas-new:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make atlas-new NAME=create_users_table"; \
		exit 1; \
	fi
	cd ${ATLAS_DIR} && atlas migrate new --env ${ATLAS_ENV} $(NAME)

# Generate a migration by comparing schema files with the current database state
# This will create a new migration file based on the diff
atlas-diff:
	@echo "Generating migration from schema diff..."
	cd ${ATLAS_DIR} && atlas migrate diff --env ${ATLAS_ENV}

# Apply pending migrations to the database
# Usage: make atlas-apply [ENV=docker]
atlas-apply:
	@echo "Applying migrations to database (env: ${ATLAS_ENV})..."
	cd ${ATLAS_DIR} && atlas migrate apply --env ${ATLAS_ENV}

# Apply migrations with auto-approval (use with caution)
atlas-apply-auto:
	@echo "Applying migrations with auto-approval (env: ${ATLAS_ENV})..."
	cd ${ATLAS_DIR} && atlas migrate apply --env ${ATLAS_ENV} --auto-approve

# Show migration status (applied vs pending)
atlas-status:
	@echo "Migration status (env: ${ATLAS_ENV}):"
	cd ${ATLAS_DIR} && atlas migrate status --env ${ATLAS_ENV}

# Rehash migration directory (update atlas.sum)
# Run this after manually creating or editing migration files
atlas-hash:
	@echo "Rehashing migration directory..."
	cd ${ATLAS_DIR} && atlas migrate hash

# Lint migration files for issues
atlas-lint:
	@echo "Linting migrations..."
	cd ${ATLAS_DIR} && atlas migrate lint --env ${ATLAS_ENV}

# Validate migration directory integrity
atlas-validate:
	@echo "Validating migration directory..."
	cd ${ATLAS_DIR} && atlas migrate validate

# Inspect current database schema
atlas-inspect:
	@echo "Inspecting database schema (env: ${ATLAS_ENV})..."
	cd ${ATLAS_DIR} && atlas schema inspect --env ${ATLAS_ENV}

# Apply schema directly (declarative approach - bypasses migrations)
# WARNING: This can be destructive. Use for development only.
atlas-schema-apply:
	@echo "Applying schema directly (env: ${ATLAS_ENV})..."
	cd ${ATLAS_DIR} && atlas schema apply --env ${ATLAS_ENV} --auto-approve

# Clean/reset dev database (for development)
# This will drop and recreate the dev database
atlas-clean:
	@echo "Cleaning dev database..."
	@echo "WARNING: This will drop all data in the dev database."
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		cd ${ATLAS_DIR} && atlas schema clean --env ${ATLAS_ENV} --auto-approve; \
	fi

#############
# Terraform #
#############
.PHONY: tf-fmt

tf-fmt:
	cd ${IAC_DIR} && terraform fmt -recursive
