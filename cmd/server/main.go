package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Yintc123/BlockScope/internal/app"
	"github.com/Yintc123/BlockScope/internal/transport/http/handler"
	"github.com/Yintc123/BlockScope/internal/transport/http/middleware"
	http "github.com/Yintc123/BlockScope/internal/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 定義啟動模式
	mode := flag.String("mode", "api", "Mode: api, crawler or all")
	flag.Parse()

	switch *mode {
	case "api":
		runAPIServer() // 只啟動 API server
	case "crawler":
		runCrawler() // 只啟動 crawler
	case "all":
		// local 開發，同時啟動兩個實例
		// 啟動有順序性，一定要先啟動 crawler，由於 API server 是阻塞(blocking)主執行緒
		// 先啟動 crawler 在背景運作
		go runCrawler()
		// 再啟動 API server
		runAPIServer()
	default:
		log.Fatal("Invalid mode")
	}
}

func runAPIServer() {
	// 1. 取得環境變數，預設為 local
	env := getEnv()

	// 1. 取得所有依賴
	appDependencies, err := app.SetupDependencies(env)
	if err != nil {
		log.Fatalf("Critical initial dependencies error: %v", err)
	}
	var appPort string = fmt.Sprintf("%d", appDependencies.Config.App.Port)

	// 2. 初始化 handlers
	var handlers *handler.Handlers = handler.NewHandlers(appDependencies.Services)
	// 3. 初始化 router
	router, err := initRouter(handlers)
	if err != nil {
		log.Fatalf("Critical initial router error: %v", err)
	}
	// 4. 啟動 Gin server
	log.Printf("Server starting in [%s] mode on port %s", env, appPort)
	if err := router.Run(":" + appPort); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}

func initRouter(handlers *handler.Handlers) (*gin.Engine, error) {
	// 初始化 router
	router := gin.Default()
	// 註冊處理錯誤的 middleware
	router.Use(middleware.ErrorHandler())
	// 註冊根路由
	http.RegisterRootRoutes(router, handlers.Healthcheck)
	// 註冊 stats 路由
	statsGroup := router.Group("/stats")
	http.RegisterStatsRoutes(statsGroup, handlers.Stats)
	return router, nil
}

func runCrawler() {

}

func getEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}
	return env
}
