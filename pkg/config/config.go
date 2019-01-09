package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ListenAddress      string
	LogLevel           string
	LogFormat          string
}

func Get() Config {
	return Config{
		ListenAddress:      fmt.Sprintf(":%v", getEnvInt("PORT", 3000)),
		LogLevel:           strings.ToLower(getEnv("LOG_LEVEL", "info")),
		LogFormat:          strings.ToLower(getEnv("LOG_FORMAT", "text")), //can be text or json
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
