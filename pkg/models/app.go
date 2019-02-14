package models

// App is the model struct for apps
// swagger:model App
type App struct {
	ID                    string     `json:"id,omitempty"`
	AppID                 string     `json:"appId,omitempty"`
	AppName               string     `json:"appName,omitempty"`
	NumOfDeployedVersions int        `json:"numOfDeployedVersions"`
	NumOfClients          int        `json:"numOfClients"`
	NumOfAppLaunches      int64      `json:"numOfAppLaunches"`
	DeployedVersions      *[]Version `json:"deployedVersions,omitempty"`
}
