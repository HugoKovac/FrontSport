package authservice

import (
	"GoNext/base/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Authenticate(username string, password string) (string, error) {
	// Get user by email
	user, err := s.userRepo.FindByEmail(username)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	return jwt.GenerateToken(user.Id.String(), s.jwtSecret, user.Role)
}
