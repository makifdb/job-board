package entities

type Organization struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
}
