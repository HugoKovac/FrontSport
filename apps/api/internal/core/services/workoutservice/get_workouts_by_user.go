package workoutservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutService) GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error) {
	return s.WorkoutRepository.GetWorkoutsByUser(userId)
}
