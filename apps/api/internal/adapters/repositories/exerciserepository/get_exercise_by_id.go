package exerciserepository

import (
	"GoNext/base/ent/exercise"
	"GoNext/base/internal/core/domain"
	"context"
)

func (r *ExerciseRepository) GetExerciseById(id int) (*domain.Exercise, error) {
	ctx := context.Background()
	exercise, err := r.client.Exercise.Query().Where(exercise.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return exercise.ToDomain(), nil
}