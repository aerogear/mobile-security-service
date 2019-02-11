package models

// App is the model struct for apps
// swagger:model App
type App struct {
	ID                    string    `json:"id"`
	AppID                 string    `json:"appId"`
	AppName               string    `json:"appName"`
	NumOfDeployedVersions int64     `json:"numOfDeployedVersions,omitempty"`
	NumOfClients          int64     `json:"numOfClients,omitempty"`
	NumOfAppLaunches      int64     `json:"numOfAppLaunches,omitempty"`
	DeployedVersions      []Version `json:"deployedVersions,omitempty"`
}
