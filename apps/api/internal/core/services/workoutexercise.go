package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"

	"github.com/google/uuid"
)

type WorkoutExerciseService struct {
	WorkoutExerciseRepository ports.WorkoutExerciseRepository
}

func NewWorkoutExerciseService(workoutExerciseRepo ports.WorkoutExerciseRepository) ports.WorkoutExerciseService {
	return &WorkoutExerciseService{
		WorkoutExerciseRepository: workoutExerciseRepo,
	}
}

func (s *WorkoutExerciseService) CreateWorkoutExercise(exerciseId int, workoutId uuid.UUID) (*domain.WorkoutExercise, error) {
	return s.WorkoutExerciseRepository.CreateWorkoutExercise(exerciseId, workoutId)
}

func (s *WorkoutExerciseService) GetWorkoutExerciseByWorkoutId(id uuid.UUID) ([]*domain.WorkoutExercise, error) {
	return s.WorkoutExerciseRepository.GetWorkoutExerciseByWorkoutId(id)
}
