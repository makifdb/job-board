package entities

type User struct {
	Base
	Name           string `json:"name"`
	Email          string `json:"email"`
	OrganizationID int64  `json:"organization_id"`
}
