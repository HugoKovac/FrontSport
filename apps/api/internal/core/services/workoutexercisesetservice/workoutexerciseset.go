package workoutexercisesetservice

import (
	"GoNext/base/internal/core/ports"
)

type WorkoutExerciseSetService struct {
	WorkoutExerciseSetRepository ports.WorkoutExerciseSetRepository
}

func NewWorkoutExerciseSetService(workoutExerciseSetRepo ports.WorkoutExerciseSetRepository) ports.WorkoutExerciseSetService {
	return &WorkoutExerciseSetService{
		WorkoutExerciseSetRepository: workoutExerciseSetRepo,
	}
}
