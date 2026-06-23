package schema

import (
	"message-platform/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.String("first_name").MaxLen(50).SchemaType(map[string]string{
			dialect.Postgres: "varchar(50)",
		}),
		field.String("last_name").MaxLen(50).SchemaType(map[string]string{
			dialect.Postgres: "varchar(70)",
		}),
		field.String("mobile").MaxLen(11).SchemaType(map[string]string{
			dialect.Postgres: "char(11)",
		}),
		field.String("email").MaxLen(100).SchemaType(map[string]string{
			dialect.Postgres: "varchar(100)",
		}),
		field.String("password_hash"),
		field.Bool("status"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.DeleteMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).Ref("users").Unique().Required(),
	}
}
