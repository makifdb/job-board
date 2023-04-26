package entities

type Company struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	LocationID  int64  `json:"location_id"`
}
