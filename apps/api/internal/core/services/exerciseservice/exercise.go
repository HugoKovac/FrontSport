package exerciseservice

import (
	"GoNext/base/internal/core/ports"
)

type ExerciseService struct {
	ExerciseRepository ports.ExerciseRepository
}

func NewExerciseService(exerciseRepo ports.ExerciseRepository) ports.ExerciseService {
	return &ExerciseService{
		ExerciseRepository: exerciseRepo,
	}
}
