package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entmixin "entgo.io/ent/schema/mixin"
)

type AuditMixin struct {
	entmixin.Schema
}

func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_by").
			Default(time.Now).
			Immutable(),

		field.Time("updated_by").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
