// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/B6113759/app/ent/deceasedreceive"
	"github.com/B6113759/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CoolroomUpdate is the builder for updating Coolroom entities.
type CoolroomUpdate struct {
	config
	hooks      []Hook
	mutation   *CoolroomMutation
	predicates []predicate.Coolroom
}

// Where adds a new predicate for the builder.
func (cu *CoolroomUpdate) Where(ps ...predicate.Coolroom) *CoolroomUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetCoolroomName sets the coolroom_name field.
func (cu *CoolroomUpdate) SetCoolroomName(s string) *CoolroomUpdate {
	cu.mutation.SetCoolroomName(s)
	return cu
}

// SetCoolroomCapacity sets the coolroom_capacity field.
func (cu *CoolroomUpdate) SetCoolroomCapacity(i int) *CoolroomUpdate {
	cu.mutation.ResetCoolroomCapacity()
	cu.mutation.SetCoolroomCapacity(i)
	return cu
}

// AddCoolroomCapacity adds i to coolroom_capacity.
func (cu *CoolroomUpdate) AddCoolroomCapacity(i int) *CoolroomUpdate {
	cu.mutation.AddCoolroomCapacity(i)
	return cu
}

// AddDeceasedreceifeIDs adds the deceasedreceives edge to DeceasedReceive by ids.
func (cu *CoolroomUpdate) AddDeceasedreceifeIDs(ids ...int) *CoolroomUpdate {
	cu.mutation.AddDeceasedreceifeIDs(ids...)
	return cu
}

// AddDeceasedreceives adds the deceasedreceives edges to DeceasedReceive.
func (cu *CoolroomUpdate) AddDeceasedreceives(d ...*DeceasedReceive) *CoolroomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.AddDeceasedreceifeIDs(ids...)
}

// SetCoolroomtypeID sets the coolroomtype edge to CoolroomType by id.
func (cu *CoolroomUpdate) SetCoolroomtypeID(id int) *CoolroomUpdate {
	cu.mutation.SetCoolroomtypeID(id)
	return cu
}

// SetNillableCoolroomtypeID sets the coolroomtype edge to CoolroomType by id if the given value is not nil.
func (cu *CoolroomUpdate) SetNillableCoolroomtypeID(id *int) *CoolroomUpdate {
	if id != nil {
		cu = cu.SetCoolroomtypeID(*id)
	}
	return cu
}

// SetCoolroomtype sets the coolroomtype edge to CoolroomType.
func (cu *CoolroomUpdate) SetCoolroomtype(c *CoolroomType) *CoolroomUpdate {
	return cu.SetCoolroomtypeID(c.ID)
}

// Mutation returns the CoolroomMutation object of the builder.
func (cu *CoolroomUpdate) Mutation() *CoolroomMutation {
	return cu.mutation
}

// RemoveDeceasedreceifeIDs removes the deceasedreceives edge to DeceasedReceive by ids.
func (cu *CoolroomUpdate) RemoveDeceasedreceifeIDs(ids ...int) *CoolroomUpdate {
	cu.mutation.RemoveDeceasedreceifeIDs(ids...)
	return cu
}

// RemoveDeceasedreceives removes deceasedreceives edges to DeceasedReceive.
func (cu *CoolroomUpdate) RemoveDeceasedreceives(d ...*DeceasedReceive) *CoolroomUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cu.RemoveDeceasedreceifeIDs(ids...)
}

// ClearCoolroomtype clears the coolroomtype edge to CoolroomType.
func (cu *CoolroomUpdate) ClearCoolroomtype() *CoolroomUpdate {
	cu.mutation.ClearCoolroomtype()
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CoolroomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cu.mutation.CoolroomName(); ok {
		if err := coolroom.CoolroomNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "coolroom_name", err: fmt.Errorf("ent: validator failed for field \"coolroom_name\": %w", err)}
		}
	}
	if v, ok := cu.mutation.CoolroomCapacity(); ok {
		if err := coolroom.CoolroomCapacityValidator(v); err != nil {
			return 0, &ValidationError{Name: "coolroom_capacity", err: fmt.Errorf("ent: validator failed for field \"coolroom_capacity\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoolroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CoolroomUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CoolroomUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CoolroomUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CoolroomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coolroom.Table,
			Columns: coolroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coolroom.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CoolroomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coolroom.FieldCoolroomName,
		})
	}
	if value, ok := cu.mutation.CoolroomCapacity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coolroom.FieldCoolroomCapacity,
		})
	}
	if value, ok := cu.mutation.AddedCoolroomCapacity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coolroom.FieldCoolroomCapacity,
		})
	}
	if nodes := cu.mutation.RemovedDeceasedreceivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coolroom.DeceasedreceivesTable,
			Columns: []string{coolroom.DeceasedreceivesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deceasedreceive.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.DeceasedreceivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coolroom.DeceasedreceivesTable,
			Columns: []string{coolroom.DeceasedreceivesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deceasedreceive.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CoolroomtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coolroom.CoolroomtypeTable,
			Columns: []string{coolroom.CoolroomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coolroomtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CoolroomtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coolroom.CoolroomtypeTable,
			Columns: []string{coolroom.CoolroomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coolroomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coolroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CoolroomUpdateOne is the builder for updating a single Coolroom entity.
type CoolroomUpdateOne struct {
	config
	hooks    []Hook
	mutation *CoolroomMutation
}

// SetCoolroomName sets the coolroom_name field.
func (cuo *CoolroomUpdateOne) SetCoolroomName(s string) *CoolroomUpdateOne {
	cuo.mutation.SetCoolroomName(s)
	return cuo
}

// SetCoolroomCapacity sets the coolroom_capacity field.
func (cuo *CoolroomUpdateOne) SetCoolroomCapacity(i int) *CoolroomUpdateOne {
	cuo.mutation.ResetCoolroomCapacity()
	cuo.mutation.SetCoolroomCapacity(i)
	return cuo
}

// AddCoolroomCapacity adds i to coolroom_capacity.
func (cuo *CoolroomUpdateOne) AddCoolroomCapacity(i int) *CoolroomUpdateOne {
	cuo.mutation.AddCoolroomCapacity(i)
	return cuo
}

// AddDeceasedreceifeIDs adds the deceasedreceives edge to DeceasedReceive by ids.
func (cuo *CoolroomUpdateOne) AddDeceasedreceifeIDs(ids ...int) *CoolroomUpdateOne {
	cuo.mutation.AddDeceasedreceifeIDs(ids...)
	return cuo
}

// AddDeceasedreceives adds the deceasedreceives edges to DeceasedReceive.
func (cuo *CoolroomUpdateOne) AddDeceasedreceives(d ...*DeceasedReceive) *CoolroomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.AddDeceasedreceifeIDs(ids...)
}

// SetCoolroomtypeID sets the coolroomtype edge to CoolroomType by id.
func (cuo *CoolroomUpdateOne) SetCoolroomtypeID(id int) *CoolroomUpdateOne {
	cuo.mutation.SetCoolroomtypeID(id)
	return cuo
}

// SetNillableCoolroomtypeID sets the coolroomtype edge to CoolroomType by id if the given value is not nil.
func (cuo *CoolroomUpdateOne) SetNillableCoolroomtypeID(id *int) *CoolroomUpdateOne {
	if id != nil {
		cuo = cuo.SetCoolroomtypeID(*id)
	}
	return cuo
}

// SetCoolroomtype sets the coolroomtype edge to CoolroomType.
func (cuo *CoolroomUpdateOne) SetCoolroomtype(c *CoolroomType) *CoolroomUpdateOne {
	return cuo.SetCoolroomtypeID(c.ID)
}

// Mutation returns the CoolroomMutation object of the builder.
func (cuo *CoolroomUpdateOne) Mutation() *CoolroomMutation {
	return cuo.mutation
}

// RemoveDeceasedreceifeIDs removes the deceasedreceives edge to DeceasedReceive by ids.
func (cuo *CoolroomUpdateOne) RemoveDeceasedreceifeIDs(ids ...int) *CoolroomUpdateOne {
	cuo.mutation.RemoveDeceasedreceifeIDs(ids...)
	return cuo
}

// RemoveDeceasedreceives removes deceasedreceives edges to DeceasedReceive.
func (cuo *CoolroomUpdateOne) RemoveDeceasedreceives(d ...*DeceasedReceive) *CoolroomUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cuo.RemoveDeceasedreceifeIDs(ids...)
}

// ClearCoolroomtype clears the coolroomtype edge to CoolroomType.
func (cuo *CoolroomUpdateOne) ClearCoolroomtype() *CoolroomUpdateOne {
	cuo.mutation.ClearCoolroomtype()
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CoolroomUpdateOne) Save(ctx context.Context) (*Coolroom, error) {
	if v, ok := cuo.mutation.CoolroomName(); ok {
		if err := coolroom.CoolroomNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "coolroom_name", err: fmt.Errorf("ent: validator failed for field \"coolroom_name\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.CoolroomCapacity(); ok {
		if err := coolroom.CoolroomCapacityValidator(v); err != nil {
			return nil, &ValidationError{Name: "coolroom_capacity", err: fmt.Errorf("ent: validator failed for field \"coolroom_capacity\": %w", err)}
		}
	}

	var (
		err  error
		node *Coolroom
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoolroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CoolroomUpdateOne) SaveX(ctx context.Context) *Coolroom {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CoolroomUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CoolroomUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CoolroomUpdateOne) sqlSave(ctx context.Context) (c *Coolroom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coolroom.Table,
			Columns: coolroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coolroom.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Coolroom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.CoolroomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coolroom.FieldCoolroomName,
		})
	}
	if value, ok := cuo.mutation.CoolroomCapacity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coolroom.FieldCoolroomCapacity,
		})
	}
	if value, ok := cuo.mutation.AddedCoolroomCapacity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: coolroom.FieldCoolroomCapacity,
		})
	}
	if nodes := cuo.mutation.RemovedDeceasedreceivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coolroom.DeceasedreceivesTable,
			Columns: []string{coolroom.DeceasedreceivesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deceasedreceive.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.DeceasedreceivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coolroom.DeceasedreceivesTable,
			Columns: []string{coolroom.DeceasedreceivesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deceasedreceive.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CoolroomtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coolroom.CoolroomtypeTable,
			Columns: []string{coolroom.CoolroomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coolroomtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CoolroomtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   coolroom.CoolroomtypeTable,
			Columns: []string{coolroom.CoolroomtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coolroomtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Coolroom{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coolroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
