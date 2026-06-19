package schema

import "entgo.io/ent"

// Scope holds the schema definition for the Scope entity.
type Scope struct {
	ent.Schema
}

// Fields of the Scope.
func (Scope) Fields() []ent.Field {
	return nil
}

// Edges of the Scope.
func (Scope) Edges() []ent.Edge {
	return nil
}
