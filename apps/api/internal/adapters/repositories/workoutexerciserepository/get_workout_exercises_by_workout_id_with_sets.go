package workoutexerciserepository

import (
	"GoNext/base/ent"
	"GoNext/base/ent/workoutexercise"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutExerciseRepository) GetWorkoutExercisesByWorkoutIdWithExAndSets(id uuid.UUID) ([]*domain.WorkoutExercise, error) {
	ctx := context.Background()
	rtn, err := r.client.WorkoutExercise.Query().
		Where(workoutexercise.WorkoutID(id)).
		WithExercise().
		WithSets().
		All(ctx)
	if err != nil {
		return nil, err
	}
	var we ent.WorkoutExercises = rtn
	return we.ToDomain(), nil
}
