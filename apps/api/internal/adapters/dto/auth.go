package dto

type UserCredentials struct {
	Firstname string `json:"firstname" validate:"required,max=100"`
	Lastname  string `json:"lastname" validate:"required,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
}

type RegisterCredentials struct {
	UserCredentials
	Confirm string `json:"confirm" validate:"required,password"`
}
