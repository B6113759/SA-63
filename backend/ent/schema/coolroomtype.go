package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CoolroomType holds the schema definition for the CoolroomType entity.
type CoolroomType struct {
	ent.Schema
}

// Fields of the CoolroomType.
func (CoolroomType) Fields() []ent.Field {
	return []ent.Field{
		field.String("coolroomtype_name").NotEmpty(),
	}
}

// Edges of the CoolroomType.
func (CoolroomType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coolrooms", Coolroom.Type).StorageKey(edge.Column("coolroomtype_id")),
	}
}
