package dtos

type CompanyCreateDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Tags        []string `json:"tags"`
}

type CompanyUpdateDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Tags        []string `json:"tags"`
}

type CompanyDTO struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Website     string   `json:"website"`
	Tags        []string `json:"tags"`
}
