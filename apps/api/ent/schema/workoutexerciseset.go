package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Set holds the schema definition for the Set entity.
type WorkoutExerciseSet struct {
	ent.Schema
}

// Fields of the Set.
func (WorkoutExerciseSet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("weight").Default(0),
		field.Int("reps").Default(0),
	}
}

// Edges of the Set.
func (WorkoutExerciseSet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workout_exercise", WorkoutExercise.Type).
			Ref("sets").
			Unique(),
	}
}

func (WorkoutExerciseSet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
