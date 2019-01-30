package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/httperrors"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps(ctx echo.Context) (*[]models.App, error)
	}

	appsService struct {
		psqlRepository PostgreSQLRepository
	}
)

// NewService instantiates this service
func NewService(psqlRepository PostgreSQLRepository) Service {
	return &appsService{
		psqlRepository: psqlRepository,
	}
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) GetApps(c echo.Context) (*[]models.App, error) {
	apps, err := a.psqlRepository.GetApps(c)

	if err != nil {
		return nil, httperrors.NotFound(c, "No apps found")
	}

	return apps, nil
}
