package workoutrepository

import (
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

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
