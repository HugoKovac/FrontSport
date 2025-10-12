package workoutrepository

import (
	"GoNext/base/ent/workout"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutRepository) GetWorkoutById(id uuid.UUID) (*domain.Workout, error) {
	ctx := context.Background()
	workout, err := r.client.Workout.Query().Where(workout.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return workout.ToDomain(), nil
}
