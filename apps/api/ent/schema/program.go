package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Program holds the schema definition for the Program entity.
type Program struct {
	ent.Schema
}

// Fields of the Program.
func (Program) Fields() []ent.Field {
    return []ent.Field{
        field.String("Name"),
        field.UUID("user_id", uuid.UUID{}),
    }
}

// Edges of the Program.
func (Program) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("user", User.Type).Ref("programs").Field("user_id").Unique().Required(),

        edge.To("exercises", Exercise.Type),
    }
}

func (Program) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
