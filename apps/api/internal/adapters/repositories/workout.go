package repositories

import (
	"GoNext/base/ent"
	"GoNext/base/ent/user"
	"GoNext/base/ent/workout"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"context"

	"github.com/google/uuid"
)

type WorkoutRepository struct {
	client *ent.Client
}

func NewWorkoutRepository(client *ent.Client) ports.WorkoutRepository {
	return &WorkoutRepository{
		client: client,
	}
}

func (r *WorkoutRepository) CreateWorkout(userId uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	wrk, err := r.client.Workout.
		Create().
		SetUserID(userId).
		SetActive(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return wrk.ToDomain(), nil
}

func (r *WorkoutRepository) GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error) {
	ctx := context.Background()
	rtn, err := r.client.Workout.Query().
		Where(
			workout.HasUserWith(user.ID(userId)),
		).
		Order(ent.Desc(workout.FieldCreatedAt)).
		WithWorkoutExercise(func(weq *ent.WorkoutExerciseQuery) { weq.WithExercise() }).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var workouts ent.Workouts = rtn
	return workouts.ToDomain(), nil
}

func (r *WorkoutRepository) GetWorkoutById(id uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	workout, err := r.client.Workout.Query().Where(workout.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return workout.ToDomain(), nil
}

func (r *WorkoutRepository) GetActiveWorkoutByUser(userId uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	workout, err := r.client.Workout.Query().
		Where(
			workout.And(
				workout.HasUserWith(
					user.ID(userId),
				),
				workout.ActiveEQ(true),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}
	return workout.ToDomain(), nil
}

func (r *WorkoutRepository) UpdateWorkoutToNotActive(id uuid.UUID) error {
	ctx := context.Background()
	return r.client.Workout.UpdateOneID(id).
		SetActive(false).
		Exec(ctx)
}
