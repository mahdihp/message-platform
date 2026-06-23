package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("code"),
		field.String("name"),
		field.String("module"),
		field.String("description").MaxLen(100),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
