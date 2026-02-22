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
