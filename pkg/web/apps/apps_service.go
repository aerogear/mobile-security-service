package apps

import (
	"github.com/aerogear/mobile-security-service/models"
)

type (
	// AppService defines the interface methods to be used
	AppService interface {
		GetApps() (*[]models.App, error)
	}

	appsService struct {
		repository Repository
	}
)

// NewService instantiates this service
func NewService(repository Repository) AppService {
	return &appsService{
		repository: repository,
	}
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) GetApps() (*[]models.App, error) {
	apps, err := a.repository.GetApps()

	if err != nil {
		return nil, models.ErrNotFound
	}

	return apps, nil
}
