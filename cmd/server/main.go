package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yintc123/BlockScope/internal/config"
	"github.com/Yintc123/BlockScope/internal/db"
	"github.com/Yintc123/BlockScope/internal/repository"
	"github.com/Yintc123/BlockScope/internal/service"
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	http "github.com/Yintc123/BlockScope/internal/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 取得環境變數，預設為 local
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// 2. 啟動依賴注入與伺服器
	router, port, err := bootstrap(env)
	if err != nil {
		log.Fatalf("Critical boot error: %v", err)
	}

	// 啟動 Gin server
	log.Printf("Server starting in [%s] mode on port %s", env, port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}

// bootstrap 封裝所有初始化邏輯，負責初始化 config / db / repository / service / handler，提升可測試性
func bootstrap(env string) (*gin.Engine, string, error) {
	// A. 載入 config
	cfg, err := config.LoadConfig(env)
	if err != nil {
		return nil, "", fmt.Errorf("config error: %w", err)
	}

	// B. 初始化 DB
	dbConn, err := db.NewDB(cfg.DB)
	if err != nil {
		return nil, "", fmt.Errorf("db error: %w", err)
	}

	// C. 依賴注入(Dependency Injection，DI)
	// 每日活躍地址模組
	repo := repository.NewDailyActiveAddressRepository(dbConn)
	svc := service.NewDailyActiveAddressService(repo)
	statsHandler := handler.NewStatsHandler(svc)

	// 健康檢查模組
	healthSvc := service.NewHealthCheckService(dbConn)
	healthHandler := handler.NewHealthcheckHandler(healthSvc)

	// D. 初始化 router
	router := gin.Default()

	// 註冊根路由
	http.RegisterRootRoutes(
		router,
		healthHandler,
	)

	// 註冊 stats 路由
	statsGroup := router.Group("/stats")
	http.RegisterStatsRoutes(
		statsGroup,
		statsHandler,
	)

	return router, fmt.Sprintf("%d", cfg.App.Port), nil
}
