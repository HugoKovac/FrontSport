package workoutrepository

import (
	"GoNext/base/ent/user"
	"GoNext/base/ent/workout"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

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
