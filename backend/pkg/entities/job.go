package entities

type Job struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CompanyName string `json:"company_name"`
	CompanyLogo string `json:"company_logo"`
	Source      string `json:"source"`
}
