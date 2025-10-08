package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/ent/workoutexercise"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"context"

	"github.com/google/uuid"
)

type WorkoutExerciseRepository struct {
	client *ent.Client
}

func NewWorkoutExerciseRepository(client *ent.Client) ports.WorkoutExerciseRepository {
	return &WorkoutExerciseRepository{
		client: client,
	}
}

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

func (r *WorkoutExerciseRepository) GetWorkoutExerciseByWorkoutId(id uuid.UUID) ([]*domain.WorkoutExercise, error) {
	ctx := context.Background()
	rtn, err := r.client.WorkoutExercise.Query().
		Where(workoutexercise.WorkoutID(id)).
		WithExercise().
		All(ctx)
	if err != nil {
		return nil, err
	}
	var we ent.WorkoutExercises = rtn
	return we.ToDomain(), nil
}


