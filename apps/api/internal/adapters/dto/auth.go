package dto

type UserCredentials struct {
	Email           string `json:"email" validate:"required,email"`
	Confirm string `json:"confirm" validate:"required,password"`
	Password        string `json:"password" validate:"required,password"`
}
