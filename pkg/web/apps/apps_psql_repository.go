package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

type (
	// PostgreSQLRepository interface defines the methods to be implemented
	PostgreSQLRepository interface {
		GetApps(c echo.Context) (*[]models.App, error)
	}

	appsPostgreSQLRepository struct {
		// TODO: Add Db connection
	}
)

// NewPostgreSQLRepository creates a new instance of appsPostgreSQLRepository
func NewPostgreSQLRepository() PostgreSQLRepository {
	return &appsPostgreSQLRepository{}
}

// GetApps retrieves all apps from the database
func (a *appsPostgreSQLRepository) GetApps(c echo.Context) (*[]models.App, error) {
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
