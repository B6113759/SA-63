package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Coolroom holds the schema definition for the Coolroom entity.
type Coolroom struct {
	ent.Schema
}

// Fields of the Coolroom.
func (Coolroom) Fields() []ent.Field {
	return []ent.Field{
		field.String("coolroom_name").NotEmpty(),
	}
}

// Edges of the Coolroom.
func (Coolroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deceasedreceives", DeceasedReceive.Type).StorageKey(edge.Column("coolroom_id")),
		edge.From("coolroomtype", CoolroomType.Type).
			Ref("coolrooms").
			Unique(),
	}
}
