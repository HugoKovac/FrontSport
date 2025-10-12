package userservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *UserService) GetById(id uuid.UUID) (*domain.User, error) {
	user, err := s.UserRepository.FindById(id)
	return user, err
}
