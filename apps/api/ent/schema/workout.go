package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Optional(),
		field.Bool("active").Default(false),
		field.UUID("user_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workout_exercise", WorkoutExercise.Type),
		edge.From("user", User.Type).
			Ref("workouts").
			Field("user_id").
			Unique(),
	}
}

func (Workout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}

func (Workout) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "active").
			Unique().
			Annotations(
				entsql.IndexWhere("active = true"),
			),
	}
}
