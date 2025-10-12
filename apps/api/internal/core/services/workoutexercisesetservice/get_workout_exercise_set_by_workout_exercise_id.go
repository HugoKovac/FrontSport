package workoutexercisesetservice

import (
	"GoNext/base/internal/core/domain"
)

func (s *WorkoutExerciseSetService) GetWorkoutExerciseSetByWorkoutExerciseId(id int) ([]*domain.WorkoutExerciseSet, error) {
	return s.WorkoutExerciseSetRepository.GetWorkoutExerciseSetByWorkoutExerciseId(id)
}
