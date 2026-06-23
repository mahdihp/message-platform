package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entmixin "entgo.io/ent/schema/mixin"
)

type DeleteMixin struct {
	entmixin.Schema
}

func (DeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Default(time.Now).
			Immutable(),
	}
}
