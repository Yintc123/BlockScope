package service

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type HealthcheckService struct {
	db *gorm.DB
}

func NewHealthCheckService(db *gorm.DB) *HealthcheckService {
	return &HealthcheckService{db: db}
}

// 檢查 server 是否有回應
func (service *HealthcheckService) CheckServer() bool {
	return true
}

// 檢查 DB 是否能連線，設計概念為“err == nil，回應沒報錯誤，就可以判斷 DB 能正常連線”
func (service *HealthcheckService) CheckDB(ctx context.Context) error {
	var sqlDB *sql.DB
	var err error
	sqlDB, err = service.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
