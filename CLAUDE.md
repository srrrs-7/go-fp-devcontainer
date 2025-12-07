# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go monorepo with functional programming patterns. Uses Go 1.25.4 with workspaces, Docker Compose for services, and PostgreSQL 18.

## Repository Structure

```
apps/
  api/          # Backend API (go-chi router, port 8080)
  pkgs/         # Shared packages (db, logger, env, types)
  web/          # Frontend (WASM experimentation, port 3000)
  iac/          # Terraform infrastructure (AWS)
```

Go workspace: `apps/go.work` manages `api`, `pkgs`, and `web` modules.

## Development Commands

```bash
# Local services (from repo root)
docker compose up -d          # Start all services (api, web, db)
docker compose up -d db       # Start only PostgreSQL

# Tests (all modules)
make test

# Run single test
cd apps/api && go test -run TestListHandler ./src/routes/tasks/

# Code quality
make fmt      # Format Go code
make vet      # Static analysis
make tidy     # go mod tidy
make tf-fmt   # Format Terraform files

# Database migrations (Atlas)
make atlas-diff              # Generate migration from schema changes
make atlas-apply             # Apply pending migrations
make atlas-new NAME=<name>   # Create new migration file
make atlas-status            # Show migration status
# Use ATLAS_ENV=docker for Docker environment: make atlas-apply ATLAS_ENV=docker

# sqlc code generation
make sqlc-gen      # Generate Go code from SQL queries
make sqlc-compile  # Validate SQL queries

# WASM build
make wasm          # Build WebAssembly binary
make wasm-clean    # Remove WASM artifacts
```

## Architecture

### Functional Error Handling with Result Type

The codebase uses a `Result[T, E]` monad (`apps/pkgs/types/result.go`) for functional error handling instead of Go's traditional `(value, error)` pattern:

```go
// Pipeline example from handlers
res := types.Pipe2(
    newListRequest(r).validate(),
    func(req listRequest) types.Result[[]model.Task, model.AppError] {
        return task_repository.FindAllTasks()
    },
    func(tasks []model.Task) listResponse { ... },
)

res.Match(
    func(resp listResponse) { response.OK(w, resp) },
    func(e model.AppError) { response.HandleAppError(w, e) },
)
```

Key functions: `Ok()`, `Err()`, `Map()`, `FlatMap()`, `Pipe2-5()`, `Match()`, `Combine()`

### Domain Model Pattern

- Value objects with type safety: `TaskID`, `TaskTitle`, `TaskDescription`, `TaskCompleted`
- Domain errors implement `AppError` interface with `ErrorName()` and `DomainName()`
- Error types: `ValidationError`, `NotFoundError`, `DatabaseError`, etc.

### API Layer Structure (apps/api/src)

```
cmd/main.go              # Entry point, graceful shutdown
routes/
  routes.go              # Chi router setup, /api/v1 prefix
  response/response.go   # JSON response helpers, error mapping
  tasks/                 # Handler per endpoint (list.go, post.go, get.go, put.go)
domain/model/            # Domain types and errors
infra/rds/               # Repository implementations
```

Route pattern: `/api/v1/tasks`, `/api/v1/tasks/{id}`

### Web Frontend (apps/web)

Go + WebAssembly frontend using `syscall/js` for DOM manipulation:
- `main.go` - WASM entry point, exposes `fetchTasks()` and `addTask()` to JavaScript
- `static/index.html` - HTML/CSS template, loads `wasm_exec.js` and `main.wasm`
- Dockerfile uses multi-stage build: Go WASM compilation → nginx for serving

Build: `GOOS=js GOARCH=wasm go build -o main.wasm .`

### Database Layer (apps/pkgs/db)

- **Atlas**: Schema-first migrations in `migrations/`
- **sqlc**: Type-safe query generation from `queries/*.sql` → `db/*.go`

Configuration files:
- `atlas.hcl` - Migration environments (local, docker, ci)
- `sqlc.yaml` - Code generation config

## Testing Patterns

Tests use table-driven pattern with `args`/`expected` structs:

```go
tests := []struct {
    testName string
    args     args
    expected expected
}{...}

for _, tt := range tests {
    t.Run(tt.testName, func(t *testing.T) { ... })
}
```

HTTP handlers tested with `httptest.NewRequest` and `httptest.NewRecorder`.

## Infrastructure (apps/iac)

Terraform modules for AWS deployment:

```
apps/iac/
  environments/       # Environment configs (dev, stg, prd)
  modules/           # Reusable Terraform modules
    vpc/             # Multi-AZ VPC with public/private/database subnets
    ecs/             # ECS Fargate service
    ecr/             # Container registry
    alb/             # Application Load Balancer
    aurora/          # Aurora PostgreSQL Serverless v2
    cognito/         # Authentication (MFA-enabled)
    cloudfront/      # CDN (S3 + ALB origins)
    s3/              # Static assets
    waf/             # Web Application Firewall
    acm/             # SSL/TLS certificates
    route53/         # DNS records
    iam/             # GitHub Actions OIDC, ECS task roles
    security-groups/ # ALB, ECS, Aurora security groups
```

Naming convention: `${project}-${environment}-${resource}`

Terraform setup:
```bash
cd apps/iac/environments/dev
cp backend.hcl.example backend.hcl     # Configure S3 state backend
cp terraform.tfvars.example terraform.tfvars  # Set environment variables
terraform init -backend-config=backend.hcl
terraform plan && terraform apply
```

## CI/CD Pipeline

**CI**: GitHub Actions runs in devcontainer on push/PR to main: `make vet && make test`

**CD**: Triggered by push to `main` (→ dev) or manual workflow dispatch (→ dev/stg/prd)
- Flow: Database Migration → Build & Push to ECR → Deploy to ECS
- Environments configured in GitHub Settings → Environments with AWS OIDC credentials
