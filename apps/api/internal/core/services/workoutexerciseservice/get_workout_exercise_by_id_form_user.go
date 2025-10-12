package workoutexerciseservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutExerciseService) GetWorkoutExerciseByIdFromUser(weID int, userID uuid.UUID) (*domain.WorkoutExercise, error) {
	return s.WorkoutExerciseRepository.GetWorkoutExerciseByIdFromUser(weID, userID)
}
