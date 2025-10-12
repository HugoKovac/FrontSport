package userservice

import (
	"GoNext/base/internal/core/domain"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Update(userId uuid.UUID, email string, oldPassword string, newPassword string) (domain.User, error) {
	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}
	if user == nil {
		return domain.User{}, errors.New("user not found")
	}
	if email != "" && email != user.Email {
		existingUser, err := s.UserRepository.FindByEmail(email)
		if existingUser != nil || err == nil {
			return domain.User{}, errors.New("user with this email already exists")
		}
		user.Email = email
	}
	if oldPassword != "" && newPassword != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
		if err != nil {
			return domain.User{}, errors.New("old password is incorrect")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, err
		}
		user.Password = string(hashedPassword)
	}
	user.UpdatedAt = time.Now()
	err = s.UserRepository.Update(user)
	if err != nil {
		return domain.User{}, errors.New("failed to update user")
	}
	return *user, nil
}
