package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps(ctx echo.Context) (*[]models.App, error)
	}

	appsService struct {
		repository Repository
	}
)

// NewService instantiates this service
func NewService(repository Repository) Service {
	return &appsService{
		repository: repository,
	}
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) GetApps(c echo.Context) (*[]models.App, error) {
	apps, err := a.repository.GetApps(c)

	if err != nil {
		return nil, models.ErrNotFound
	}

	return apps, nil
}
