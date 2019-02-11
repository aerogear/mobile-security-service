package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

// Repository represent the app's repository contract
type Repository interface {
	GetApps(c echo.Context) (*[]models.App, error)
}
