package workoutexercisesetrepository

import (
	"GoNext/base/ent/user"
	"GoNext/base/ent/workout"
	"GoNext/base/ent/workoutexercise"
	"GoNext/base/ent/workoutexerciseset"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutExerciseSetRepository) GetWorkoutExerciseSetByIdFromUser(id uuid.UUID, userID uuid.UUID) (*domain.WorkoutExerciseSet, error) {
	ctx := context.Background()
	we, err := r.client.WorkoutExerciseSet.Query().
		Where(workoutexerciseset.And(
			workoutexerciseset.ID(id),
			workoutexerciseset.HasWorkoutExerciseWith(
				workoutexercise.HasWorkoutWith(
					workout.HasUserWith(user.ID(userID)),
				),
			),
		)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return we.ToDomain(), nil
}
