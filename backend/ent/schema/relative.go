package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Relative holds the schema definition for the Relative entity.
type Relative struct {
	ent.Schema
}

// Fields of the Relative.
func (Relative) Fields() []ent.Field {
	return []ent.Field{
		field.String("relative_name").NotEmpty(),
	}
}

// Edges of the Relative.
func (Relative) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deceasedreceives", DeceasedReceive.Type).StorageKey(edge.Column("relative_id")),
	}
}
