package workoutexerciserepository

import (
	"GoNext/base/ent/user"
	"GoNext/base/ent/workout"
	"GoNext/base/ent/workoutexercise"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

func (r *WorkoutExerciseRepository) GetWorkoutExerciseByIdFromUser(weID int, userID uuid.UUID) (*domain.WorkoutExercise, error) {
	ctx := context.Background()
	we, err := r.client.WorkoutExercise.Query().
		Where(
			workoutexercise.And(
				workoutexercise.ID(weID),
				workoutexercise.HasWorkoutWith(
					workout.HasUserWith(user.ID(userID)),
				),
			),
		).
		WithExercise().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return we.ToDomain(), nil
}
