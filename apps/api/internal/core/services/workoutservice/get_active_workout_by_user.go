package workoutservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutService) GetActiveWorkoutByUser(userId uuid.UUID) (*domain.Workout, error) {
	return s.WorkoutRepository.GetActiveWorkoutByUser(userId)
}
