package workoutexerciseservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutExerciseService) GetWorkoutExercisesByWorkoutIdWithExAndSets(id uuid.UUID) ([]*domain.WorkoutExercise, error) {
	return s.WorkoutExerciseRepository.GetWorkoutExercisesByWorkoutIdWithExAndSets(id)
}
