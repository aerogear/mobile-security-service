package models

// App is the model struct for apps
// swagger:model App
type App struct {
	ID                    string     `json:"id"`
	AppID                 string     `json:"appId"`
	AppName               string     `json:"appName"`
	NumOfDeployedVersions *int       `json:"numOfDeployedVersions,omitempty"`
	NumOfClients          *int       `json:"numOfClients"`
	NumOfAppLaunches      *int       `json:"numOfAppLaunches"`
	DeployedVersions      *[]Version `json:"deployedVersions,omitempty"`
}
