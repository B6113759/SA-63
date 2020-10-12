package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// DeceasedReceive holds the schema definition for the DeceasedReceive entity.
type DeceasedReceive struct {
	ent.Schema
}

// Fields of the DeceasedReceive.
func (DeceasedReceive) Fields() []ent.Field {
	return []ent.Field{
		field.Time("death_time"),
	}
}

// Edges of the DeceasedReceive.
func (DeceasedReceive) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", User.Type).
			Ref("deceasedreceives").
			Unique(),
		edge.From("relative", Relative.Type).
			Ref("deceasedreceives").
			Unique(),
		edge.From("coolroom", Coolroom.Type).
			Ref("deceasedreceives").
			Unique(),
		edge.From("patient", Patient.Type).
			Ref("deceasedreceives").
			Unique().
			Required(),
	}
}
