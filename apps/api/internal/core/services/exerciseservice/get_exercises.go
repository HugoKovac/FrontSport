package exerciseservice

import "GoNext/base/internal/core/domain"

func (s *ExerciseService) GetExercises() ([]*domain.Exercise, error) {
	return s.ExerciseRepository.GetExercises()
}
