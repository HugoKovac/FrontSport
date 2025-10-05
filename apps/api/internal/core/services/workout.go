package services

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"

	"github.com/google/uuid"
)

type WorkoutService struct {
	WorkoutRepository ports.WorkoutRepository
}

func NewWorkoutService(WorkoutRepo ports.WorkoutRepository) ports.WorkoutService {
	return &WorkoutService{
		WorkoutRepository: WorkoutRepo,
	}
}

func (s *WorkoutService) CreateWorkout(userId uuid.UUID) (*domain.Workout, error) {
	return s.WorkoutRepository.CreateWorkout(userId)
}

func (s *WorkoutService) GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error) {
	return s.WorkoutRepository.GetWorkoutsByUser(userId)
}

func (s *WorkoutService) GetWorkoutById(id uuid.UUID) (*domain.Workout, error) {
	return s.WorkoutRepository.GetWorkoutById(id)
}
