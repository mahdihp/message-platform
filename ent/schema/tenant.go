package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("parent_id"),
		field.String("name"),
		field.String("type"),
		field.String("status"),
		field.String("brand_name").Optional().Unique(),
		field.String("domain").Optional().Unique(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
