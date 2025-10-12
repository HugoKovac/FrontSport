package workoutexerciseservice

import (
	"GoNext/base/internal/core/domain"

	"github.com/google/uuid"
)

func (s *WorkoutExerciseService) CreateWorkoutExercise(exerciseId int, workoutId uuid.UUID) (*domain.WorkoutExercise, error) {
	return s.WorkoutExerciseRepository.CreateWorkoutExercise(exerciseId, workoutId)
}
