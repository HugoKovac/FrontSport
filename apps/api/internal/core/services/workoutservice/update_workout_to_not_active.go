package workoutservice

import "github.com/google/uuid"

func (s *WorkoutService) UpdateWorkoutToNotActive(id uuid.UUID) error {
	return s.WorkoutRepository.UpdateWorkoutToNotActive(id)
}
