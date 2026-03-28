package app

import (
	"github.com/Yintc123/BlockScope/internal/config"
	"github.com/Yintc123/BlockScope/internal/db"
	"github.com/Yintc123/BlockScope/internal/repository"
	"github.com/Yintc123/BlockScope/internal/service"
	"gorm.io/gorm"
)

// 負責初始化所有參數
type AppDependencies struct {
	Config   *config.Config
	DB       *gorm.DB
	Repos    *repository.Repositories
	Services *service.Services
	// Handlers *handler.Handlers // crawler 不需要 handlers
}

func SetupDependencies(env string) (*AppDependencies, error) {
	// 載入配置
	cfg, err := config.LoadConfig(env)
	if err != nil {
		return nil, err
	}
	// 初始化 DB，依賴注入(Dependency Injection，DI)
	db, err := db.NewDB(cfg.DB)
	if err != nil {
		return nil, err
	}
	// 初始化 repos，依賴注入(Dependency Injection，DI)
	var repos *repository.Repositories = repository.NewRepositories(db)
	// 初始化 services，依賴注入(Dependency Injection，DI)
	var services *service.Services = service.NewServices(repos, db)
	// 初始化 handlers，依賴注入(Dependency Injection，DI)
	// var handlers *handler.Handlers = handler.NewHandlers(services)
	return &AppDependencies{
		Config:   cfg,
		DB:       db,
		Repos:    repos,
		Services: services,
		// Handlers: handlers,
	}, nil
}
