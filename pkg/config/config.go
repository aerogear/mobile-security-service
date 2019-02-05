package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ListenAddress  string
	LogLevel       string
	LogFormat      string
	CORS           CORSConfig
	StaticFilesDir string
	ApiRoutePrefix string
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowCredentials bool
}

func Get() Config {
	return Config{
		ListenAddress: fmt.Sprintf(":%v", getEnvInt("PORT", 3000)),
		LogLevel:      strings.ToLower(getEnv("LOG_LEVEL", "info")),
		LogFormat:     strings.ToLower(getEnv("LOG_FORMAT", "text")), //can be text or json
		CORS: CORSConfig{
			AllowOrigins:     getEnvSlice("ACCESS_CONTROL_ALLOW_ORIGIN", []string{"*"}, ","),
			AllowCredentials: getEnvBool("ACCESS_CONTROL_ALLOW_CREDENTIALS", false),
		},
		StaticFilesDir: getEnv("STATIC_FILES_DIR", ""),
		ApiRoutePrefix: "/api", //should start with a "/",
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

// Helper to read an environment variable into a bool or return default value
func getEnvBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
