package ports

import (
	"GoNext/base/internal/core/domain"
)

type UserRepository interface {
	Create(user domain.User) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
}

type ExerciseRepository interface {
	GetExercises() ([]*domain.Exercise, error)
	GetExerciseById(id int) (*domain.Exercise, error)
}
