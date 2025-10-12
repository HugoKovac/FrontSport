package authservice

import "GoNext/base/pkg/jwt"

func (s *AuthService) ValidateToken(tokenString string) (string, error) {
	return jwt.ValidateToken(tokenString, s.jwtSecret)
}
