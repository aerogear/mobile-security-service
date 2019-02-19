package apps

import (
	"database/sql"
	"github.com/aerogear/mobile-security-service/models"
)

type (
	// PostgreSQLRepository interface defines the methods to be implemented
	appsPostgreSQLRepository struct {
		db *sql.DB
	}
)

// NewPostgreSQLRepository creates a new instance of appsPostgreSQLRepository
func NewPostgreSQLRepository(db *sql.DB) Repository {
	return &appsPostgreSQLRepository{db}
}

// GetApps retrieves all apps from the database
func (a *appsPostgreSQLRepository) GetApps() (*[]models.App, error) {
	app1 := models.App{
		ID:                    "a0874c82-2b7f-11e9-b210-d663bd873d93",
		AppID:                 "com.aerogear.app1",
		AppName:               "app1",
		NumOfDeployedVersions: 1,
		NumOfClients:          1,
		NumOfAppLaunches:      1,
	}

	apps := []models.App{app1}

	return &apps, nil
}
