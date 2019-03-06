package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

// Repository represent the app's repository contract
type Repository interface {
	GetApps() (*[]models.App, error)
	GetActiveAppByID(ID string) (*models.App, error)
	GetAppVersionsByAppID(ID string) (*[]models.Version, error)
	UpdateAppVersions(versions []models.Version) error
	DisableAllAppVersionsByAppID(appID string, message string) error
	DeleteAppByAppID(appId string) error
	CreateApp(id, appId, name string) error
	GetAppByAppID(appID string) (*models.App, error)
	UnDeleteAppByAppID(appID string) error
}
