// +build integration

package db

import (
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	config := config.Get()
	dbConn, err := Connect(config.DB.ConnectionString, config.DB.MaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %v", err.Error())
	}

	err = dbConn.Ping()

	if err != nil {
		t.Errorf("Failed to ping database after Connect(): %v", err.Error())
	}
}

func TestSetup(t *testing.T) {
	config := config.Get()
	dbConn, err := Connect(config.DB.ConnectionString, config.DB.MaxConnections)

	if err != nil {
		t.Errorf("Connect() returned an error: %v", err.Error())
	}

	err = Setup(dbConn)

	if err != nil {
		t.Errorf("Setup() returned an error: %v", err.Error())
	}

	var exists bool

	err = dbConn.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'app');").Scan(&exists)

	if err != nil {
		t.Errorf("Database returned an error while checking if table exists: %v", err.Error())
	}

	if !exists {
		t.Error("Expected table app does not exist")
	}

	err = dbConn.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'version');").Scan(&exists)

	if err != nil {
		t.Errorf("Database returned an error while checking if table exists: %v", err.Error())
	}

	if !exists {
		t.Error("Expected table version does not exist")
	}
}
