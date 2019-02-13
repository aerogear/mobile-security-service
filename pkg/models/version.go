package models

// Version model
// swagger:model Version
type Version struct {
	ID                   string   `json:"id"`
	Version              string   `json:"version"`
	AppID                string   `json:"appId"`
	Disabled             bool     `json:"disabled"`
	DisabledMessage      string   `json:"disabledMessage,omitempty"`
	NumOfCurrentInstalls int64    `json:"numOfCurrentInstalls,omitempty"`
	NumOfAppLaunches     int64    `json:"numOfAppLaunches,omitempty"`
	LastLaunchedAt       string   `json:"lastLaunchedAt,omitempty"`
	Devices              []Device `json:"devices,omitempty"`
}
