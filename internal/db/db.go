package db

import (
	"errors"
	"fmt"

	"github.com/Yintc123/BlockScope/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	switch dbConfig.Driver {
	case "postgres":
		var dataSourceName string = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode,
		)
		db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		// 檢查連線是否可用
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		if err := sqlDB.Ping(); err != nil {
			return nil, fmt.Errorf("db ping failed: %w", err)
		}

		return db, nil
	default:
		return nil, errors.New("unsupport db driver.")
	}
}
