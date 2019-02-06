package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

// Repository represent the app's repository contract
type Repository interface {
	GetApps() (*[]models.App, error)
	GetAppVersionsByAppID(id string) (*[]models.Version, error)
	GetAppByID(ID string) (*models.App, error)
}
