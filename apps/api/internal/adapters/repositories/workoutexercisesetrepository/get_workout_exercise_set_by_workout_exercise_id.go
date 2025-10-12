package workoutexercisesetrepository

import (
	"GoNext/base/ent"
	"GoNext/base/ent/workoutexercise"
	"GoNext/base/ent/workoutexerciseset"
	"GoNext/base/internal/core/domain"
	"context"
)

func (r *WorkoutExerciseSetRepository) GetWorkoutExerciseSetByWorkoutExerciseId(id int) ([]*domain.WorkoutExerciseSet, error) {
	ctx := context.Background()
	rtn, err := r.client.WorkoutExerciseSet.Query().
		Where(workoutexerciseset.HasWorkoutExerciseWith(workoutexercise.ID(id))).
		WithWorkoutExercise().
		All(ctx)
	if err != nil {
		return nil, err
	}
	var we ent.WorkoutExerciseSets = rtn
	return we.ToDomain(), nil
}
