package workoutservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutService) GetWorkoutById(id uuid.UUID) (*domain.Workout, error) {
	return s.WorkoutRepository.GetWorkoutById(id)
}
