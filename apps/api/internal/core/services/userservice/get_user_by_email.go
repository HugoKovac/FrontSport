package userservice

import "GoNext/base/internal/core/domain"

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	user, err := s.UserRepository.FindByEmail(email)
	return user, err
}
