package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
)

type ExerciseService struct {
	ExerciseRepository ports.ExerciseRepository
}

func NewExerciseService(exerciseRepo ports.ExerciseRepository) ports.ExerciseService {
	return &ExerciseService{
		ExerciseRepository: exerciseRepo,
	}
}

func (s *ExerciseService) GetExercises() ([]*domain.Exercise, error) {
	return s.ExerciseRepository.GetExercises()
}
