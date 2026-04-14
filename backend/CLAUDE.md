# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Instructions for CLAUDE

- If you see something missing in following sections, add it
- If you see that something have been deleted in project but not in this file, inform and propose a deletion
- If you see that something needs to be updated, update it

## Commands

All commands are run from the `backend/` directory unless noted.

### Backend (Go)

```bash
make code/build          # Build binary for current OS
make code/lint           # Run golangci-lint
make test/unit           # Run unit tests only
make test/integration    # Run integration tests only (requires services)
make test/all            # Run all tests with coverage
make code/generate       # Regenerate GraphQL resolvers + DAOs + go generate
make setup/services      # Start required Docker services (PostgreSQL, Keycloak, OPA, RabbitMQ)
```

Run a single test:

```bash
go test -tags=unit ./pkg/golang-graphql-example/... -run TestFunctionName
go test -tags=integration ./pkg/golang-graphql-example/... -run TestFunctionName
```

Run a single test inside a testify suite:

```bash
# Pattern: TestSuiteName/TestMethodName
go test -tags=unit ./pkg/golang-graphql-example/... -run TestMySuite/TestFunctionName
go test -tags=integration ./pkg/golang-graphql-example/... -run TestMySuite/TestFunctionName
```

## Technology Stack

### Language & Runtime

|          |                                                |
| -------- | ---------------------------------------------- |
| Language | Go 1.26                                        |
| Module   | `github.com/oxyno-zeta/golang-graphql-example` |

### Web & GraphQL

| Library                                  | Role                            |
| ---------------------------------------- | ------------------------------- |
| `github.com/gin-gonic/gin`               | HTTP framework                  |
| `github.com/99designs/gqlgen`            | GraphQL code generation         |
| `github.com/vektah/gqlparser/v2`         | GraphQL schema parser           |
| `github.com/graph-gophers/dataloader/v7` | Batching/DataLoader for GraphQL |
| `github.com/99designs/gqlgen-contrib`    | gqlgen extensions               |

### Database & ORM

| Library                       | Role                          |
| ----------------------------- | ----------------------------- |
| `gorm.io/gorm`                | ORM                           |
| `gorm.io/driver/postgres`     | PostgreSQL driver (primary)   |
| `gorm.io/driver/sqlite`       | SQLite driver (fallback/test) |
| `go-gormigrate/gormigrate/v2` | Schema migrations             |
| `gorm.io/plugin/dbresolver`   | Read replica support          |
| `DATA-DOG/go-sqlmock`         | SQL mock for unit tests       |

### Authentication & Authorization

| Library                        | Role                                  |
| ------------------------------ | ------------------------------------- |
| `github.com/coreos/go-oidc/v3` | OIDC/Keycloak token validation        |
| `golang.org/x/oauth2`          | OAuth 2.0                             |
| OPA (Open Policy Agent)        | Policy enforcement (external service) |

### Messaging

| Library                          | Role                       |
| -------------------------------- | -------------------------- |
| `github.com/rabbitmq/amqp091-go` | RabbitMQ AMQP 0.9.1 client |

### Observability

| Library                               | Role                               |
| ------------------------------------- | ---------------------------------- |
| `go.opentelemetry.io/otel`            | Tracing (OTLP/HTTP exporter)       |
| `otelgin` / `ravilushqa/otelgqlgen`   | Gin + GraphQL auto-instrumentation |
| B3, Jaeger, OT propagators            | Trace context propagation          |
| `github.com/prometheus/client_golang` | Prometheus metrics                 |
| `gorm.io/plugin/opentelemetry`        | GORM tracing                       |
| `gorm.io/plugin/prometheus`           | GORM metrics                       |

### Logging & Configuration

| Library                         | Role                         |
| ------------------------------- | ---------------------------- |
| `go.uber.org/zap`               | Structured logging           |
| `github.com/samber/slog-zap/v2` | slog → zap bridge            |
| `github.com/spf13/viper`        | YAML config with live reload |
| `github.com/spf13/cobra`        | CLI framework                |

### Utilities

| Library                    | Role                                  |
| -------------------------- | ------------------------------------- |
| `github.com/samber/lo`     | Functional helpers (map, filter, …)   |
| `emperror.dev/errors`      | Error handling/wrapping               |
| `github.com/gofrs/uuid`    | UUID generation                       |
| `cirello.io/pglock`        | PostgreSQL-backed distributed locking |
| `go.uber.org/automaxprocs` | Auto GOMAXPROCS from cgroup limits    |

### Testing

| Library                       | Role                      |
| ----------------------------- | ------------------------- |
| `github.com/stretchr/testify` | Assertions, mocks, suites |
| `go.uber.org/mock` (mockgen)  | Mock generation           |
| `go.uber.org/goleak`          | Goroutine leak detection  |
| `DATA-DOG/go-sqlmock`         | SQL layer mocking         |

### Code Generation

| Tool                            | Role                                   |
| ------------------------------- | -------------------------------------- |
| `gqlgen`                        | GraphQL resolvers + models from schema |
| `mockgen`                       | Interface mocks for tests              |
| `tools/generator/daogen/`       | DAO boilerplate generation             |
| `tools/generator/graphqlgen/`   | Extra GraphQL generation               |
| `tools/generator/modeltagsgen/` | GORM model tag generation              |

### Linting & Formatting

- **golangci-lint** — 50+ linters enabled (`gocritic`, `staticcheck`, `gosec`, `revive`, `errchkjson`, `modernize`, …)
- **gofmt** with `interface{}` → `any` rewrite
- **gofumpt** (extra strictness), **goimports**, **gci** (import section ordering), **golines** (max 160 chars)

## Architecture

### Multi-target Application

The backend supports multiple launch targets defined in `cmd/golang-graphql-example/main.go`:

- `server` — HTTP/GraphQL server
- `migrate-db` — standalone migration runner
- `all` — runs all applicable targets

Those are just examples. Other targets can be found in `cmd/golang-graphql-example/main.go` file.

### Backend Package Layout

All application code lives under `pkg/golang-graphql-example/`:

| Package                 | Responsibility                                           |
| ----------------------- | -------------------------------------------------------- |
| `config/`               | YAML config with live reload via Viper                   |
| `authx/authentication/` | Keycloak/OIDC token validation                           |
| `authx/authorization/`  | OPA (Open Policy Agent) policy enforcement               |
| `business/`             | Domain services (e.g., `todos/`) — add new entities here |
| `business/migration/`   | gormigrate database schema migrations                    |
| `common/`               | Shared utilities (correlation IDs, error helpers)        |
| `database/`             | GORM setup; supports PostgreSQL (primary) and SQLite     |
| `email/`                | Email sending service                                    |
| `log/`                  | Logger setup, GORM/lock distributor log adapters         |
| `messagebus/amqp/`      | RabbitMQ AMQP client                                     |
| `lockdistributor/`      | Distributed locking backed by PostgreSQL                 |
| `metrics/`              | Prometheus metrics                                       |
| `server/graphql/`       | gqlgen resolvers, generated models, custom scalars       |
| `signalhandler/`        | OS signal handling (graceful shutdown)                   |
| `tracing/`              | OpenTelemetry setup (exports to Jaeger/Tempo)            |
| `version/`              | Build version info                                       |

### GraphQL Code Generation

Schema files live in `graphql/*.graphql`. After modifying the schema:

```bash
make code/graphql/generate   # regenerates resolvers via gqlgen
```

Config: `gqlgen.yml` — models autobind from `pkg/.../server/graphql/model`, resolvers follow the schema layout.

### Service Initialization Order

In `main.go`, services initialize in a strict order:

1. Logger → Config → Version
2. Metrics → Tracing → Database → Email → Lock Distributor → Signal Handler → AMQP → Auth
3. Business units (todos, etc.)

This ordering matters for dependency injection — later services can depend on earlier ones.

### Testing Strategy

Tests use build tags to separate concerns:

- `//go:build unit` — fast, no external services
- `//go:build integration` — requires the full service stack started via `make setup/services`

### Adding a New Business Entity

1. Create a package under `business/<entity>/`
2. Define GORM model + gormigrate migration in `business/migration/`
3. Add GraphQL types to `graphql/<entity>.graphql`
4. Run `make code/generate`
5. Implement the generated resolver interface in `server/graphql/`

### Configuration

Config files live in `conf/` (YAML). The config manager supports live reload on file changes. Use the `config/` package to access typed config structs — do not use Viper directly in business logic.

### Commit Convention

Follow Angular commit convention: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`, etc. This is enforced by pre-commit hooks via Commitizen.
