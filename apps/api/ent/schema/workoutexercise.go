package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkoutExercise holds the schema definition for the WorkoutExercise entity.
type WorkoutExercise struct {
	ent.Schema
}

// Fields of the WorkoutExercise.
func (WorkoutExercise) Fields() []ent.Field {
	return []ent.Field{
		field.Int("exercise_id"),
		field.UUID("workout_id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the WorkoutExercise.
func (WorkoutExercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workout", Workout.Type).
			Field("workout_id").
			Ref("workout_exercise").
			Unique().
			Required(),
		edge.From("exercise", Exercise.Type).
			Field("exercise_id").
			Ref("workout_exercise").
			Unique().
			Required(),
	}
}
