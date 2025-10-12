package userservice

import (
	"GoNext/base/internal/core/ports"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}
