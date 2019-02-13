package models

import (
	"database/sql"
)

// Version model
// swagger:model Version
type Version struct {
	ID               string         `json:"id"`
	Version          string         `json:"version"`
	AppID            string         `json:"appId"`
	Disabled         bool           `json:"disabled"`
	DisabledMessage  sql.NullString `json:"disabledMessage,omitempty"`
	NumOfClients     int64          `json:"numOfClients"`
	NumOfAppLaunches int64          `json:"numOfAppLaunches"`
	Devices          []Device       `json:"devices,omitempty"`
}
