package workoutexercisesetrepository

import (
	"GoNext/base/internal/core/domain"
	"context"
)

func (r *WorkoutExerciseSetRepository) CreateWorkoutExerciseSet(workoutExerciseId, weight, reps int) (*domain.WorkoutExerciseSet, error) {
	ctx := context.Background()
	wrk, err := r.client.WorkoutExerciseSet.Create().
		SetWorkoutExerciseID(workoutExerciseId).
		SetWeight(weight).
		SetReps(reps).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return wrk.ToDomain(), nil
}
