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
	CreateApp(app *models.App) error
	GetAppByAppID(appID string) (*models.App, error)
	GetActiveAppByAppID(appID string) (*models.App, error)
	UnDeleteAppByAppID(appID string) error
	GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error)
	GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error)
	GetDeviceByVersionAndAppID(versionID string, appID string) (*models.Device, error)
	CreateNewVersion(version *models.Version) error
	IncrementVersionTotals(versionID string, isNewInstall bool) error
	InsertDeviceOrUpdateVersionID(device *models.Device) error
}
