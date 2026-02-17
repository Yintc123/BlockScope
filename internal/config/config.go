package config

import (
	"os"
	"strconv"
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
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func mustEnvInt(key string) int {
	var valueFromEnv string = os.Getenv(key)
	if valueFromEnv == "" {
		panic("Missing env: " + key)
	}
	var valueInt int
	var err error
	valueInt, err = strconv.Atoi(valueFromEnv)
	if err != nil {
		panic("invalid int env: " + key)
	}

	return valueInt
}

// LoadConfig 初始化 Config，根據 env 選擇 local / production
func LoadConfig(env string) *Config {
	var envToLower = strings.ToLower(env)
	if envToLower != "production" {
		envToLower = "local"
	}

	// godotenv.Load(".env." + envToLower)
	switch envToLower {
	case "local":
		godotenv.Load(".env.local")
	case "production":
		godotenv.Load(".env")
	}

	var commonChains []string = []string{"eth", "btc", "sol"}
	var commonAppName string = "BlockScope"
	// 由於 env 取得的值一定為 string，需要轉型為 int
	var dbPort int = mustEnvInt(os.Getenv("DB_PORT"))

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
	switch envToLower {
	case "local":
		config.App.Port = "8080"
		config.DB.Host = os.Getenv("DB_HOST")
		config.DB.Port = dbPort
		config.DB.Name = os.Getenv("DB_NAME")
		config.DB.User = os.Getenv("DB_USER")
		config.DB.Password = os.Getenv("DB_PASSWORD")
	case "production":
		config.App.Port = "80"
		config.DB.Host = os.Getenv("DB_HOST")
		config.DB.Port = dbPort
		config.DB.Name = os.Getenv("DB_NAME")
		config.DB.User = os.Getenv("DB_USER")
		config.DB.Password = os.Getenv("DB_PASSWORD")
	}

	return config
}
