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

func (r *WorkoutRepository) ToDomainWorkout(entWorkout *ent.Workout) *domain.Workout {
	return &domain.Workout{
		Id:        entWorkout.ID.String(),
		CreatedAt: entWorkout.CreatedAt,
		UpdatedAt: entWorkout.UpdatedAt,
	}
}

func (r *WorkoutRepository) ToDomainWorkouts(entWorkouts []*ent.Workout) (exs []*domain.Workout) {
	for _, v := range entWorkouts {
		exs = append(exs, r.ToDomainWorkout(v))
	}
	return
}

func (r *WorkoutRepository) CreateWorkout(userId uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	wrk, err := r.client.Workout.
		Create().
		SetUserID(userId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.ToDomainWorkout(wrk), nil
}

func (r *WorkoutRepository) GetWorkoutsByUser(userId uuid.UUID) ([]*domain.Workout, error) {
	ctx := context.Background()
	workouts, err := r.client.Workout.Query().
		Where(
			workout.HasUserWith(user.ID(userId)),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return r.ToDomainWorkouts(workouts), nil
}

func (r *WorkoutRepository) GetWorkoutById(id uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	workout, err := r.client.Workout.Query().Where(workout.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.ToDomainWorkout(workout), nil
}
