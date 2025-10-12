package workoutexercisesetservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutExerciseSetService) GetWorkoutExerciseSetByIdFromUser(id uuid.UUID, userID uuid.UUID) (*domain.WorkoutExerciseSet, error) {
	return s.WorkoutExerciseSetRepository.GetWorkoutExerciseSetByIdFromUser(id, userID)
}
