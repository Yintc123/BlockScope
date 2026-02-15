package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var dataSourceName string = "host=localhost user=postgres password=postgres dbname=blockscope port=5432 sslmode=disable"

	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to connect db: ", err)
	}

	DB = db
}
