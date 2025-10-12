// internal/core/services/auth_service.go
package authservice

import (
	"GoNext/base/internal/core/ports"
)

type AuthService struct {
	userRepo  ports.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo ports.UserRepository, jwtSecret string) ports.AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}
