package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"
	"GoNext/base/internal/primitive/userprimitive"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("firstname").MaxLen(100).Optional(),
		field.String("lastname").MaxLen(100).Optional(),
		field.String("email").Unique(),
		field.String("password").NotEmpty(),
		field.Enum("role").GoType(userprimitive.Roles("")).Default(userprimitive.RoleUser.String()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("programs", Program.Type),
		edge.To("workouts", Workout.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
