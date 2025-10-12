package exerciseservice

import "GoNext/base/internal/core/domain"

func (s *ExerciseService) GetExerciseById(id int) (*domain.Exercise, error) {
	return s.ExerciseRepository.GetExerciseById(id)
}
