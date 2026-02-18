# BlockScope

## English Version

### 📖 Project Introduction

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

### 📄 License

[Specify your license]

### 👤 Author

Yintc123

---

---

## 中文版本

### 📖 項目介紹

**BlockScope** 是一個用 Go 語言開發的區塊鏈日活躍地址追踪系統。該項目提供了一套完整的 REST API，用於查詢和統計在不同區塊鏈上的每日活躍地址數據。

### ✨ 主要功能

- 📊 **每日活躍地址統計**：追踪不同區塊鏈的每日活躍地址數量
- 🔗 **多鏈支援**：支持 Ethereum、Bitcoin、Solana 等多條區塊鏈
- 💾 **數據持久化**：使用 PostgreSQL 保存歷史數據
- 🏥 **健康監測**：提供健康檢查接口，監控服務狀態
- 🎯 **查詢接口**：根據日期和區塊鏈查詢統計數據

### 🏗️ 項目架構

```
BlockScope/
├── cmd/
│   └── server/              # 應用入口點
├── internal/
│   ├── config/              # 配置管理（支持本地和生產環境）
│   ├── db/                  # 數據庫連接管理
│   ├── domain/              # 業務模型
│   ├── repository/          # 數據訪問層（DAO）
│   ├── service/             # 業務邏輯層
│   ├── transport/
│   │   └── http/
│   │       ├── handler/     # HTTP 請求處理器
│   │       ├── request/     # 請求對象定義
│   │       └── routes/      # 路由定義
│   └── validator/           # 請求驗證
└── migrations/              # 數據庫遷移腳本
```

### 🛠️ 技術棧

- **語言**：Go 1.25.7
- **Web 框架**：Gin
- **數據庫**：PostgreSQL
- **ORM**：GORM
- **數據庫遷移**：golang-migrate
- **配置管理**：godotenv
- **驗證器**：validator/v10

### 🚀 快速開始

#### 環境要求

- Go 1.25.7 或更高版本
- PostgreSQL 12 或更高版本
- Docker（可選）

#### 安裝步驟

1. **克隆項目**
```bash
git clone https://github.com/Yintc123/BlockScope.git
cd BlockScope
```

2. **安裝依賴**
```bash
go mod download
```

3. **配置環境變數**

創建 `.env.local` 文件（開發環境）：
```env
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blockscope
```

或創建 `.env` 文件（生產環境）：
```env
APP_ENV=production
DB_HOST=your_host
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=blockscope
```

4. **初始化數據庫**

手動執行遷移腳本：
```bash
# 創建表
psql -U postgres -d blockscope -f migrations/000001_create_daily_active_address.up.sql
```

5. **運行服務**

```bash
# 開發環境（默認端口 8080）
APP_ENV=local go run cmd/server/main.go

# 生產環境（默認端口 80）
APP_ENV=production go run cmd/server/main.go
```

服務將在 `http://localhost:8080` 啟動。

### 📡 API 接口

#### 健康檢查接口

**請求**
```http
GET /healthcheck
```

**響應**
```json
{
  "status": "ok"
}
```

#### 獲取每日活躍地址統計

**請求**
```http
GET /stats/daily-active-address?date=2024-01-01&chain=eth
```

**參數**
- `date`：日期（格式：YYYY-MM-DD）【必需】
- `chain`：區塊鏈名稱（支持：eth, btc, sol）【必需】

**響應**
```json
{
  "id": 1,
  "date": "2024-01-01",
  "chain": "eth",
  "count": 1234567
}
```

**錯誤響應**
```json
{
  "error": "data not found"
}
```

### 📋 數據模型

#### DailyActiveAddress（每日活躍地址）

| 字段 | 類型 | 說明 |
|------|------|------|
| ID | uint | 主鍵 |
| Date | time.Time | 日期（與 Chain 組成唯一索引） |
| Chain | string | 區塊鏈名稱（與 Date 組成唯一索引） |
| Count | int64 | 活躍地址數量 |

### 🌍 支持的區塊鏈

- **eth** - Ethereum（以太坊）
- **btc** - Bitcoin（比特幣）
- **sol** - Solana（索拉納）

### 📝 配置說明

項目支持不同環境配置：

| 環境 | 配置文件 | 端口 | 用途 |
|------|---------|------|------|
| 開發 | `.env.local` | 8080 | 本地開發 |
| 生產 | `.env` | 80 | 生產部署 |

### 🔍 主要模塊說明

- **config**：負責加載和管理應用配置，支持環境隔離
- **db**：管理數據庫連接，初始化 GORM
- **repository**：數據訪問層，與數據庫交互
- **service**：業務邏輯層，處理應用的核心邏輯
- **handler**：HTTP 請求處理器，接收和返回 HTTP 響應
- **validator**：請求參數驗證，確保數據有效性

### 📄 許可證

[指定您的許可證]

### 👤 作者

Yintc123
