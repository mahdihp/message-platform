package schema

import (
	"message-platform/internal/shared"

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
		field.String("name").MaxLen(50).Unique(),
		field.Enum("type").Values(string(shared.Platform), string(shared.Reseller), string(shared.Customer)).Default(string(shared.Platform)),
		field.Enum("status").Values(string(shared.Draft), string(shared.Pending), string(shared.Verified), string(shared.Suspended)).Default(string(shared.Draft)),
		field.String("brand_name").MaxLen(50).Optional().Unique(),
		field.String("domain").MaxLen(200).Optional().Unique(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Tenant.Type).From("parent").Unique().Field("parent_id"),
		edge.To("users", User.Type),
	}
}
