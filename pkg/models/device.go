package models

// Device model
// swagger:model Device
type Device struct {
	ID            string `json:"id"`
	VersionID     string `json:"versionId"`
	DeviceID      string `json:"deviceId"`
	DeviceVersion string `json:"deviceVersion"`
	DeviceType    string `json:"deviceType"`
}
