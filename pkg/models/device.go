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

// NewDevice returns a new Device model
func NewDevice(sdkInfo *Device, app *App, version *Version) *Device {
	dev := new(Device)
	dev.VersionID = version.ID
	dev.Version = version.Version
	dev.AppID = app.ID
	dev.DeviceVersion = sdkInfo.DeviceVersion
	dev.DeviceType = sdkInfo.DeviceType
	return dev
}
