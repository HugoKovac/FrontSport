package workoutexerciserepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/ports"
)

type WorkoutExerciseRepository struct {
	client *ent.Client
}

func NewWorkoutExerciseRepository(client *ent.Client) ports.WorkoutExerciseRepository {
	return &WorkoutExerciseRepository{
		client: client,
	}
}
