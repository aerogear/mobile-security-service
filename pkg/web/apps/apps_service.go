package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
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
		InitClientApp(sdkInfo *models.Device) (*models.Version, error)
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
	app, err := a.repository.GetAppByAppID(appId)

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

func (a *appsService) checkDeviceWithThisVersion(sdkInfo *models.Device, hasVersion bool, hasApp bool) (bool, *models.Device, error) {
	// Has no need to do the queries if has not an app or version in the DB with the SDK info
	if !hasApp || !hasVersion {
		return false, nil, nil
	}

	device, err := a.repository.GetDeviceByVersionAndAppID(sdkInfo.Version, sdkInfo.AppID)
	if err != nil && err != models.ErrNotFound {
		return false, nil, nil
	}

	return true, device, nil
}

func (a *appsService) checkVersionForThisApp(sdkInfo *models.Device, hasApp bool) (bool, *models.Version, error) {
	// Has no need to do the queries if has not an app in the DB with the SDK info
	if !hasApp {
		return false, nil, nil
	}

	version, err := a.repository.GetVersionByAppIDAndVersion(sdkInfo.AppID, sdkInfo.Version)
	if err != nil && err != models.ErrNotFound {
		return false, nil, nil
	}
	return true, version, nil
}

func (a *appsService) checkAnyVersionOfAppForThisDevice(sdkInfo *models.Device, hasVersion bool, hasApp bool) (bool, *models.Device, error) {
	// Has no need to do the queries if has not an app or version in the DB with the SDK info
	if !hasApp || hasVersion {
		return false, nil, nil
	}

	device, err := a.repository.GetDeviceByDeviceIDAndAppID(sdkInfo.DeviceID, sdkInfo.AppID)
	if err != nil && err != models.ErrNotFound {
		return false, nil, nil
	}
	return true, device, nil
}

func (a *appsService) checkIfHasApp(sdkInfo *models.Device) (bool, *models.App, error) {
	app, err := a.repository.GetAppByAppID(sdkInfo.AppID)

	if err != nil && err != models.ErrNotFound {
		return false, nil, nil
	}
	return true, app, nil
}

// Init call made from the SDK
// Response: Version details
// If the app, version and or device are not tracked in the database then it will be made
func (a *appsService) InitClientApp(sdkInfo *models.Device) (*models.Version, error) {
	hasApp, app, err := a.checkIfHasApp(sdkInfo)
	if err != nil {
		return nil, err
	}

	//If the app be deleted then is required re-active it
	if hasApp && app.DeletedAt != "" {
		a.repository.UnDeleteAppByAppID(app.AppID)
		if err != nil {
			return nil, err
		}
	}

	hasVersion, version, err := a.checkVersionForThisApp(sdkInfo, hasApp)
	if err != nil {
		return nil, err
	}

	//Check if the version is not disable before continue
	// If it is disable than just response it
	if hasVersion && version.Disabled {
		// It is to return just the valid information for the device
		response, err := a.preperResponse(version)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	hasAnyOtherVersion, deviceDiffVersion, err := a.checkAnyVersionOfAppForThisDevice(sdkInfo, hasVersion, hasApp)
	if err != nil {
		return nil, err
	}

	hasDeviceVersion, deviceWithVersion, err := a.checkDeviceWithThisVersion(sdkInfo, hasVersion, hasApp)
	if err != nil {
		return nil, err
	}

	//** Launch Scenario **
	// If has a version and has tracked a device using this version then it is the launch scenario
	if hasDeviceVersion {
		if err = a.repository.IncrementVersionTotals(deviceWithVersion.VersionID, false); err != nil {
			return nil, err
		}
	}

	//** New install Scenario **
	//If has a version but has NOT tracked a device using this version then it is the new install
	if hasVersion && !hasDeviceVersion {
		if err := a.repository.InsertDeviceOrUpdateVersionID(sdkInfo); err != nil {
			return nil, err
		}
		a.repository.IncrementVersionTotals(version.ID, true)
	}

	// ** New version and new install Scenario **
	// If has NOT the version then it is a new device install of a new version
	if !hasVersion && !hasDeviceVersion {
		if err = a.handleNewVersionAndInstall(sdkInfo, deviceDiffVersion, app, hasAnyOtherVersion); err != nil {
			return nil, err
		}
	}

	//** New APP Scenario **
	if !hasVersion && !hasDeviceVersion && !hasApp {
		// Create new App with the data sent froom the SDK
		if err = a.repository.CreateApp(helpers.GetUUID(), sdkInfo.AppID, sdkInfo.AppName); err != nil {
			return nil, err
		}
		// Create the new Version, and Device for this call
		app, err := a.repository.GetAppByAppID(app.AppID)
		if err = a.handleNewVersionAndInstall(sdkInfo, deviceDiffVersion, app, hasAnyOtherVersion); err != nil {
			return nil, err
		}
	}

	// It is to return just the valid information for the device
	response, err := a.preperResponse(version)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func (a *appsService) handleNewVersionAndInstall(sdkInfo, deviceDiffVersion *models.Device, app *models.App, hasAnyOtherVersion bool) error {
	version := models.NewVersion(sdkInfo, app)
	version.ID = helpers.GetUUID()

	if err := a.repository.CreateNewVersion(version); err != nil {
		return err
	}
	if err := a.repository.InsertDeviceOrUpdateVersionID(sdkInfo); err != nil {
		return err
	}
	a.repository.IncrementVersionTotals(version.ID, true)
	return nil
}

//Remove the data info that should be not returned to the SDK
func (a *appsService) preperResponse(version *models.Version) (*models.Version, error) {
	// clear these values before returning the data
	version.LastLaunchedAt = ""
	version.NumOfAppLaunches = 0
	version.NumOfCurrentInstalls = 0
	return version, nil
}
