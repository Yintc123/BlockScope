package config

import (
	"os"
	"testing"
)

func LoadTestConfig(test *testing.T) *Config {
	test.Helper()
	os.Setenv("APP_ENV", "test")
	return LoadConfig("test")
}
