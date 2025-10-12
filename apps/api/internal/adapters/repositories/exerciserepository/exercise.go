package exerciserepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/ports"
)

type ExerciseRepository struct {
	client *ent.Client
}

func NewExerciseRepository(client *ent.Client) ports.ExerciseRepository {
	return &ExerciseRepository{
		client: client,
	}
}
