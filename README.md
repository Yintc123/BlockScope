# BlockScope

🌍 **Languages**: [English](README.md) | [中文](README.zh.md)

## 📖 Project Introduction

**BlockScope** is a blockchain daily active address tracking system developed in Go. This project provides a complete set of REST APIs for querying and analyzing daily active addresses data across different blockchains.

### ✨ Key Features

- 📊 **Daily Active Address Statistics**: Track the number of daily active addresses on different blockchains
- 🔗 **Multi-chain Support**: Support for Ethereum, Bitcoin, Solana and other blockchains
- 💾 **Data Persistence**: Store historical data using PostgreSQL
- 🏥 **Health Monitoring**: Provides health check interface to monitor service status
- 🎯 **Query Interface**: Query statistical data by date and blockchain

### 🏗️ Project Architecture

```
BlockScope/
├── cmd/
│   └── server/              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── db/                  # Database connection management
│   ├── domain/              # Business models
│   ├── repository/          # Data access layer (DAO)
│   ├── service/             # Business logic layer
│   ├── transport/
│   │   └── http/
│   │       ├── handler/     # HTTP request handlers
│   │       ├── request/     # Request object definitions
│   │       └── routes/      # Route definitions
│   └── validator/           # Request validation
└── migrations/              # Database migration scripts
```

### 🛠️ Technology Stack

- **Language**: Go 1.25.7
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Database Migration**: golang-migrate
- **Configuration**: godotenv
- **Validator**: validator/v10

### 🚀 Quick Start

#### Requirements

- Go 1.25.7 or higher
- PostgreSQL 12 or higher
- Docker (optional)

#### Installation Steps

1. **Clone the Repository**
```bash
git clone https://github.com/Yintc123/BlockScope.git
cd BlockScope
```

2. **Install Dependencies**
```bash
go mod download
```

3. **Configure Environment Variables**

Create `.env.local` file (Development):
```env
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blockscope
```

Or create `.env` file (Production):
```env
APP_ENV=production
DB_HOST=your_host
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=blockscope
```

4. **Initialize Database**

Execute the migration script manually:
```bash
# Create table
psql -U postgres -d blockscope -f migrations/000001_create_daily_active_address.up.sql
```

5. **Run the Service**

```bash
# Development (default port 8080)
APP_ENV=local go run cmd/server/main.go

# Production (default port 80)
APP_ENV=production go run cmd/server/main.go
```

The service will start on `http://localhost:8080`.

### 📡 API Endpoints

#### Health Check Endpoint

**Request**
```http
GET /healthcheck
```

**Response**
```json
{
  "status": "ok"
}
```

#### Get Daily Active Address Statistics

**Request**
```http
GET /stats/daily-active-address?date=2024-01-01&chain=eth
```

**Parameters**
- `date`: Date in format YYYY-MM-DD [Required]
- `chain`: Blockchain name (supported: eth, btc, sol) [Required]

**Response**
```json
{
  "id": 1,
  "date": "2024-01-01",
  "chain": "eth",
  "count": 1234567
}
```

**Error Response**
```json
{
  "error": "data not found"
}
```

### 📋 Data Models

#### DailyActiveAddress

| Field | Type | Description |
|-------|------|-------------|
| ID | uint | Primary key |
| Date | time.Time | Date (forms unique index with Chain) |
| Chain | string | Blockchain name (forms unique index with Date) |
| Count | int64 | Number of active addresses |

### 🌍 Supported Blockchains

- **eth** - Ethereum
- **btc** - Bitcoin
- **sol** - Solana

### 📝 Configuration Details

The project supports multiple environment configurations:

| Environment | Config File | Port | Purpose |
|-------------|------------|------|---------|
| Development | `.env.local` | 8080 | Local development |
| Production | `.env` | 80 | Production deployment |

### 🔍 Module Descriptions

- **config**: Loads and manages application configuration with environment isolation
- **db**: Manages database connection and initializes GORM
- **repository**: Data access layer that interacts with the database
- **service**: Business logic layer that handles core application logic
- **handler**: HTTP request handlers that receive and return HTTP responses
- **validator**: Request parameter validation to ensure data validity

### � Design Decisions and Best Practices

#### Context Lifecycle Management

**Related Code Location**
- Path: [internal/transport/http/handler/healthcheck_handler.go](internal/transport/http/handler/healthcheck_handler.go)
- Method: `Check()`
- Dependency: `CheckDB(ctx context.Context)` in [internal/service/healthcheck_service.go](internal/service/healthcheck_service.go)

**Background**

The `CheckDB()` method requires a valid context object to control the database health check operation through `sqlDB.PingContext(ctx)`.

**Frequently Asked Questions**

**Q1: Does using `context.Background()` waste resources?**

No. `context.Background()` creates a lightweight root context with the following characteristics:

- Consumes minimal memory resources
- Does not spawn additional goroutines
- Does not hold external resources (e.g., database connections, file handles)
- Is a non-cancellable, timeout-free root context

**Q2: Does repeatedly creating context objects when calling APIs accumulate over time?**

No. Go's garbage collection (GC) mechanism automatically reclaims unused context objects, so accumulation is not a concern.

**Recommended Optimization**

In production applications, you should use the request context provided by the Gin framework instead of `context.Background()`:

```go
func (handler *HealthcheckHandler) Check(c *gin.Context) {
	ctx := c.Request.Context()  // Use HTTP request's context
	
	dbErr := handler.service.CheckDB(ctx)
	// ...
}
```

**Optimization Benefits**

| Benefit | Description |
|---------|-------------|
| **Resource Efficiency** | Avoids creating a new context object for each request |
| **Lifecycle Synchronization** | When HTTP request is cancelled or times out, downstream database operations also stop automatically |
| **Trace Consistency** | Retains HTTP request trace context information for distributed tracing and monitoring |
#### Environment Variable Loading Strategy for Tests

**Related Code Location**
- Path: [internal/config/config.go](internal/config/config.go)
- Method: `LoadConfig(env string)`
- Utility Function: `GetProjectRoot(skip int, targetFileName string)` in [internal/util/path.go](internal/util/path.go)

**Background**

When executing `go test`, you may find that environment variables fail to load (returning nil/empty values). This occurs because the working directory during test execution is the directory containing the test file, not the project root. If you attempt to load `.env.test` from the root directory using `LoadConfig("test")` directly, the configuration module will be unable to locate the environment file due to the working directory discrepancy.

**Frequently Asked Questions**

**Q1: Why not use relative paths directly to load .env.test in tests?**

Because `go test` executes with the working directory set to the test file's directory, not the project root. For example, when running tests in the `repository` directory, relative paths will be resolved from `internal/repository`, causing the file lookup to fail for `.env.test` in the project root.

**Recommended Approach**

Create a `util` module that uses `runtime.Caller()` to dynamically obtain the caller's file path, then traverse upward through the directory tree until the target file is found:

**skip Parameter Explanation**

The `skip` parameter specifies which level of the call stack to retrieve the file path from:

| skip Value | Description | Returned File |
|------------|-------------|---------------|
| 0 | Current function's file | `path.go` (the `GetProjectRoot` function itself) |
| 1 | Direct caller's file | The file calling `GetProjectRoot` (e.g., `config.go`) |
| 2 | Caller's caller's file | The file calling `config.go` |

**Practical Example:**
```
config.go (skip=1) calls → GetProjectRoot(1) → returns config.go's path
                                ↓
                         path.go (skip=0) file itself
```

**Why is skip=0 recommended?**

When calling from the configuration module, you should pass `skip=0` for the following reasons:

1. **Call Stack Determinism**: Different callers might invoke `GetProjectRoot()` (in tests, other modules, etc.). Using `skip=1` would require knowing the exact stack level and is error-prone. Using `skip=0` calculates directly from the `GetProjectRoot()` function itself, independent of call chain depth.

2. **Consistent Traversal Starting Point**: With `skip=0`, `GetProjectRoot()` starts searching upward from `path.go`'s directory. Regardless of the caller's location, it will eventually traverse to the project root, ensuring files like `.env.test` are found.

3. **API Universality**: When calling `GetProjectRoot()` from multiple modules (`config.go`, `db.go`, etc.), using `skip=0` requires no adjustment—the same value works everywhere, reducing error risk.

```go
// internal/util/path.go
func GetProjectRoot(skip int, targetFileName string) string {
	// Dynamically locate file path using runtime.Caller
	_, b, _, _ := runtime.Caller(skip)
	dir := filepath.Dir(b)

	// Traverse upward until finding the target file
	for {
		if _, err := os.Stat(filepath.Join(dir, targetFileName)); err == nil {
			return dir
		}
		
		parent := filepath.Dir(dir)
		if parent == dir {
			break // Reached system root
		}
		dir = parent
	}

	// Return current working directory if target not found
	cwd, _ := os.Getwd()
	return cwd
}
```

Usage in the configuration module:

```go
// internal/config/config.go
case "test":
	var rootPath string = util.GetProjectRoot(0, ".env.test")
	err := godotenv.Load(filepath.Join(rootPath, ".env.test"))
	if err != nil {
		return nil, fmt.Errorf("could not load .env.test from %s: %w", rootPath, err)
	}
```

**Optimization Benefits**

| Benefit | Description |
|---------|-------------|
| **Working Directory Independent** | Does not depend on runtime working directory; tests can execute from any location |
| **Dynamic Path Resolution** | Automatically locates the correct project root without manual path configuration |
| **Scalability** | The same mechanism can be applied to other environment files (e.g., `.env.staging`) |
### �📄 License

[Specify your license]

### 👤 Author

Yintc123

