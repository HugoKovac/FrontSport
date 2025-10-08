package domain

import (
	"time"
)

type Workout struct {
	Id               string             `json:"id"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	WorkoutExercises []*WorkoutExercise `json:"workout_exercises"`
	Active           bool               `json:"active"`
}
