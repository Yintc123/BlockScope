package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Yintc123/BlockScope/internal/util"
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

func getEnv(key string) (string, error) {
	var value string = os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("Missing required env: %s", key)
	}

	return value, nil
}

func getEnvWIthDefault(key string, defaultValue string) string {
	var value string = os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	var valueFromEnv string = os.Getenv(key)
	if valueFromEnv == "" {
		return defaultValue
	}
	var valueInt int
	var err error
	valueInt, err = strconv.Atoi(valueFromEnv)
	if err != nil {
		return defaultValue
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
func LoadConfig(env string) (*Config, error) {
	env = normalizeEnv(env)

	switch env {
	case "production":
		godotenv.Load(".env")
	case "local":
		godotenv.Load(".env.local")
	case "test":
		// 要依據 path.go 的位置尋找 .env.test 的檔案路徑，故傳入 skip = 0
		var rootPath string = util.GetProjectRoot(0, ".env.test")
		// 由於 go test 載入 .env.test 的位置是測試檔的位置，故需要回到根目錄載入 .env.test
		err := godotenv.Load(filepath.Join(rootPath, ".env.test"))
		if err != nil {
			return nil, fmt.Errorf("could not load .env.test from %s: %w", rootPath, err)
		}
	}

	// 關鍵參數建議使用 getEnv 進行嚴格檢查
	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	// 共用參數
	var config *Config = &Config{
		App: AppConfig{
			Name: "BlockScope",
			// Port: "8080",
			ChainList: []string{"eth", "btc", "sol"},
		},
		DB: DBConfig{
			Driver:   "postgres",
			Host:     dbHost,
			Port:     getEnvInt("DB_PORT", 5432),
			Name:     getEnvWIthDefault("DB_NAME", "blockscope"),
			User:     getEnvWIthDefault("DB_USER", "postgres"),
			Password: os.Getenv("DB_PASSWORD"), // 密碼允許為空
			SSLMode:  getEnvWIthDefault("DB_SSLMODE", "disable"),
		},
	}

	// 環境差異
	switch env {
	case "production":
		config.App.Port = 80
	default:
		config.App.Port = 8080
	}

	return config, nil
}
