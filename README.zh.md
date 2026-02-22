# BlockScope

🌍 **Languages**: [English](README.md) | [中文](README.zh.md)

## 📖 項目介紹

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
- **Web 框架**：[Gin](https://github.com/gin-gonic/gin) - 輕量級的 HTTP 框架
- **數據庫**：PostgreSQL 12+
- **ORM**：[GORM](https://gorm.io/) - 物件關係對應庫
- **數據庫遷移**：[golang-migrate](https://github.com/golang-migrate/migrate) - 數據庫版本控制
- **配置管理**：[godotenv](https://github.com/joho/godotenv) - 環境變數管理
- **驗證**：[validator/v10](https://github.com/go-playground/validator) - 請求參數驗證
- **測試**：[testify](https://github.com/stretchr/testify) - 測試斷言庫

### 🚀 快速開始

#### 環境要求

- Go 1.25.7 或更高版本
- PostgreSQL 12 或更高版本
- Git

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

用於 **開發環境**（本地測試）：
```bash
# 建立 .env.local 檔案
cat > .env.local << EOF
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=blockscope
EOF
```

用於 **生產環境**（部署）：
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

4. **初始化數據庫**

#### 使用 golang-migrate

使用 `golang-migrate` 管理資料庫遷移，遷移檔位於 `migrations/`。重要指令：

執行遷移：
```bash
migrate -path migrations -database "postgres://user:password@host:5432/blockscope?sslmode=disable" up
```

回滾（單步）：
```bash
migrate -path migrations -database "postgres://user:password@host:5432/blockscope?sslmode=disable" down
```

檔名範例：`000001_create_daily_active_address.up.sql` / `000001_create_daily_active_address.down.sql`。

簡短建議：在測試資料庫先驗證遷移，並將遷移檔納入版本控制。

5. **運行服務**

**開發模式**（預設端口 8080）：
```bash
APP_ENV=local go run cmd/server/main.go
```

**生產模式**（預設端口 80 - 需要管理員權限）：
```bash
APP_ENV=production go run cmd/server/main.go
```

應用將啟動並監聽 `http://localhost:8080`（或您設定的端口）。

### 📡 API 接口

#### 健康檢查

驗證服務和資料庫連接：

**請求**
```http
GET /health
```

**響應** (200 OK)
```json
{
  "status": "ok",
  "checks": {
    "API server": true,
    "DB": "ok"
  }
}
```

備註：若檢查失敗，`status` 可能為 `"fail"`；`checks.DB` 可包含錯誤字串，如 `"fail: <error>"`，且 `API server` 將為 `false`。

#### 查詢每日活躍地址統計

取得特定區塊鏈和日期的每日活躍地址數量：

**請求**
```http
GET /stats/daily-active-address?date=2024-01-01&chain=eth
```

**查詢參數**

| 參數 | 型態 | 必需 | 說明 |
|------|------|------|------|
| `date` | string | 是 | 日期格式 `YYYY-MM-DD` |
| `chain` | string | 是 | 區塊鏈識別碼（eth、btc、sol） |

**響應** (200 OK)
```json
{
  "id": 1,
  "date": "2024-01-01",
  "chain": "eth",
  "count": 1234567
}
```

**錯誤響應** (404 Not Found)
```json
{
  "error": "data not found"
}
```

### 📋 數據模型

#### DailyActiveAddress

| 字段 | 型態 | 約束條件 | 說明 |
|------|------|--------|------|
| ID | `uint` | 主鍵 | 唯一識別碼 |
| Date | `time.Time` | UNIQUE（与Chain）、已索引 | 交易日期，資料庫使用 `DATE` 類型 |
| Chain | `string` | UNIQUE（与Date）、已索引 | 區塊鏈名稱 |
| Count | `int64` | Not Null | 活躍地址數量 |

**數據庫結構**
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

### 🌍 支持的區塊鏈

目前已配置和支持以下區塊鏈：

| 代碼 | 名稱 | 狀態 |
|------|------|------|
| `eth` | Ethereum（以太坊） | ✅ 支持 |
| `btc` | Bitcoin（比特幣） | ✅ 支持 |
| `sol` | Solana（索拉納） | ✅ 支持 |

**說明**：支持的區塊鏈清單定義在 `internal/config/config.go`，可通過修改配置輕鬆擴展。

### 📝 配置說明

應用通過 `.env` 文件支持多環境配置：

| 環境 | 檔案 | 端口 | 模式 | 用途 |
|------|------|------|------|------|
| 開發 | `.env.local` | 8080 | 調試友好 | 本地開發和測試 |
| 生產 | `.env` | 80 | 優化 | 生產部署 |
| 測試 | `.env.test` | N/A | 測試專用 | 單元和集成測試 |

**配置優先級**：
1. 環境變數（來自 `.env` 檔案）
2. 預設值（程式碼中硬編碼）

**配置示例**：
```env
# 應用設定
APP_ENV=local                   # local、production 或 test

# 數據庫連接
DB_HOST=localhost               # 資料庫伺服器地址
DB_PORT=5432                    # PostgreSQL 預設端口
DB_USER=postgres                # 資料庫使用者
DB_PASSWORD=your_secure_password # 使用者密碼（如無需可留空）
DB_NAME=blockscope              # 資料庫名稱
DB_SSLMODE=disable              # SSL 模式（disable、require 等）
```

### 🏗️ 項目架構

#### 分層架構

項目採用清潔的分層架構模式，實現關注點分離：

```
┌─────────────────────────────────────────┐
│  傳輸層（HTTP）                          │
│  ├── handler/      (請求處理器)         │
│  ├── request/      (DTO 定義)           │
│  └── routes/       (路由定義)           │
├─────────────────────────────────────────┤
│  業務邏輯層（Service）                   │
│  ├── 業務邏輯和算法                      │
│  └── 業務規則驗證                        │
├─────────────────────────────────────────┤
│  數據訪問層（Repository）                 │
│  ├── 數據庫查詢（通過 GORM）            │
│  └── 數據轉換                           │
├─────────────────────────────────────────┤
│  領域層（Domain）                        │
│  └── DailyActiveAddress 結構體           │
├─────────────────────────────────────────┤
│  基礎設施層                              │
│  ├── config/       (配置管理)            │
│  ├── db/           (數據庫連接)          │
│  ├── util/         (工具函式)            │
│  └── validator/    (請求驗證)            │
└─────────────────────────────────────────┘
```

#### 依賴注入

應用在 `bootstrap()` 函式（[cmd/server/main.go](cmd/server/main.go)）中使用手動依賴注入來組裝所有組件：

```go
// bootstrap 組裝所有依賴並返回已配置的 Gin 路由器
func bootstrap(env string) (*gin.Engine, string, error) {
	// 1. 加載配置
	cfg, err := config.LoadConfig(env)
	
	// 2. 初始化數據庫
	dbConn, err := db.NewDB(cfg.DB)
	
	// 3. 組裝依賴（DI）
	repo := repository.NewDailyActiveAddressRepository(dbConn)
	svc := service.NewDailyActiveAddressService(repo)
	handler := handler.NewStatsHandler(svc)
	
	// 4. 註冊路由並返回
	router := setupRoutes(handler)
	return router, port, nil
}
```

**優勢**：
- ✅ 顯式依賴管理
- ✅ 易於測試（模擬依賴）
- ✅ 清晰的組件關係
- ✅ 無反射魔法

### 🧪 開發

#### 運行測試

執行單元和集成測試：

```bash
# 執行所有測試
go test ./...

# 運行覆蓋率測試
go test -cover ./...

# 運行特定套件的測試
go test ./internal/repository

# 詳細輸出
go test -v ./...
```

**說明**：測試需要 `.env.test` 檔案，應用程式會使用 `util.GetProjectRoot()` 函式自動加載此檔案。

#### 測試數據庫設置

```bash
# 建立測試數據庫
psql -U postgres -c "CREATE DATABASE blockscope_test;"

# 使用 golang-migrate 執行遷移
migrate -path migrations -database "postgres://postgres:password@localhost:5432/blockscope_test?sslmode=disable" up

# 或手動執行遷移
psql -U postgres -d blockscope_test -f migrations/000001_create_daily_active_address.up.sql
```

#### 代碼結構約定

- `*_test.go` 檔案用於單元/集成測試
- 表驅動測試用於驗證器/處理器測試
- 模擬儲存庫用於業務邏輯層隔離
- 在儲存庫層進行集成測試

### 📚 模塊說明

| 模塊 | 路徑 | 責任 |
|------|------|------|
| **Config** | `internal/config/` | 加載和管理基於環境的配置 |
| **DB** | `internal/db/` | 通過 GORM 初始化數據庫連接 |
| **Domain** | `internal/domain/` | 業務實體模型（DailyActiveAddress） |
| **Repository** | `internal/repository/` | 數據持久化和 CRUD 操作 |
| **Service** | `internal/service/` | 核心業務邏輯和驗證 |
| **Transport/HTTP** | `internal/transport/http/` | HTTP 處理器、路由和請求/響應模型 |
| **Validator** | `internal/validator/` | 請求參數驗證 |
| **Util** | `internal/util/` | 工具函式（路徑解析等） |

### 📄 許可證

[指定您的許可證]

### 👤 作者

Yintc123

### 📌 設計決策與最佳實踐

#### Context 生命週期管理

**相關代碼位置**
- 位置：[internal/transport/http/handler/healthcheck_handler.go](internal/transport/http/handler/healthcheck_handler.go)
- 方法：`Check()`
- 依賴：[internal/service/healthcheck_service.go](internal/service/healthcheck_service.go) 中的 `CheckDB(ctx context.Context)`

**背景**

`CheckDB()` 方法需要透過 `sqlDB.PingContext(ctx)` 對資料庫進行健康檢查，因此需要傳入一個有效的 context 物件以控制該操作的生命週期。

**常見問題**

**Q1：使用 `context.Background()` 會造成資源浪費嗎？**

不會。`context.Background()` 建立的是一個輕量級的根 context，具有以下特點：

- 幾乎不佔用記憶體資源
- 不啟動額外的 goroutine
- 不持有外部資源（如数据库連接、檔案句柄等）
- 是一個不可取消、無超時限制的 root context

**Q2：頻繁呼叫 API 時，重複建立 context 物件是否會累積？**

不會。Go 的垃圾回收機制（GC）會自動回收未被引用的 context 物件，因此不存在累積問題。

**推薦優化方案**

在實務應用中，應直接使用 Gin 框架提供的 request context，而非 `context.Background()`：

```go
func (handler *HealthcheckHandler) Check(c *gin.Context) {
	ctx := c.Request.Context()  // 使用 HTTP 請求的 context
	
	dbErr := handler.service.CheckDB(ctx)
	// ...
}
```

**優化的優勢**

| 優勢 | 說明 |
|------|------|
| **資源效率** | 避免為每次請求建立新的 context 物件 |
| **生命週期同步** | HTTP 請求被取消或逾時時，下游的資料庫操作也會自動停止 |
| **追蹤一致性** | 保留 HTTP 請求的追蹤上下文信息，便於分散式追蹤和監控 |

#### 測試環境變數加載策略

**相關代碼位置**
- 位置：[internal/config/config.go](internal/config/config.go)
- 方法：`LoadConfig(env string)`
- 工具函式：[internal/util/path.go](internal/util/path.go) 中的 `GetProjectRoot(skip int, targetFileName string)`

**背景**

執行 `go test` 時會發現環境變數無法載入（皆為空值）。這是因為測試檔的啟動工作目錄為測試檔所在的資料夾，而非項目根目錄。如果直接在測試中引入 config 模組使用 `LoadConfig("test")` 欲載入根目錄的 `.env.test`，會因工作目錄差異而找不到環境檔案。

**常見問題**

**Q1：為什麼不在測試中直接使用相對路徑載入 .env.test？**

因為 `go test` 執行時的工作目錄是測試檔所在的目錄，不是項目根目錄。例如在 `repository` 目錄運行測試時，相對路徑會從 `internal/repository` 開始查找，導致無法找到項目根目錄的 `.env.test` 檔案。

**推薦優化方案**

創建一個 `util` 模組，使用 `runtime.Caller()` 動態取得調用者的檔案路徑，然後向上遍歷目錄樹直到找到目標檔案：

**skip 參數說明**

`skip` 參數用於指定從呼叫堆棧中取得哪一層檔案的路徑：

| skip 值 | 說明 | 返回的檔案 |
|---------|------|----------|
| 0 | 當前函式所在的檔案 | `path.go`（`GetProjectRoot` 函式本身） |
| 1 | 直接呼叫者的檔案 | 呼叫 `GetProjectRoot` 的檔案（如 `config.go`） |
| 2 | 呼叫者的呼叫者的檔案 | 呼叫 `config.go` 的檔案 |

**實用例子：**
```
config.go (skip=1) 呼叫 → GetProjectRoot(1) → 返回 config.go 的路徑
                                ↓
                         path.go (skip=0) 的檔案本身
```

**為什麼建議 skip 參數設定為 0？**

在配置模組中呼叫時，應傳入 `skip=0`，原因如下：

1. **堆棧層級確定性**：不同的調用者可能呼叫 `GetProjectRoot()`（在測試中呼叫、在其他模組呼叫等），若使用 `skip=1` 需要知道確切的堆棧層級，容易出錯。使用 `skip=0` 則直接從 `GetProjectRoot()` 函式本身開始計算，不受呼叫鏈深度影響。

2. **向上遍歷的基準點**：使用 `skip=0` 時，`GetProjectRoot()` 會從 `path.go` 的目錄開始向上查找目標檔案。無論呼叫者在哪裡，最終都會向上遍歷到項目根目錄，確保能找到 `.env.test` 等檔案。

3. **API 的通用性**：若在多個模組（`config.go`、`db.go` 等）呼叫 `GetProjectRoot()`，使用 `skip=0` 無需調整，統一使用同一個值，降低出錯風險。

```go
// internal/util/path.go
func GetProjectRoot(skip int, targetFileName string) string {
	// runtime.Caller 動態定位檔案路徑
	_, b, _, _ := runtime.Caller(skip)
	dir := filepath.Dir(b)

	// 向上層尋找，直到找到目標檔案
	for {
		if _, err := os.Stat(filepath.Join(dir, targetFileName)); err == nil {
			return dir
		}
		
		parent := filepath.Dir(dir)
		if parent == dir {
			break // 已到達系統頂層
		}
		dir = parent
	}

	// 找不到時回傳當前執行目錄
	cwd, _ := os.Getwd()
	return cwd
}
```

在配置模組中使用：

```go
// internal/config/config.go
case "test":
	var rootPath string = util.GetProjectRoot(0, ".env.test")
	err := godotenv.Load(filepath.Join(rootPath, ".env.test"))
	if err != nil {
		return nil, fmt.Errorf("could not load .env.test from %s: %w", rootPath, err)
	}
```

**優化的優勢**

| 優勢 | 說明 |
|------|------|
| **工作目錄無關** | 不依賴執行時工作目錄，測試可從任意位置執行 |
| **動態定位** | 自動找到正確的項目根目錄，無須手動配置路徑 |
| **可擴展性** | 同樣的機制可用於其他環境變數檔案（如 `.env.staging`） |
