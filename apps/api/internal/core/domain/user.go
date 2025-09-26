package domain

import "time"

type User struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
}

type UserFront struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
}

func (u *User) ToFront() *UserFront {
	return &UserFront{
		Id:        u.Id,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Role:      u.Role,
	}
}
