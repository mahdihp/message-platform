package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int("tenant_id").Optional(),
		field.String("first_name").MaxLen(50).MinLen(3),
		field.String("last_name").MaxLen(50),
		field.String("mobile").MaxLen(11),
		field.String("email").MaxLen(100),
		field.String("password_hash"),
		field.Bool("status"),
		field.Bool("status2"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("users").Unique().Required(),
	}
}
