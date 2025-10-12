package workoutrepository

import (
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutRepository) UpdateWorkoutToNotActive(id uuid.UUID) error {
	ctx := context.Background()
	return r.client.Workout.UpdateOneID(id).
		SetActive(false).
		Exec(ctx)
}
