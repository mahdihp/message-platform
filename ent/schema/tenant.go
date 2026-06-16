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
		field.Strings("name"),
		field.Strings("type"),
		field.Strings("status"),
		field.Strings("brand_name"),
		field.Strings("domain"),
		field.Strings("domain3"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}
