// +build integration

package db

import (
	"database/sql"
	"testing"

	"github.com/aerogear/mobile-security-service/pkg/config"
	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	config := config.Get()

	type args struct {
		connString     string
		maxConnections int
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		{
			name: "Connect() should successfully make a connection to the database and return the connection",
			args: args{
				config.DB.ConnectionString,
				config.DB.MaxConnections,
			},
			wantErr: false,
		},
		{
			name: "Connect() should return an error when invalid database connection string is supplied",
			args: args{
				"connect_timeout=5 dbname=mobile_security_service_test host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				config.DB.MaxConnections,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbConn, err := Connect(tt.args.connString, tt.args.maxConnections)

			if dbConn == nil && !tt.wantErr {
				t.Errorf("Connect() expected database connection to be returned")
			}

			if dbConn != nil && !tt.wantErr {
				err = dbConn.Ping()

				if err != nil {
					t.Errorf("Could not ping database after successfully connecting: %v", err)
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	config := config.Get()

	type args struct {
		connString     string
		maxConnections int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantDB  bool
	}{
		{
			name: "Setup() should return an error when using invalid database connection",
			args: args{
				connString:     "connect_timeout=5 dbname=mobile_security_service_test host=localhost password=postgres port=5432 sslmode=disable user=postgresql",
				maxConnections: 10,
			},
			wantErr: true,
		},
		{
			name: "Setup() should be return a valid database connection",
			args: args{
				connString:     config.DB.ConnectionString,
				maxConnections: config.DB.MaxConnections,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, err := Connect(tt.args.connString, tt.args.maxConnections)

			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := Setup(db); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}

			// These tests only need to be checked if we have a valid database connection
			if db != nil {
				var exists bool

				err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'app');").Scan(&exists)

				if err != nil {
					t.Errorf("Database returned an error while checking if table exists: %v", err.Error())
				}

				if !exists {
					t.Error("Expected table app does not exist")
				}

				err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'version');").Scan(&exists)

				if err != nil {
					t.Errorf("Database returned an error while checking if table exists: %v", err.Error())
				}

				if !exists {
					t.Error("Expected table version does not exist")
				}
			}
		})
	}
}
