package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/ent/exercise"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"context"
)

type ExerciseRepository struct {
	client *ent.Client
}

func NewExerciseRepository(client *ent.Client) ports.ExerciseRepository {
	return &ExerciseRepository{
		client: client,
	}
}

func (r *ExerciseRepository) GetExercises() ([]*domain.Exercise, error) {
	ctx := context.Background()
	rtn, err := r.client.Exercise.Query().All(ctx)
	var exercises ent.Exercises = rtn
	if err != nil {
		return nil, err
	}
	return exercises.ToDomain(), nil

}

func (r *ExerciseRepository) GetExerciseById(id int) (*domain.Exercise, error) {
	ctx := context.Background()
	exercise, err := r.client.Exercise.Query().Where(exercise.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return exercise.ToDomain(), nil
}
