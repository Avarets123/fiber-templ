package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() error {
	return godotenv.Load(".env")
}

type DatabaseCfg struct {
	Url string
}

func NewDbCfg() *DatabaseCfg {
	return &DatabaseCfg{
		Url: getEnvStr("DATABASE_URL", ""),
	}
}

func getEnvStr(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parseValue, _ := strconv.Atoi(value)
	return parseValue
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parseValue, _ := strconv.ParseBool(value)
	return parseValue
}
