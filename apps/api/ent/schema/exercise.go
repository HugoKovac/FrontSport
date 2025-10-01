package schema

import (
	"GoNext/base/ent/schema/mixin/timestamps"


	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Exercise struct {
	ent.Schema
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("video_url").Optional().MaxLen(256),
		field.String("image_url").Optional().MaxLen(256),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Exercise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		timestamps.Mixin{},
	}
}
