package workoutexercisesetrepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/ports"
)

type WorkoutExerciseSetRepository struct {
	client *ent.Client
}

func NewWorkoutExerciseSetRepository(client *ent.Client) ports.WorkoutExerciseSetRepository {
	return &WorkoutExerciseSetRepository{
		client: client,
	}
}
