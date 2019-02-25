package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps() (*[]models.App, error)
		GetAppByID(ID string) (*models.App, error)
		UpdateAppVersions(versions []models.Version) error
		DisableAllAppVersionsByAppID(id string, message string) error
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

	return apps, nil
}

// GetAppByID retrieves app by id from the repository
func (a *appsService) GetAppByID(id string) (*models.App, error) {

	app, err := a.repository.GetAppByID(id)

	if err != nil {
		return nil, err
	}

	deployedVersions, err := a.repository.GetAppVersionsByAppID(app.AppID)
	if err != nil {
		return nil, err
	}
	app.DeployedVersions = deployedVersions

	return app, nil
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) UpdateAppVersions(versions []models.Version) error {
	err := a.repository.UpdateAppVersions(versions)

	// Check for errors and return the appropriate error to the handler
	if err != nil {
		return err
	}

	return nil
}

// Update all versions
func (a *appsService) DisableAllAppVersionsByAppID(id string, message string) error {

	// get the app id to send it to the re
	app, err := a.repository.GetAppByID(id)

	if err != nil {
		return err
	}

	return a.repository.DisableAllAppVersionsByAppID(app.AppID, message)
}
