// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/deceasedreceive"
	"github.com/B6113759/app/ent/patient"
	"github.com/B6113759/app/ent/relative"
	"github.com/B6113759/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeceasedReceiveCreate is the builder for creating a DeceasedReceive entity.
type DeceasedReceiveCreate struct {
	config
	mutation *DeceasedReceiveMutation
	hooks    []Hook
}

// SetDeathTime sets the death_time field.
func (drc *DeceasedReceiveCreate) SetDeathTime(t time.Time) *DeceasedReceiveCreate {
	drc.mutation.SetDeathTime(t)
	return drc
}

// SetDoctorID sets the doctor edge to User by id.
func (drc *DeceasedReceiveCreate) SetDoctorID(id int) *DeceasedReceiveCreate {
	drc.mutation.SetDoctorID(id)
	return drc
}

// SetNillableDoctorID sets the doctor edge to User by id if the given value is not nil.
func (drc *DeceasedReceiveCreate) SetNillableDoctorID(id *int) *DeceasedReceiveCreate {
	if id != nil {
		drc = drc.SetDoctorID(*id)
	}
	return drc
}

// SetDoctor sets the doctor edge to User.
func (drc *DeceasedReceiveCreate) SetDoctor(u *User) *DeceasedReceiveCreate {
	return drc.SetDoctorID(u.ID)
}

// SetRelativeID sets the relative edge to Relative by id.
func (drc *DeceasedReceiveCreate) SetRelativeID(id int) *DeceasedReceiveCreate {
	drc.mutation.SetRelativeID(id)
	return drc
}

// SetNillableRelativeID sets the relative edge to Relative by id if the given value is not nil.
func (drc *DeceasedReceiveCreate) SetNillableRelativeID(id *int) *DeceasedReceiveCreate {
	if id != nil {
		drc = drc.SetRelativeID(*id)
	}
	return drc
}

// SetRelative sets the relative edge to Relative.
func (drc *DeceasedReceiveCreate) SetRelative(r *Relative) *DeceasedReceiveCreate {
	return drc.SetRelativeID(r.ID)
}

// SetCoolroomID sets the coolroom edge to Coolroom by id.
func (drc *DeceasedReceiveCreate) SetCoolroomID(id int) *DeceasedReceiveCreate {
	drc.mutation.SetCoolroomID(id)
	return drc
}

// SetNillableCoolroomID sets the coolroom edge to Coolroom by id if the given value is not nil.
func (drc *DeceasedReceiveCreate) SetNillableCoolroomID(id *int) *DeceasedReceiveCreate {
	if id != nil {
		drc = drc.SetCoolroomID(*id)
	}
	return drc
}

// SetCoolroom sets the coolroom edge to Coolroom.
func (drc *DeceasedReceiveCreate) SetCoolroom(c *Coolroom) *DeceasedReceiveCreate {
	return drc.SetCoolroomID(c.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (drc *DeceasedReceiveCreate) SetPatientID(id int) *DeceasedReceiveCreate {
	drc.mutation.SetPatientID(id)
	return drc
}

// SetPatient sets the patient edge to Patient.
func (drc *DeceasedReceiveCreate) SetPatient(p *Patient) *DeceasedReceiveCreate {
	return drc.SetPatientID(p.ID)
}

// Mutation returns the DeceasedReceiveMutation object of the builder.
func (drc *DeceasedReceiveCreate) Mutation() *DeceasedReceiveMutation {
	return drc.mutation
}

// Save creates the DeceasedReceive in the database.
func (drc *DeceasedReceiveCreate) Save(ctx context.Context) (*DeceasedReceive, error) {
	if _, ok := drc.mutation.DeathTime(); !ok {
		return nil, &ValidationError{Name: "death_time", err: errors.New("ent: missing required field \"death_time\"")}
	}
	if _, ok := drc.mutation.PatientID(); !ok {
		return nil, &ValidationError{Name: "patient", err: errors.New("ent: missing required edge \"patient\"")}
	}
	var (
		err  error
		node *DeceasedReceive
	)
	if len(drc.hooks) == 0 {
		node, err = drc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeceasedReceiveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			drc.mutation = mutation
			node, err = drc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(drc.hooks) - 1; i >= 0; i-- {
			mut = drc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, drc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (drc *DeceasedReceiveCreate) SaveX(ctx context.Context) *DeceasedReceive {
	v, err := drc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drc *DeceasedReceiveCreate) sqlSave(ctx context.Context) (*DeceasedReceive, error) {
	dr, _spec := drc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	dr.ID = int(id)
	return dr, nil
}

func (drc *DeceasedReceiveCreate) createSpec() (*DeceasedReceive, *sqlgraph.CreateSpec) {
	var (
		dr    = &DeceasedReceive{config: drc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deceasedreceive.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deceasedreceive.FieldID,
			},
		}
	)
	if value, ok := drc.mutation.DeathTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: deceasedreceive.FieldDeathTime,
		})
		dr.DeathTime = value
	}
	if nodes := drc.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deceasedreceive.DoctorTable,
			Columns: []string{deceasedreceive.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.RelativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deceasedreceive.RelativeTable,
			Columns: []string{deceasedreceive.RelativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.CoolroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deceasedreceive.CoolroomTable,
			Columns: []string{deceasedreceive.CoolroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coolroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   deceasedreceive.PatientTable,
			Columns: []string{deceasedreceive.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return dr, _spec
}
