package workoutservice

import (
	"GoNext/base/internal/core/ports"
)

type WorkoutService struct {
	WorkoutRepository ports.WorkoutRepository
}

func NewWorkoutService(WorkoutRepo ports.WorkoutRepository) ports.WorkoutService {
	return &WorkoutService{
		WorkoutRepository: WorkoutRepo,
	}
}
