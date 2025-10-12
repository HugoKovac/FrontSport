package exerciserepository

import (
	"GoNext/base/ent"
	"GoNext/base/internal/core/domain"
	"context"
)

func (r *ExerciseRepository) GetExercises() ([]*domain.Exercise, error) {
	ctx := context.Background()
	rtn, err := r.client.Exercise.Query().All(ctx)
	var exercises ent.Exercises = rtn
	if err != nil {
		return nil, err
	}
	return exercises.ToDomain(), nil
}