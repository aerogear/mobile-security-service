package apps

import (
	"github.com/aerogear/mobile-security-service/models"
)

// Repository represent the app's repository contract
type Repository interface {
	GetApps() (*[]models.App, error)
}
