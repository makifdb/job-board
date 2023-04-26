package entities

import "time"

type External struct {
	Base
	Name   string `json:"name"`
	Url    string `json:"url"`
	Active bool   `json:"active"`
}

type ExternalData struct {
	Base
	Url         string    `json:"url"`
	Data        string    `json:"data"`
	CollectedAt time.Time `json:"collected_at"`
	CollectedBy string    `json:"collected_by"`
	ExternalID  int64     `json:"external_id"`
	Valid       bool      `json:"valid"`
}
