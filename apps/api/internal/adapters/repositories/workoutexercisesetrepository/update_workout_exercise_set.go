package workoutexercisesetrepository

import (
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutExerciseSetRepository) UpdateWorkoutExerciseSet(id uuid.UUID, weight int, reps int) error {
	ctx := context.Background()
	return r.client.WorkoutExerciseSet.UpdateOneID(id).
		SetWeight(weight).
		SetReps(reps).
		Exec(ctx)
}
