package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("url").MaxLen(2048),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("programs", Program.Type).Ref("exercises").Unique(),
        edge.From("workouts", Workout.Type).Ref("exercises").Unique(),
    }
}

func (Exercise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
