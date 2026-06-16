package schema

import (
	"message-platform/configs"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.Int("parent_id").Optional(),
		field.String("name"),
		field.Enum("type").Values(string(configs.Platform)).Default(string(configs.Platform)),
		field.String("status"),
		field.String("brand_name").Optional().Unique(),
		field.String("domain").Optional().Unique(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Tenant.Type).
			From("parent").Unique().Field("parent_id"),
	}
}
