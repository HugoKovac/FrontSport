package workoutexercisesetservice

import (
	"GoNext/base/internal/core/domain"
)

func (s *WorkoutExerciseSetService) CreateWorkoutExerciseSet(workoutExerciseId, weight, reps int) (*domain.WorkoutExerciseSet, error) {
	return s.WorkoutExerciseSetRepository.CreateWorkoutExerciseSet(workoutExerciseId, weight, reps)
}

