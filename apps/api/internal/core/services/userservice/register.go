package userservice

import (
	"GoNext/base/internal/core/domain"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)


func (s *UserService) Register(user domain.User) (*domain.User, error) {
	existingUser, err := s.UserRepository.FindByEmail(user.Email)
	if existingUser != nil || err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return s.UserRepository.Create(user)
}
