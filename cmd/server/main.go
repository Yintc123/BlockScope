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
	// bootstrap 初始化所有依賴
	router, port, err := bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	// 啟動 Gin server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("fail to run server: ", err)
	}
}

// bootstrap 負責初始化 config / db / repository / service / handler
func bootstrap() (*gin.Engine, string, error) {
	// 1. 載入 config
	cfg := config.LoadConfig(os.Getenv("APP_ENV"))

	// 2. 初始化 DB
	dbConn, err := db.NewDB(cfg.DB)
	if err != nil {
		return nil, "", fmt.Errorf("fail to connect db: %w", err)
	}

	// 3. 初始化 repository / service / handler
	repo := repository.NewDailyActiveAddressRepository(dbConn)
	svc := service.NewDailyActiveAddressService(repo)
	statsHandler := handler.NewStatsHandler(svc)

	healthSvc := service.NewHealthCheckService(dbConn)
	healthHandler := handler.NewHealthcheckHandler(healthSvc)

	// 4. 初始化 router
	router := gin.Default()
	http.RegisterRootRoutes(
		router,
		healthHandler,
	)

	statsGroup := router.Group("/stats")
	http.RegisterStatsRoutes(
		statsGroup,
		statsHandler,
	)

	return router, fmt.Sprintf("%d", cfg.App.Port), nil
}
