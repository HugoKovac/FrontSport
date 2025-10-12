package domain

import (
	"github.com/google/uuid"
)

type WorkoutExercise struct {
	Id         int                   `json:"id"`
	ExerciseID int                   `json:"exercise_id"`
	WorkoutID  uuid.UUID             `json:"workout_id"`
	Workout    *Workout              `json:"workout,omitempty"`
	Exercise   *Exercise             `json:"exercise,omitempty"`
	Sets       []*WorkoutExerciseSet `json:"sets,omitempty"`
}
