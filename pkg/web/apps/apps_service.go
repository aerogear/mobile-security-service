package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/google/uuid"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps() (*[]models.App, error)
		GetActiveAppByID(ID string) (*models.App, error)
		UpdateAppVersions(versions []models.Version) error
		DisableAllAppVersionsByAppID(id string, message string) error
		UnbindingAppByAppID(appID string) error
		BindingAppByApp(appId, name string) error
		InitClientApp(deviceInfo *models.Device) (*models.Version, error)
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

// GetActiveAppByID retrieves app by id from the repository
func (a *appsService) GetActiveAppByID(id string) (*models.App, error) {

	app, err := a.repository.GetActiveAppByID(id)

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
	app, err := a.repository.GetActiveAppByID(id)

	if err != nil {
		return err
	}

	return a.repository.DisableAllAppVersionsByAppID(app.AppID, message)
}

func (a *appsService) UnbindingAppByAppID(appID string) error {
	err := a.repository.DeleteAppByAppID(appID)
	if err != nil {
		return err
	}
	return nil
}

func (a *appsService) BindingAppByApp(appId, name string) error {

	// Check if it exist
	app, err := a.repository.GetActiveAppByAppID(appId)

	// If it is new then create an app
	if err != nil && err == models.ErrNotFound {
		id := helpers.GetUUID()
		return a.repository.CreateApp(id, appId, name)
	}

	if err != nil {
		return err
	}

	// if is deleted so just reactive the existent app
	return a.repository.UnDeleteAppByAppID(app.AppID)
}

// // InitClientApp returns information about the current state of the app - its disabled status
func (a *appsService) InitClientApp(deviceInfo *models.Device) (*models.Version, error) {
	if _, err := a.repository.GetActiveAppByAppID(deviceInfo.AppID); err != nil {
		return nil, err
	}

	version, err := a.repository.GetVersionByAppIDAndVersion(deviceInfo.AppID, deviceInfo.Version)

	// If any error other Not Found error occurred, return
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	// If the version does not exist, create it
	if err == models.ErrNotFound {
		version = &models.Version{
			ID:      uuid.New().String(),
			Version: deviceInfo.Version,
			AppID:   deviceInfo.AppID,
		}
	}

	// Update the existing version or create a new one
	if err := a.repository.UpsertVersionWithAppLaunchesAndLastLaunched(version); err != nil {
		return nil, err
	}

	device, err := a.repository.GetDeviceByDeviceIDAndAppID(deviceInfo.DeviceID, deviceInfo.AppID)

	var newDevice = false
	if err != nil {
		// If we can't find the device by device ID and app ID
		if err != models.ErrNotFound {
			return nil, err
		}

		// Build a new device to save to the database
		device = models.NewDevice(version.ID, version.Version, deviceInfo.AppID, deviceInfo.DeviceID, deviceInfo.DeviceVersion, deviceInfo.DeviceType)
		newDevice = true
	}

	// indicates whether properties of the device are updated
	var isUpdated bool

	if device.VersionID != version.ID {
		device.VersionID = version.ID
		isUpdated = true
	}

	if device.DeviceVersion != deviceInfo.DeviceVersion {
		device.DeviceVersion = deviceInfo.DeviceVersion
		isUpdated = true
	}

	if newDevice || isUpdated {
		err := a.repository.InsertDeviceOrUpdateVersionID(*device)

		if err != nil {
			return nil, err
		}
	}

	// clear these values before returning the data
	version.LastLaunchedAt = ""
	version.NumOfAppLaunches = 0
	version.NumOfCurrentInstalls = 0

	return version, nil
}
