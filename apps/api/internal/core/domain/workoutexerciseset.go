package domain

import (
	"github.com/google/uuid"
)

type WorkoutExerciseSet struct {
	Id                uuid.UUID        `json:"id"`
	WorkoutExerciseID uuid.UUID        `json:"workout_exercise_id"`
	WorkoutExercise   *WorkoutExercise `json:"workout_exercise,omitempty"`
	Weight            int              `json:"weight"`
	Reps              int              `json:"reps"`
}
