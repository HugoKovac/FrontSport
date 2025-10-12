package workoutrepository

import (
	"GoNext/base/ent"
	"GoNext/base/ent/user"
	"GoNext/base/ent/workout"
	"GoNext/base/internal/core/domain"
	"context"

	"github.com/google/uuid"
)

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
