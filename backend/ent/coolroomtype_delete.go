// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/B6113759/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CoolroomTypeDelete is the builder for deleting a CoolroomType entity.
type CoolroomTypeDelete struct {
	config
	hooks      []Hook
	mutation   *CoolroomTypeMutation
	predicates []predicate.CoolroomType
}

// Where adds a new predicate to the delete builder.
func (ctd *CoolroomTypeDelete) Where(ps ...predicate.CoolroomType) *CoolroomTypeDelete {
	ctd.predicates = append(ctd.predicates, ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *CoolroomTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctd.hooks) == 0 {
		affected, err = ctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoolroomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctd.mutation = mutation
			affected, err = ctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctd.hooks) - 1; i >= 0; i-- {
			mut = ctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *CoolroomTypeDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *CoolroomTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coolroomtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coolroomtype.FieldID,
			},
		},
	}
	if ps := ctd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
}

// CoolroomTypeDeleteOne is the builder for deleting a single CoolroomType entity.
type CoolroomTypeDeleteOne struct {
	ctd *CoolroomTypeDelete
}

// Exec executes the deletion query.
func (ctdo *CoolroomTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coolroomtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *CoolroomTypeDeleteOne) ExecX(ctx context.Context) {
	ctdo.ctd.ExecX(ctx)
}
