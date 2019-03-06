package models

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

func NewApp(sdkInfo *Device) *App {
	app := new(App)
	app.AppID = sdkInfo.AppID
	app.AppName = sdkInfo.AppName
	return app
}
