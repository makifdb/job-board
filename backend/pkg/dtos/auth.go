package dtos

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
