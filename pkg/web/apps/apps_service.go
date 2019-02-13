package apps

import (
	// "github.com/sirupsen/logrus"
	"github.com/aerogear/mobile-security-service/pkg/models"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps() (*[]models.App, error)
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
func (a *appsService) GetApps() (*[]models.App, error) {
	apps, err := a.repository.GetApps()

	// Check for errors and return the appropriate error to the handler
	if err != nil {
		return nil, err
	}

	// iterate over the list of apps and get the deployed versions for each
	for i, app := range *apps {
		deployedVersions, err := a.repository.GetAppVersionsByAppID(app.AppID)

		if err == nil {
			app.DeployedVersions = deployedVersions
			(*apps)[i] = app
		}
	}

	return apps, nil
}
