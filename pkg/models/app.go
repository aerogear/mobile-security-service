package models

// App is the model struct for apps
// swagger:model App
type App struct {
	ID                    string     `json:"id"`
	AppID                 string     `json:"appId"`
	AppName               string     `json:"appName,omitempty"`
	NumOfDeployedVersions *int       `json:"numOfDeployedVersions,omitempty"`
	NumOfCurrentInstalls  *int       `json:"numOfCurrentInstalls,omitempty"`
	NumOfAppLaunches      *int       `json:"numOfAppLaunches,omitempty"`
	DeployedVersions      *[]Version `json:"deployedVersions,omitempty"`
	DeletedAt             string     `json:"deletedAt,omitempty"`
}

// NewAppByNameAndAppID will create a new App object based on the name and appId which indeed are the only values
// according to the business that can be used/provided in order to the service create an app. This will be used in the
// systems which would consume the REST API to create new apps. ( E.g Operator )
func NewAppByNameAndAppID(name, appId string) *App {
	app := new(App)
	app.AppName = name
	app.AppID = appId
	return app
}
