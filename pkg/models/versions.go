package models

// Version is the model struct for Version
// swagger:model Version
type Version struct {
	ID               int64  `json:"id"`
	Version          string `json:"version"`
	AppID            string `json:"appId"`
	Disabled         bool   `json:"disabled"`
	DisabledMessage  string `json:"disabledMessage,omitempty"`
	NumOfClients     int64  `json:"numOfClients"`
	NumOfAppStartups int64  `json:"numOfAppStartups"`
}
