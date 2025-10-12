package domain

import (
	"time"

	"github.com/google/uuid"
)

type Workout struct {
	Id               uuid.UUID             `json:"id"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	WorkoutExercises []*WorkoutExercise `json:"workout_exercises"`
	Active           bool               `json:"active"`
}
