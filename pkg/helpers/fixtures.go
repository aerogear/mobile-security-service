package helpers

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

// GetMockAppList returns some dummy apps
func GetMockAppList() []models.App {
	apps := []models.App{
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc025",
			AppID:   "com.aerogear.mobile_app_one",
			AppName: "Mobile App One",
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc026",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Two",
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc027",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Three",
		},
	}
	return apps
}

// GetMockAppVersionList returns some dummy app versions
func GetMockAppVersionList() []models.Version {
	versions := []models.Version{
		models.Version{
			ID:                   "55ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:              "1.0",
			AppID:                "com.aerogear.mobile_app_one",
			DisabledMessage:      "Please contact an administrator",
			Disabled:             false,
			NumOfCurrentInstalls: 1,
			NumOfAppLaunches:     2,
		},
		models.Version{
			ID:                   "59ebd387-9c68-4137-a367-a12025cc1cdb",
			Version:              "1.1",
			AppID:                "com.aerogear.mobile_app_one",
			Disabled:             false,
			NumOfCurrentInstalls: 0,
			NumOfAppLaunches:     0,
		},
		models.Version{
			ID:                   "59dbd387-9c68-4137-a367-a12025cc2cdb",
			Version:              "1.0",
			AppID:                "com.aerogear.mobile_app_two",
			Disabled:             false,
			NumOfCurrentInstalls: 0,
			NumOfAppLaunches:     0,
		},
	}

	return versions
}

// GetMockAppL returns some dummy app
func GetMockApp() *models.App {
	return &models.App{
		ID:      "7f89ce49-a736-459e-9110-e52d049fc025",
		AppID:   "com.aerogear.mobile_app_one",
		AppName: "Mobile App One",
	}
}
