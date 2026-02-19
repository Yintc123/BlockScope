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
	Port      int
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

func mustEnv(key string) string {
	var value string = os.Getenv(key)
	if value == "" {
		panic("Missing env: " + key)
	}

	return value
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

func normalizeEnv(env string) string {
	switch strings.ToLower(env) {
	case "production", "local", "test":
		return strings.ToLower(env)
	default:
		return "local"
	}
}

// LoadConfig 初始化 Config，根據 env 選擇 local / production
func LoadConfig(env string) *Config {
	env = normalizeEnv(env)

	switch env {
	case "production":
		godotenv.Load(".env")
	case "local":
		godotenv.Load(".env.local")
	case "test":
		godotenv.Load(".env.test")
	}

	var commonChains []string = []string{"eth", "btc", "sol"}
	var commonAppName string = "BlockScope"
	// 由於 env 取得的值一定為 string，需要轉型為 int
	var dbPort int = mustEnvInt("DB_PORT")

	// 共用參數
	var config *Config = &Config{
		App: AppConfig{
			Name: commonAppName,
			// Port: "8080",
			ChainList: commonChains,
		},
		DB: DBConfig{
			Driver:   "postgres",
			Host:     mustEnv("DB_HOST"),
			Port:     dbPort,
			Name:     mustEnv("DB_NAME"),
			User:     mustEnv("DB_USER"),
			Password: mustEnv("DB_PASSWORD"),
			SSLMode:  "disable",
		},
	}

	// 環境差異
	switch env {
	case "local":
		config.App.Port = 8080
	case "production":
		config.App.Port = 80
	}

	return config
}
