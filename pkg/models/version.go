package models

// Version model
// swagger:model Version
type Version struct {
	ID                   string   `json:"id"`
	Version              string   `json:"version"`
	AppID                string   `json:"appId"`
	Disabled             bool     `json:"disabled"`
	DisabledMessage      string   `json:"disabledMessage,omitempty"`
	NumOfCurrentInstalls int64    `json:"numOfCurrentInstalls"`
	NumOfAppLaunches     int64    `json:"numOfAppLaunches"`
	Devices              []Device `json:"devices,omitempty"`
	LastLaunchedAt       string   `json:"lastLaunchedAt,omitempty"`
}

func NewVersion(sdkInfo *Device, app *App) *Version {
	ver := new(Version)
	ver.Version = sdkInfo.Version
	ver.AppID = app.AppID
	return ver
}
