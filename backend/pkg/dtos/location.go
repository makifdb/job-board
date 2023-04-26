package dtos

type LocationCreateDTO struct {
	Name     string   `json:"name"`
	LongName string   `json:"long_name"`
	Tags     []string `json:"tags"`
}

type LocationUpdateDTO struct {
	Name     string   `json:"name"`
	LongName string   `json:"long_name"`
	Tags     []string `json:"tags"`
}

type LocationDTO struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	LongName string   `json:"long_name"`
	Tags     []string `json:"tags"`
}
