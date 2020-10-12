// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/deceasedreceive"
	"github.com/B6113759/app/ent/patient"
	"github.com/B6113759/app/ent/predicate"
	"github.com/B6113759/app/ent/relative"
	"github.com/B6113759/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeceasedReceiveQuery is the builder for querying DeceasedReceive entities.
type DeceasedReceiveQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.DeceasedReceive
	// eager-loading edges.
	withDoctor   *UserQuery
	withRelative *RelativeQuery
	withCoolroom *CoolroomQuery
	withPatient  *PatientQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (drq *DeceasedReceiveQuery) Where(ps ...predicate.DeceasedReceive) *DeceasedReceiveQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit adds a limit step to the query.
func (drq *DeceasedReceiveQuery) Limit(limit int) *DeceasedReceiveQuery {
	drq.limit = &limit
	return drq
}

// Offset adds an offset step to the query.
func (drq *DeceasedReceiveQuery) Offset(offset int) *DeceasedReceiveQuery {
	drq.offset = &offset
	return drq
}

// Order adds an order step to the query.
func (drq *DeceasedReceiveQuery) Order(o ...OrderFunc) *DeceasedReceiveQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// QueryDoctor chains the current query on the doctor edge.
func (drq *DeceasedReceiveQuery) QueryDoctor() *UserQuery {
	query := &UserQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, drq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.DoctorTable, deceasedreceive.DoctorColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelative chains the current query on the relative edge.
func (drq *DeceasedReceiveQuery) QueryRelative() *RelativeQuery {
	query := &RelativeQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, drq.sqlQuery()),
			sqlgraph.To(relative.Table, relative.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.RelativeTable, deceasedreceive.RelativeColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCoolroom chains the current query on the coolroom edge.
func (drq *DeceasedReceiveQuery) QueryCoolroom() *CoolroomQuery {
	query := &CoolroomQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, drq.sqlQuery()),
			sqlgraph.To(coolroom.Table, coolroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.CoolroomTable, deceasedreceive.CoolroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the patient edge.
func (drq *DeceasedReceiveQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, drq.sqlQuery()),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.PatientTable, deceasedreceive.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeceasedReceive entity in the query. Returns *NotFoundError when no deceasedreceive was found.
func (drq *DeceasedReceiveQuery) First(ctx context.Context) (*DeceasedReceive, error) {
	drs, err := drq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(drs) == 0 {
		return nil, &NotFoundError{deceasedreceive.Label}
	}
	return drs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) FirstX(ctx context.Context) *DeceasedReceive {
	dr, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return dr
}

// FirstID returns the first DeceasedReceive id in the query. Returns *NotFoundError when no id was found.
func (drq *DeceasedReceiveQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deceasedreceive.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) FirstXID(ctx context.Context) int {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only DeceasedReceive entity in the query, returns an error if not exactly one entity was returned.
func (drq *DeceasedReceiveQuery) Only(ctx context.Context) (*DeceasedReceive, error) {
	drs, err := drq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(drs) {
	case 1:
		return drs[0], nil
	case 0:
		return nil, &NotFoundError{deceasedreceive.Label}
	default:
		return nil, &NotSingularError{deceasedreceive.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) OnlyX(ctx context.Context) *DeceasedReceive {
	dr, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return dr
}

// OnlyID returns the only DeceasedReceive id in the query, returns an error if not exactly one id was returned.
func (drq *DeceasedReceiveQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = &NotSingularError{deceasedreceive.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) OnlyIDX(ctx context.Context) int {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeceasedReceives.
func (drq *DeceasedReceiveQuery) All(ctx context.Context) ([]*DeceasedReceive, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return drq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) AllX(ctx context.Context) []*DeceasedReceive {
	drs, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return drs
}

// IDs executes the query and returns a list of DeceasedReceive ids.
func (drq *DeceasedReceiveQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := drq.Select(deceasedreceive.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) IDsX(ctx context.Context) []int {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DeceasedReceiveQuery) Count(ctx context.Context) (int, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return drq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DeceasedReceiveQuery) Exist(ctx context.Context) (bool, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return drq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DeceasedReceiveQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DeceasedReceiveQuery) Clone() *DeceasedReceiveQuery {
	return &DeceasedReceiveQuery{
		config:     drq.config,
		limit:      drq.limit,
		offset:     drq.offset,
		order:      append([]OrderFunc{}, drq.order...),
		unique:     append([]string{}, drq.unique...),
		predicates: append([]predicate.DeceasedReceive{}, drq.predicates...),
		// clone intermediate query.
		sql:  drq.sql.Clone(),
		path: drq.path,
	}
}

//  WithDoctor tells the query-builder to eager-loads the nodes that are connected to
// the "doctor" edge. The optional arguments used to configure the query builder of the edge.
func (drq *DeceasedReceiveQuery) WithDoctor(opts ...func(*UserQuery)) *DeceasedReceiveQuery {
	query := &UserQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withDoctor = query
	return drq
}

//  WithRelative tells the query-builder to eager-loads the nodes that are connected to
// the "relative" edge. The optional arguments used to configure the query builder of the edge.
func (drq *DeceasedReceiveQuery) WithRelative(opts ...func(*RelativeQuery)) *DeceasedReceiveQuery {
	query := &RelativeQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withRelative = query
	return drq
}

//  WithCoolroom tells the query-builder to eager-loads the nodes that are connected to
// the "coolroom" edge. The optional arguments used to configure the query builder of the edge.
func (drq *DeceasedReceiveQuery) WithCoolroom(opts ...func(*CoolroomQuery)) *DeceasedReceiveQuery {
	query := &CoolroomQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withCoolroom = query
	return drq
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (drq *DeceasedReceiveQuery) WithPatient(opts ...func(*PatientQuery)) *DeceasedReceiveQuery {
	query := &PatientQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withPatient = query
	return drq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeathTime time.Time `json:"death_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeceasedReceive.Query().
//		GroupBy(deceasedreceive.FieldDeathTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (drq *DeceasedReceiveQuery) GroupBy(field string, fields ...string) *DeceasedReceiveGroupBy {
	group := &DeceasedReceiveGroupBy{config: drq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		DeathTime time.Time `json:"death_time,omitempty"`
//	}
//
//	client.DeceasedReceive.Query().
//		Select(deceasedreceive.FieldDeathTime).
//		Scan(ctx, &v)
//
func (drq *DeceasedReceiveQuery) Select(field string, fields ...string) *DeceasedReceiveSelect {
	selector := &DeceasedReceiveSelect{config: drq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(), nil
	}
	return selector
}

func (drq *DeceasedReceiveQuery) prepareQuery(ctx context.Context) error {
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DeceasedReceiveQuery) sqlAll(ctx context.Context) ([]*DeceasedReceive, error) {
	var (
		nodes       = []*DeceasedReceive{}
		withFKs     = drq.withFKs
		_spec       = drq.querySpec()
		loadedTypes = [4]bool{
			drq.withDoctor != nil,
			drq.withRelative != nil,
			drq.withCoolroom != nil,
			drq.withPatient != nil,
		}
	)
	if drq.withDoctor != nil || drq.withRelative != nil || drq.withCoolroom != nil || drq.withPatient != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deceasedreceive.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &DeceasedReceive{config: drq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := drq.withDoctor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeceasedReceive)
		for i := range nodes {
			if fk := nodes[i].doctor_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Doctor = n
			}
		}
	}

	if query := drq.withRelative; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeceasedReceive)
		for i := range nodes {
			if fk := nodes[i].relative_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(relative.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "relative_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Relative = n
			}
		}
	}

	if query := drq.withCoolroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeceasedReceive)
		for i := range nodes {
			if fk := nodes[i].coolroom_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(coolroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coolroom_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Coolroom = n
			}
		}
	}

	if query := drq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DeceasedReceive)
		for i := range nodes {
			if fk := nodes[i].patient_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	return nodes, nil
}

func (drq *DeceasedReceiveQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DeceasedReceiveQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := drq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (drq *DeceasedReceiveQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deceasedreceive.Table,
			Columns: deceasedreceive.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deceasedreceive.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (drq *DeceasedReceiveQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(deceasedreceive.Table)
	selector := builder.Select(t1.Columns(deceasedreceive.Columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(deceasedreceive.Columns...)...)
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector)
	}
	if offset := drq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeceasedReceiveGroupBy is the builder for group-by DeceasedReceive entities.
type DeceasedReceiveGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DeceasedReceiveGroupBy) Aggregate(fns ...AggregateFunc) *DeceasedReceiveGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the group-by query and scan the result into the given value.
func (drgb *DeceasedReceiveGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := drgb.path(ctx)
	if err != nil {
		return err
	}
	drgb.sql = query
	return drgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := drgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) StringsX(ctx context.Context) []string {
	v, err := drgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) StringX(ctx context.Context) string {
	v, err := drgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) IntsX(ctx context.Context) []int {
	v, err := drgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) IntX(ctx context.Context) int {
	v, err := drgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := drgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) Float64X(ctx context.Context) float64 {
	v, err := drgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := drgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (drgb *DeceasedReceiveGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drgb *DeceasedReceiveGroupBy) BoolX(ctx context.Context) bool {
	v, err := drgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drgb *DeceasedReceiveGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := drgb.sqlQuery().Query()
	if err := drgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drgb *DeceasedReceiveGroupBy) sqlQuery() *sql.Selector {
	selector := drgb.sql
	columns := make([]string, 0, len(drgb.fields)+len(drgb.fns))
	columns = append(columns, drgb.fields...)
	for _, fn := range drgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(drgb.fields...)
}

// DeceasedReceiveSelect is the builder for select fields of DeceasedReceive entities.
type DeceasedReceiveSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (drs *DeceasedReceiveSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := drs.path(ctx)
	if err != nil {
		return err
	}
	drs.sql = query
	return drs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) ScanX(ctx context.Context, v interface{}) {
	if err := drs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Strings(ctx context.Context) ([]string, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) StringsX(ctx context.Context) []string {
	v, err := drs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) StringX(ctx context.Context) string {
	v, err := drs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Ints(ctx context.Context) ([]int, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) IntsX(ctx context.Context) []int {
	v, err := drs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) IntX(ctx context.Context) int {
	v, err := drs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) Float64sX(ctx context.Context) []float64 {
	v, err := drs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) Float64X(ctx context.Context) float64 {
	v, err := drs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DeceasedReceiveSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) BoolsX(ctx context.Context) []bool {
	v, err := drs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (drs *DeceasedReceiveSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{deceasedreceive.Label}
	default:
		err = fmt.Errorf("ent: DeceasedReceiveSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drs *DeceasedReceiveSelect) BoolX(ctx context.Context) bool {
	v, err := drs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drs *DeceasedReceiveSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := drs.sqlQuery().Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drs *DeceasedReceiveSelect) sqlQuery() sql.Querier {
	selector := drs.sql
	selector.Select(selector.Columns(drs.fields...)...)
	return selector
}