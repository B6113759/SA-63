// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Coolroom is the model entity for the Coolroom schema.
type Coolroom struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CoolroomName holds the value of the "coolroom_name" field.
	CoolroomName string `json:"coolroom_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CoolroomQuery when eager-loading is set.
	Edges           CoolroomEdges `json:"edges"`
	coolroomtype_id *int
}

// CoolroomEdges holds the relations/edges for other nodes in the graph.
type CoolroomEdges struct {
	// Deceasedreceives holds the value of the deceasedreceives edge.
	Deceasedreceives []*DeceasedReceive
	// Coolroomtype holds the value of the coolroomtype edge.
	Coolroomtype *CoolroomType
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DeceasedreceivesOrErr returns the Deceasedreceives value or an error if the edge
// was not loaded in eager-loading.
func (e CoolroomEdges) DeceasedreceivesOrErr() ([]*DeceasedReceive, error) {
	if e.loadedTypes[0] {
		return e.Deceasedreceives, nil
	}
	return nil, &NotLoadedError{edge: "deceasedreceives"}
}

// CoolroomtypeOrErr returns the Coolroomtype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CoolroomEdges) CoolroomtypeOrErr() (*CoolroomType, error) {
	if e.loadedTypes[1] {
		if e.Coolroomtype == nil {
			// The edge coolroomtype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: coolroomtype.Label}
		}
		return e.Coolroomtype, nil
	}
	return nil, &NotLoadedError{edge: "coolroomtype"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Coolroom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // coolroom_name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Coolroom) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // coolroomtype_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Coolroom fields.
func (c *Coolroom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(coolroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field coolroom_name", values[0])
	} else if value.Valid {
		c.CoolroomName = value.String
	}
	values = values[1:]
	if len(values) == len(coolroom.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field coolroomtype_id", value)
		} else if value.Valid {
			c.coolroomtype_id = new(int)
			*c.coolroomtype_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDeceasedreceives queries the deceasedreceives edge of the Coolroom.
func (c *Coolroom) QueryDeceasedreceives() *DeceasedReceiveQuery {
	return (&CoolroomClient{config: c.config}).QueryDeceasedreceives(c)
}

// QueryCoolroomtype queries the coolroomtype edge of the Coolroom.
func (c *Coolroom) QueryCoolroomtype() *CoolroomTypeQuery {
	return (&CoolroomClient{config: c.config}).QueryCoolroomtype(c)
}

// Update returns a builder for updating this Coolroom.
// Note that, you need to call Coolroom.Unwrap() before calling this method, if this Coolroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Coolroom) Update() *CoolroomUpdateOne {
	return (&CoolroomClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Coolroom) Unwrap() *Coolroom {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Coolroom is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Coolroom) String() string {
	var builder strings.Builder
	builder.WriteString("Coolroom(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", coolroom_name=")
	builder.WriteString(c.CoolroomName)
	builder.WriteByte(')')
	return builder.String()
}

// Coolrooms is a parsable slice of Coolroom.
type Coolrooms []*Coolroom

func (c Coolrooms) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
