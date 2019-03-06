package models

import "github.com/google/uuid"

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

func NewVersionByDevice(sdkInfo *Device) *Version {
	ver := new(Version)
	ver.ID = uuid.New().String()
	ver.Version = sdkInfo.Version
	ver.AppID = sdkInfo.AppID
	return ver
}
