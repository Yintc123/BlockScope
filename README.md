# BlockScope

🌍 **Languages**: [English](README.md) | [中文](README.zh.md)

## 📖 Project Introduction

BlockScope tracks daily active addresses across blockchains and provides a REST API to query and aggregate daily active address counts by chain and date.

### ✨ Key Features

- Daily active address statistics
- Multi-chain support (eth, btc, sol)
- Persistent storage in PostgreSQL via GORM
- Health check endpoint
- Query API for analytics

### 🏗️ Project Structure

```
BlockScope/
├── cmd/
│   └── server/              # application entry point
├── internal/
│   ├── config/              # configuration
│   ├── db/                  # database connection
│   ├── domain/              # domain models
│   ├── repository/          # data access layer
│   ├── service/             # business logic
│   ├── transport/
│   │   └── http/
│   │       ├── handler/
│   │       ├── request/
│   │       └── routes/
│   └── validator/
└── migrations/              # database migration scripts
```

### 🛠️ Tech Stack

- Go 1.25.7
- Gin (HTTP)
- PostgreSQL 12+
- GORM (ORM)
- golang-migrate (migrations)
- godotenv (env loading)
- validator/v10 (request validation)
- testify (tests)

## 🚀 Quick Start

### Prerequisites

- Go 1.25.7 or newer
- PostgreSQL 12+
- Git

### Installation

1. Clone the repo

```bash
git clone https://github.com/Yintc123/BlockScope.git
cd BlockScope
```

2. Install dependencies

```bash
go mod download
```

3. Configure environment variables

For development (example):

```bash
cat > .env.local << EOF
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blockscope
EOF
```

For production:

```bash
cat > .env << EOF
APP_ENV=production
DB_HOST=your_host
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=blockscope
EOF
```

4. Initialize database

#### Option A — golang-migrate (recommended)

Install migrate (macOS/Homebrew example):

```bash
brew install golang-migrate
# or
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create database and run migrations:

```bash
psql -U postgres -c "CREATE DATABASE blockscope;"
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope?sslmode=disable" up
```

**Migration commands**:

```bash
# check version
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope?sslmode=disable" version

# rollback last
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope?sslmode=disable" down

# rollback all
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope?sslmode=disable" down -all
```

5. Run the service

```bash
APP_ENV=local go run cmd/server/main.go
```

The service listens on http://localhost:8080 by default.

## 🔄 Database Migration Management

Migration files live in `migrations/`. Use `golang-migrate` to apply or rollback migrations.

```bash
migrate -path migrations -database "postgres://user:password@host:5432/blockscope?sslmode=disable" up
```

Rollback (single step):

```bash
migrate -path migrations -database "postgres://user:password@host:5432/blockscope?sslmode=disable" down
```

File examples: `000001_create_daily_active_address.up.sql` / `.down.sql`.

Tip: validate migrations in a test database and keep migration files under version control.

## 📡 API

### Healthcheck

GET /health

Response 200:

```json
{
  "status": "ok",
  "checks": {
    "API server": true,
    "DB": "ok"
  }
}
```

Note: on failure `status` may be `"fail"`; the `checks.DB` field may contain an error string like `"fail: <error>"` and `API server` will be `false` when the server health check fails.

### Query daily active addresses

GET /stats/daily-active-address?date=YYYY-MM-DD&chain=eth

Response example:

```json
{
  "id": 1,
  "date": "2024-01-01",
  "chain": "eth",
  "count": 1234567
}
```

## 📋 Data Model

DailyActiveAddress

```sql
CREATE TABLE daily_active_addresses (
		id SERIAL PRIMARY KEY,
		date DATE NOT NULL,
		chain VARCHAR(50) NOT NULL,
		count BIGINT NOT NULL,
		UNIQUE(date, chain)
);

CREATE INDEX idx_date_chain ON daily_active_addresses (date, chain);
```

## 🌍 Supported Blockchains

| code | name | status |
|------|------|--------|
| eth  | Ethereum | ✅ |
| btc  | Bitcoin  | ✅ |
| sol  | Solana   | ✅ |

Supported chains are configured in `internal/config/config.go`.

## 📝 Configuration

The app supports multiple `.env` files:

| env | file | port |
|-----|------|------|
| development | .env.local | 8080 |
| production  | .env       | 80   |
| test        | .env.test  | N/A  |

Sample:

```env
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=blockscope
DB_SSLMODE=disable
```

## 🏗️ Architecture

Layered architecture and dependency injection are used; see `cmd/server/main.go` for bootstrap wiring.

## 🧪 Development & Testing

Run tests:

```bash
go test ./...
```

Tests require `.env.test`; `LoadConfig("test")` calls `util.GetProjectRoot()` to locate `.env.test` regardless of the test working directory. The helper walks upward from the caller; using `skip=0` is recommended for deterministic behavior when invoking package-level test helpers.

Test DB example:

```bash
psql -U postgres -c "CREATE DATABASE blockscope_test;"
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope_test?sslmode=disable" up
```

## 📚 Modules

- `internal/config/` — configuration
- `internal/db/` — DB connection
- `internal/domain/` — domain models
- `internal/repository/` — persistence
- `internal/service/` — business logic
- `internal/transport/http/` — handlers & routes
- `internal/util/` — helper functions

## 📌 Design Decisions & Best Practices

### Context lifecycle

Use request context from Gin (`c.Request.Context()`) when calling DB operations so cancellations/timeouts propagate.

### Test env loading

`LoadConfig("test")` calls `util.GetProjectRoot()` to locate `.env.test` regardless of the test working directory. The helper climbs the filesystem from the caller to find the target file; prefer `skip=0` for predictable results when used from shared test helpers.

## License

Add a license file or specify the project's license here.

## Author

Yintc123
