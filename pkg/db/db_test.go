// +build integration

package db

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	_ "github.com/lib/pq"
)

// Connect() should successfully make a connection to the database and return the connection
func TestConnect(t *testing.T) {
	config := config.Get()
	s := config.DB.ConnectionString
	i := config.DB.MaxConnections

	// Assertions
	got, err := Connect(s, i)
	assert.NotNil(t, got)
	assert.Nil(t, err)
}

//Connect() should return an error when invalid database connection string is supplied
func TestConnectError(t *testing.T) {
	config := config.Get()
	invalidConnectionString := "connect_timeout=5 dbname=mobile_security_service_test host=localhost password=postgres port=5432 sslmode=disable user=postgresql"
	i := config.DB.MaxConnections

	// Assertions
	got, err := Connect(invalidConnectionString, i)
	assert.Nil(t, got)
	assert.NotNil(t, err)
}

func TestSetup(t *testing.T) {
	config := config.Get()
	s := config.DB.ConnectionString
	i := config.DB.MaxConnections
	db, err := Connect(s, i)
	assert.NoError(t, Setup(db))

	// Assertions
	assert.NotNil(t, db)
	assert.Nil(t, err)

	// These tests only need to be checked if we have a valid database connection
	if db != nil {
		var exists bool
		assert.NoError(t, db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'app');").Scan(&exists))
		assert.True(t, exists)
		assert.NoError(t, db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'version');").Scan(&exists))
		assert.True(t, exists)
		assert.NoError(t, db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'device');").Scan(&exists))
		assert.True(t, exists)
	}
}

//Setup() should return an error when using invalid database connection
func TestSetupError(t *testing.T) {
	config := config.Get()
	invalidConnectionString := "connect_timeout=5 dbname=mobile_security_service_test host=localhost password=postgres port=5432 sslmode=disable user=postgresql"
	i := config.DB.MaxConnections
	db, err := Connect(invalidConnectionString, i)

	assert.NotNil(t, err);
	errS := Setup(db)
	assert.NotNil(t, errS)
}