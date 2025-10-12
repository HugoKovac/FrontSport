package workoutexerciserepository

import (
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutExerciseRepository) CreateWorkoutExercise(exerciseId int, workoutId uuid.UUID) (*domain.WorkoutExercise, error) {
	ctx := context.Background()
	wrk, err := r.client.WorkoutExercise.
		Create().
		SetExerciseID(exerciseId).
		SetWorkoutID(workoutId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return wrk.ToDomain(), nil
}
