package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
    return []ent.Field{
        field.String("Name"),
        field.UUID("user_id", uuid.UUID{}),
		// associate thread id
    }
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("workouts").Field("user_id").Unique().Required(),

        edge.To("exercises", Exercise.Type),
    }
}

func (Workout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
