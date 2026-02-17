package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Name      string
	Port      string
	ChainList []string
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// LoadConfig 初始化 Config，根據 env 選擇 local / production
func LoadConfig(env string) *Config {
	var envToLower = strings.ToLower(env)
	if envToLower != "production" {
		envToLower = "local"
	}

	godotenv.Load(".env." + envToLower)

	var commonChains []string = []string{"eth", "btc", "sol"}
	var commonAppName string = "BlockScope"

	// 共用參數
	var config *Config = &Config{
		App: AppConfig{
			Name: commonAppName,
			// Port: "8080",
			ChainList: commonChains,
		},
		DB: DBConfig{
			Driver:  "postgres",
			SSLMode: "disable",
		},
	}

	// 環境差異
	switch strings.ToLower(env) {
	case "local":
		config.App.Port = "8080"
		config.DB.Host = os.Getenv("DB_HOST")
		config.DB.Port = os.Getenv("DB_PORT")
		config.DB.Name = os.Getenv("DB_NAME")
		config.DB.User = os.Getenv("DB_USER")
		config.DB.Password = os.Getenv("DB_PASSWORD")
	case "production":
		config.App.Port = "80"
		config.DB.Host = os.Getenv("DB_HOST")
		config.DB.Port = os.Getenv("DB_PORT")
		config.DB.Name = os.Getenv("DB_Name")
		config.DB.User = os.Getenv("DB_USER")
		config.DB.Password = os.Getenv("DB_PASSWORD")
	}

	return config
}
