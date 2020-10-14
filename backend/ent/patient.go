// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6113759/app/ent/deceasedreceive"
	"github.com/B6113759/app/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Patient is the model entity for the Patient schema.
type Patient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PatientName holds the value of the "patient_name" field.
	PatientName string `json:"patient_name,omitempty"`
	// PatientAge holds the value of the "patient_age" field.
	PatientAge int `json:"patient_age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientQuery when eager-loading is set.
	Edges PatientEdges `json:"edges"`
}

// PatientEdges holds the relations/edges for other nodes in the graph.
type PatientEdges struct {
	// Deceasedreceives holds the value of the deceasedreceives edge.
	Deceasedreceives *DeceasedReceive
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeceasedreceivesOrErr returns the Deceasedreceives value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientEdges) DeceasedreceivesOrErr() (*DeceasedReceive, error) {
	if e.loadedTypes[0] {
		if e.Deceasedreceives == nil {
			// The edge deceasedreceives was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: deceasedreceive.Label}
		}
		return e.Deceasedreceives, nil
	}
	return nil, &NotLoadedError{edge: "deceasedreceives"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Patient) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // patient_name
		&sql.NullInt64{},  // patient_age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Patient fields.
func (pa *Patient) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patient.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field patient_name", values[0])
	} else if value.Valid {
		pa.PatientName = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field patient_age", values[1])
	} else if value.Valid {
		pa.PatientAge = int(value.Int64)
	}
	return nil
}

// QueryDeceasedreceives queries the deceasedreceives edge of the Patient.
func (pa *Patient) QueryDeceasedreceives() *DeceasedReceiveQuery {
	return (&PatientClient{config: pa.config}).QueryDeceasedreceives(pa)
}

// Update returns a builder for updating this Patient.
// Note that, you need to call Patient.Unwrap() before calling this method, if this Patient
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Patient) Update() *PatientUpdateOne {
	return (&PatientClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Patient) Unwrap() *Patient {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Patient is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Patient) String() string {
	var builder strings.Builder
	builder.WriteString("Patient(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", patient_name=")
	builder.WriteString(pa.PatientName)
	builder.WriteString(", patient_age=")
	builder.WriteString(fmt.Sprintf("%v", pa.PatientAge))
	builder.WriteByte(')')
	return builder.String()
}

// Patients is a parsable slice of Patient.
type Patients []*Patient

func (pa Patients) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
