package config

import (
	"fmt"
	"os"
	"sort"
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
	DB             DBConfig
}

type CORSConfig struct {
	AllowOrigins     []string
	AllowCredentials bool
}

type DBConfig struct {
	ConnectionString string
	MaxConnections   int
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
		DB: DBConfig{
			ConnectionString: getDBConnectionString(),
			MaxConnections:   getEnvInt("DB_MAX_CONNECTIONS", 100),
		},
	}
}

// builds a libpq compatible connection string e.g. "user=postgresql host=localhost password=postgres
func getDBConnectionString() string {

	// These are all of the options supported by pq
	dbConnectionConfig := map[string]string{
		"dbname":                    getEnv("PGDATABASE", "mobile_security_service"),
		"user":                      getEnv("PGUSER", "postgresql"),
		"password":                  getEnv("PGPASSWORD", "postgres"),
		"host":                      getEnv("PGHOST", "localhost"),
		"port":                      getEnv("PGPORT", "5432"),
		"sslmode":                   getEnv("PGSSLMODE", "disable"),
		"connect_timeout":           getEnv("PGCONNECT_TIMEOUT", "5"),
		"fallback_application_name": getEnv("PGAPPNAME", ""),
		"sslcert":                   getEnv("PGSSLCERT", ""),
		"sslkey":                    getEnv("PGSSLKEY", ""),
		"sslrootcert":               getEnv("PGSSLROOTCERT", ""),
	}

	var options []string

	for k, v := range dbConnectionConfig {
		if dbConnectionConfig[k] != "" {
			options = append(options, fmt.Sprintf("%v=%v", k, v))
		}
	}

	// sort the list (because ordering in maps is random)
	// this ensures connection string is the same every time. Makes testing much easier
	sort.Strings(options)

	return strings.Join(options, " ")
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
