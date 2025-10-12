package workoutexercisesetservice

import (
	"github.com/google/uuid"
)

func (s *WorkoutExerciseSetService) UpdateWorkoutExerciseSet(id uuid.UUID, weight int, reps int) error {
	return s.WorkoutExerciseSetRepository.UpdateWorkoutExerciseSet(id, weight, reps)
}
