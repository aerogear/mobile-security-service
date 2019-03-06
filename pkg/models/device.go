package models

// Device model
// swagger:model Device
type Device struct {
	ID            string `json:"id"`
	VersionID     string `json:"versionId"`
	Version       string `json:"version"`
	AppID         string `json:"appId"`
	DeviceID      string `json:"deviceId"`
	DeviceVersion string `json:"deviceVersion"`
	DeviceType    string `json:"deviceType"`
	AppName       string `json:"appName,omitempty"`
}
