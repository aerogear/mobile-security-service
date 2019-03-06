package models

import "github.com/google/uuid"

// App is the model struct for apps
// swagger:model App
type App struct {
	ID                    string     `json:"id"`
	AppID                 string     `json:"appId"`
	AppName               string     `json:"appName"`
	NumOfDeployedVersions *int       `json:"numOfDeployedVersions,omitempty"`
	NumOfCurrentInstalls  *int       `json:"numOfCurrentInstalls,omitempty"`
	NumOfAppLaunches      *int       `json:"numOfAppLaunches,omitempty"`
	DeployedVersions      *[]Version `json:"deployedVersions,omitempty"`
	DeletedAt             string     `json:"deletedAt,omitempty"`
}

func NewAppByDevice(sdkInfo *Device) *App {
	app := new(App)
	app.ID = uuid.New().String()
	app.AppID = sdkInfo.AppID
	app.AppName = sdkInfo.AppName
	return app
}
func NewApp(appId, name string) *App {
	app := new(App)
	app.ID = uuid.New().String()
	app.AppID = appId
	app.AppName = name
	return app
}
