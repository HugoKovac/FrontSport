package dto

import "github.com/google/uuid"

type WorkoutQuery struct {
	WorkoutID uuid.UUID `params:"workout_id" validator:"required,uuid"`
}
