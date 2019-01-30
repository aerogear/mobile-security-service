package models

// App is the model struct for apps
type App struct {
	ID                    int64        `json:"id"`
	AppID                 string       `json:"appId"`
	AppName               string       `json:"appName"`
	NumOfDeployedVersions int64        `json:"numOfDeployedVersions,omitempty"`
	NumOfClients          int64        `json:"numOfClients,omitempty"`
	NumOfAppLaunches      int64        `json:"numOfAppLaunches,omitempty"`
	DeployedVersions      []AppVersion `json:"deployedVersions,omitempty"`
}
