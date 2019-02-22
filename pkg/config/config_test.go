package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	defaultConfig = Config{
		ListenAddress: ":3000",
		LogLevel:      "info",
		LogFormat:     "text",
		CORS: CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowCredentials: false,
		},
		StaticFilesDir: "",
		APIRoutePrefix: "/api",
		DB: DBConfig{
			ConnectionString: "connect_timeout=5 dbname=mobile_security_service host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
			MaxConnections:   100,
		},
	}
)

// Get() should return sensible defaults when no environment variables are set
func TestGet_DefaultConfig(t *testing.T) {
	// Assertions
	got := Get()
	assert.NotNil(t, got)
	assert.Equal(t, defaultConfig, got)
}

// Get() should override defaults with environment variables when set
func TestGet_OverrideConfig(t *testing.T) {
	// New environment variables
	envVars:= map[string]string{
		"PORT":                             "4000",
		"LOG_LEVEL":                        "error",
		"LOG_FORMAT":                       "json",
		"ACCESS_CONTROL_ALLOW_ORIGIN":      "http://localhost:1234,http://localhost:2345",
		"ACCESS_CONTROL_ALLOW_CREDENTIALS": "false",
		"STATIC_FILES_DIR":                 "static",
		"PGDATABASE":                       "mobile_security_service",
		"PGUSER":                           "postgresql",
		"PGPASSWORD":                       "postgres",
		"PGHOST":                           "localhost",
		"PGPORT":                           "5432",
		"PGSSLMODE":                        "disable",
		"PGCONNECT_TIMEOUT":                "5",
		"PGAPPNAME":                        "",
		"PGSSLCERT":                        "",
		"PGSSLKEY":                         "",
		"PGSSLROOTCERT":                    "",
		"DB_MAX_CONNECTIONS":               "100",
	}

	// Set up env for the project
	for key, val := range envVars {
		os.Setenv(key, val)
	}

	// Assertions
	got := Get()
	assert.NotNil(t, got)
	assert.NotEqual(t, defaultConfig, got)
}

// getEnv() should return default value when no environment variable is set
func TestGetEnv_OverrideValues(t *testing.T) {
	key:= "LOG_FORMAT"
	value:= "json"
	got := getEnv(key, value);

	// Assertions
	assert.NotNil(t, got)
	assert.Equal(t, got, value)
}

// getEnv() should return environment variable value when set instead of default value
func TestGetEnv_OverValues(t *testing.T) {
	key:= "LOG_FORMAT"
	newValue:= "other"

	os.Setenv(key, newValue)
	got := getEnv(key, newValue);

	// Assertions
	assert.NotNil(t, got)
	assert.Equal(t, got, newValue)
}

//getDBConnectionString() should build a connection string using environment variables
func TestGetConnectionString_OverrideValues(t *testing.T) {

	envVars:= map[string]string{
		"PGDATABASE":         "testdb",
		"PGUSER":             "postgresql1",
		"PGPASSWORD":         "postgresql",
		"PGHOST":             "127.0.0.1",
		"PGPORT":             "5431",
		"PGSSLMODE":          "disable",
		"PGCONNECT_TIMEOUT":  "5",
		"PGAPPNAME":          "",
		"PGSSLCERT":          "",
		"PGSSLKEY":           "",
		"PGSSLROOTCERT":      "",
		"DB_MAX_CONNECTIONS": "100",
	}

	for key, val := range envVars {
		os.Setenv(key, val)
	}

	got := getDBConnectionString();

	// Assertions
	assert.NotNil(t, got)
	assert.Equal(t, got, "connect_timeout=5 dbname=testdb host=127.0.0.1 password=postgresql port=5431 sslmode=disable user=postgresql1")

}

//getDBConnectionString() should build a connection string using the default values when environment variables are not set
func TestGetConnectionString_DefaultValues(t *testing.T) {

	envVars:= map[string]string{
		"PGDATABASE":         "",
		"PGUSER":             "",
		"PGPASSWORD":         "",
		"PGHOST":             "",
		"PGPORT":             "",
		"PGSSLMODE":          "",
		"PGCONNECT_TIMEOUT":  "",
		"PGAPPNAME":          "",
		"PGSSLCERT":          "",
		"PGSSLKEY":           "",
		"PGSSLROOTCERT":      "",
		"DB_MAX_CONNECTIONS": "",
	}

	for key, val := range envVars {
		os.Setenv(key, val)
	}

	got := getDBConnectionString();

	// Assertions
	assert.NotNil(t, got)
	assert.Equal(t, got, "connect_timeout=5 dbname=mobile_security_service host=localhost password=postgres port=5432 sslmode=disable user=postgresql")
}