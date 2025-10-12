package workoutexerciseservice

import (
	"GoNext/base/internal/core/ports"
)

type WorkoutExerciseService struct {
	WorkoutExerciseRepository ports.WorkoutExerciseRepository
}

func NewWorkoutExerciseService(workoutExerciseRepo ports.WorkoutExerciseRepository) ports.WorkoutExerciseService {
	return &WorkoutExerciseService{
		WorkoutExerciseRepository: workoutExerciseRepo,
	}
}
