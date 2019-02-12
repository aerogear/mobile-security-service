package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

type (
	// PostgreSQLRepository interface defines the methods to be implemented
	appsPostgreSQLRepository struct {
		// TODO: Add Db connection
	}
)

// NewPostgreSQLRepository creates a new instance of appsPostgreSQLRepository
func NewPostgreSQLRepository() Repository {
	return &appsPostgreSQLRepository{}
}

// GetApps retrieves all apps from the database
func (a *appsPostgreSQLRepository) GetApps() (*[]models.App, error) {
	app1 := models.App{
		ID:                    1,
		AppID:                 "com.aerogear.app1",
		AppName:               "app1",
		NumOfDeployedVersions: 1,
		NumOfClients:          1,
		NumOfAppLaunches:      1,
	}

	apps := []models.App{app1}

	return &apps, nil
}
