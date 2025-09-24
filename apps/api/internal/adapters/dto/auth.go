package dto

type UserCredentials struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,password"`
}

type RegisterCredentials struct {
	UserCredentials
	Confirm string `json:"confirm" validate:"required,password"`
}
