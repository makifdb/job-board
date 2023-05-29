package dtos

type JobCreateDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	CompanyName string   `json:"company_name"`
	CompanyLogo string   `json:"company_logo"`
	Tags        []string `json:"tags"`
}

type JobUpdateDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	CompanyName string   `json:"company_name"`
	CompanyLogo string   `json:"company_logo"`
	Tags        []string `json:"tags"`
}

type JobDTO struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	CompanyName string   `json:"company_name"`
	CompanyLogo string   `json:"company_logo"`
	Source      string   `json:"source"`
	Tags        []string `json:"tags"`
}
