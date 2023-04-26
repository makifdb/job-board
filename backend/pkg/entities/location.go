package entities

type Location struct {
	Base
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	ParentID int64  `json:"parent_id"`
}
