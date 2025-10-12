package workoutrepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/ports"
)

type WorkoutRepository struct {
	client *ent.Client
}

func NewWorkoutRepository(client *ent.Client) ports.WorkoutRepository {
	return &WorkoutRepository{
		client: client,
	}
}
