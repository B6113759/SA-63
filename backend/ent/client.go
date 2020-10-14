// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/B6113759/app/ent/migrate"

	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/B6113759/app/ent/deceasedreceive"
	"github.com/B6113759/app/ent/patient"
	"github.com/B6113759/app/ent/relative"
	"github.com/B6113759/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Coolroom is the client for interacting with the Coolroom builders.
	Coolroom *CoolroomClient
	// CoolroomType is the client for interacting with the CoolroomType builders.
	CoolroomType *CoolroomTypeClient
	// DeceasedReceive is the client for interacting with the DeceasedReceive builders.
	DeceasedReceive *DeceasedReceiveClient
	// Patient is the client for interacting with the Patient builders.
	Patient *PatientClient
	// Relative is the client for interacting with the Relative builders.
	Relative *RelativeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Coolroom = NewCoolroomClient(c.config)
	c.CoolroomType = NewCoolroomTypeClient(c.config)
	c.DeceasedReceive = NewDeceasedReceiveClient(c.config)
	c.Patient = NewPatientClient(c.config)
	c.Relative = NewRelativeClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Coolroom:        NewCoolroomClient(cfg),
		CoolroomType:    NewCoolroomTypeClient(cfg),
		DeceasedReceive: NewDeceasedReceiveClient(cfg),
		Patient:         NewPatientClient(cfg),
		Relative:        NewRelativeClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:          cfg,
		Coolroom:        NewCoolroomClient(cfg),
		CoolroomType:    NewCoolroomTypeClient(cfg),
		DeceasedReceive: NewDeceasedReceiveClient(cfg),
		Patient:         NewPatientClient(cfg),
		Relative:        NewRelativeClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Coolroom.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Coolroom.Use(hooks...)
	c.CoolroomType.Use(hooks...)
	c.DeceasedReceive.Use(hooks...)
	c.Patient.Use(hooks...)
	c.Relative.Use(hooks...)
	c.User.Use(hooks...)
}

// CoolroomClient is a client for the Coolroom schema.
type CoolroomClient struct {
	config
}

// NewCoolroomClient returns a client for the Coolroom from the given config.
func NewCoolroomClient(c config) *CoolroomClient {
	return &CoolroomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coolroom.Hooks(f(g(h())))`.
func (c *CoolroomClient) Use(hooks ...Hook) {
	c.hooks.Coolroom = append(c.hooks.Coolroom, hooks...)
}

// Create returns a create builder for Coolroom.
func (c *CoolroomClient) Create() *CoolroomCreate {
	mutation := newCoolroomMutation(c.config, OpCreate)
	return &CoolroomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Coolroom.
func (c *CoolroomClient) Update() *CoolroomUpdate {
	mutation := newCoolroomMutation(c.config, OpUpdate)
	return &CoolroomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoolroomClient) UpdateOne(co *Coolroom) *CoolroomUpdateOne {
	mutation := newCoolroomMutation(c.config, OpUpdateOne, withCoolroom(co))
	return &CoolroomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoolroomClient) UpdateOneID(id int) *CoolroomUpdateOne {
	mutation := newCoolroomMutation(c.config, OpUpdateOne, withCoolroomID(id))
	return &CoolroomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Coolroom.
func (c *CoolroomClient) Delete() *CoolroomDelete {
	mutation := newCoolroomMutation(c.config, OpDelete)
	return &CoolroomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoolroomClient) DeleteOne(co *Coolroom) *CoolroomDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoolroomClient) DeleteOneID(id int) *CoolroomDeleteOne {
	builder := c.Delete().Where(coolroom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoolroomDeleteOne{builder}
}

// Create returns a query builder for Coolroom.
func (c *CoolroomClient) Query() *CoolroomQuery {
	return &CoolroomQuery{config: c.config}
}

// Get returns a Coolroom entity by its id.
func (c *CoolroomClient) Get(ctx context.Context, id int) (*Coolroom, error) {
	return c.Query().Where(coolroom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoolroomClient) GetX(ctx context.Context, id int) *Coolroom {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// QueryDeceasedreceives queries the deceasedreceives edge of a Coolroom.
func (c *CoolroomClient) QueryDeceasedreceives(co *Coolroom) *DeceasedReceiveQuery {
	query := &DeceasedReceiveQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coolroom.Table, coolroom.FieldID, id),
			sqlgraph.To(deceasedreceive.Table, deceasedreceive.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coolroom.DeceasedreceivesTable, coolroom.DeceasedreceivesColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCoolroomtype queries the coolroomtype edge of a Coolroom.
func (c *CoolroomClient) QueryCoolroomtype(co *Coolroom) *CoolroomTypeQuery {
	query := &CoolroomTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coolroom.Table, coolroom.FieldID, id),
			sqlgraph.To(coolroomtype.Table, coolroomtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, coolroom.CoolroomtypeTable, coolroom.CoolroomtypeColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CoolroomClient) Hooks() []Hook {
	return c.hooks.Coolroom
}

// CoolroomTypeClient is a client for the CoolroomType schema.
type CoolroomTypeClient struct {
	config
}

// NewCoolroomTypeClient returns a client for the CoolroomType from the given config.
func NewCoolroomTypeClient(c config) *CoolroomTypeClient {
	return &CoolroomTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coolroomtype.Hooks(f(g(h())))`.
func (c *CoolroomTypeClient) Use(hooks ...Hook) {
	c.hooks.CoolroomType = append(c.hooks.CoolroomType, hooks...)
}

// Create returns a create builder for CoolroomType.
func (c *CoolroomTypeClient) Create() *CoolroomTypeCreate {
	mutation := newCoolroomTypeMutation(c.config, OpCreate)
	return &CoolroomTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for CoolroomType.
func (c *CoolroomTypeClient) Update() *CoolroomTypeUpdate {
	mutation := newCoolroomTypeMutation(c.config, OpUpdate)
	return &CoolroomTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoolroomTypeClient) UpdateOne(ct *CoolroomType) *CoolroomTypeUpdateOne {
	mutation := newCoolroomTypeMutation(c.config, OpUpdateOne, withCoolroomType(ct))
	return &CoolroomTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoolroomTypeClient) UpdateOneID(id int) *CoolroomTypeUpdateOne {
	mutation := newCoolroomTypeMutation(c.config, OpUpdateOne, withCoolroomTypeID(id))
	return &CoolroomTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoolroomType.
func (c *CoolroomTypeClient) Delete() *CoolroomTypeDelete {
	mutation := newCoolroomTypeMutation(c.config, OpDelete)
	return &CoolroomTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoolroomTypeClient) DeleteOne(ct *CoolroomType) *CoolroomTypeDeleteOne {
	return c.DeleteOneID(ct.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoolroomTypeClient) DeleteOneID(id int) *CoolroomTypeDeleteOne {
	builder := c.Delete().Where(coolroomtype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoolroomTypeDeleteOne{builder}
}

// Create returns a query builder for CoolroomType.
func (c *CoolroomTypeClient) Query() *CoolroomTypeQuery {
	return &CoolroomTypeQuery{config: c.config}
}

// Get returns a CoolroomType entity by its id.
func (c *CoolroomTypeClient) Get(ctx context.Context, id int) (*CoolroomType, error) {
	return c.Query().Where(coolroomtype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoolroomTypeClient) GetX(ctx context.Context, id int) *CoolroomType {
	ct, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ct
}

// QueryCoolrooms queries the coolrooms edge of a CoolroomType.
func (c *CoolroomTypeClient) QueryCoolrooms(ct *CoolroomType) *CoolroomQuery {
	query := &CoolroomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ct.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(coolroomtype.Table, coolroomtype.FieldID, id),
			sqlgraph.To(coolroom.Table, coolroom.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, coolroomtype.CoolroomsTable, coolroomtype.CoolroomsColumn),
		)
		fromV = sqlgraph.Neighbors(ct.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CoolroomTypeClient) Hooks() []Hook {
	return c.hooks.CoolroomType
}

// DeceasedReceiveClient is a client for the DeceasedReceive schema.
type DeceasedReceiveClient struct {
	config
}

// NewDeceasedReceiveClient returns a client for the DeceasedReceive from the given config.
func NewDeceasedReceiveClient(c config) *DeceasedReceiveClient {
	return &DeceasedReceiveClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deceasedreceive.Hooks(f(g(h())))`.
func (c *DeceasedReceiveClient) Use(hooks ...Hook) {
	c.hooks.DeceasedReceive = append(c.hooks.DeceasedReceive, hooks...)
}

// Create returns a create builder for DeceasedReceive.
func (c *DeceasedReceiveClient) Create() *DeceasedReceiveCreate {
	mutation := newDeceasedReceiveMutation(c.config, OpCreate)
	return &DeceasedReceiveCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for DeceasedReceive.
func (c *DeceasedReceiveClient) Update() *DeceasedReceiveUpdate {
	mutation := newDeceasedReceiveMutation(c.config, OpUpdate)
	return &DeceasedReceiveUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeceasedReceiveClient) UpdateOne(dr *DeceasedReceive) *DeceasedReceiveUpdateOne {
	mutation := newDeceasedReceiveMutation(c.config, OpUpdateOne, withDeceasedReceive(dr))
	return &DeceasedReceiveUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeceasedReceiveClient) UpdateOneID(id int) *DeceasedReceiveUpdateOne {
	mutation := newDeceasedReceiveMutation(c.config, OpUpdateOne, withDeceasedReceiveID(id))
	return &DeceasedReceiveUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DeceasedReceive.
func (c *DeceasedReceiveClient) Delete() *DeceasedReceiveDelete {
	mutation := newDeceasedReceiveMutation(c.config, OpDelete)
	return &DeceasedReceiveDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeceasedReceiveClient) DeleteOne(dr *DeceasedReceive) *DeceasedReceiveDeleteOne {
	return c.DeleteOneID(dr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeceasedReceiveClient) DeleteOneID(id int) *DeceasedReceiveDeleteOne {
	builder := c.Delete().Where(deceasedreceive.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeceasedReceiveDeleteOne{builder}
}

// Create returns a query builder for DeceasedReceive.
func (c *DeceasedReceiveClient) Query() *DeceasedReceiveQuery {
	return &DeceasedReceiveQuery{config: c.config}
}

// Get returns a DeceasedReceive entity by its id.
func (c *DeceasedReceiveClient) Get(ctx context.Context, id int) (*DeceasedReceive, error) {
	return c.Query().Where(deceasedreceive.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeceasedReceiveClient) GetX(ctx context.Context, id int) *DeceasedReceive {
	dr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return dr
}

// QueryDoctor queries the doctor edge of a DeceasedReceive.
func (c *DeceasedReceiveClient) QueryDoctor(dr *DeceasedReceive) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.DoctorTable, deceasedreceive.DoctorColumn),
		)
		fromV = sqlgraph.Neighbors(dr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelative queries the relative edge of a DeceasedReceive.
func (c *DeceasedReceiveClient) QueryRelative(dr *DeceasedReceive) *RelativeQuery {
	query := &RelativeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, id),
			sqlgraph.To(relative.Table, relative.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.RelativeTable, deceasedreceive.RelativeColumn),
		)
		fromV = sqlgraph.Neighbors(dr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCoolroom queries the coolroom edge of a DeceasedReceive.
func (c *DeceasedReceiveClient) QueryCoolroom(dr *DeceasedReceive) *CoolroomQuery {
	query := &CoolroomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, id),
			sqlgraph.To(coolroom.Table, coolroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deceasedreceive.CoolroomTable, deceasedreceive.CoolroomColumn),
		)
		fromV = sqlgraph.Neighbors(dr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPatient queries the patient edge of a DeceasedReceive.
func (c *DeceasedReceiveClient) QueryPatient(dr *DeceasedReceive) *PatientQuery {
	query := &PatientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deceasedreceive.Table, deceasedreceive.FieldID, id),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, deceasedreceive.PatientTable, deceasedreceive.PatientColumn),
		)
		fromV = sqlgraph.Neighbors(dr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DeceasedReceiveClient) Hooks() []Hook {
	return c.hooks.DeceasedReceive
}

// PatientClient is a client for the Patient schema.
type PatientClient struct {
	config
}

// NewPatientClient returns a client for the Patient from the given config.
func NewPatientClient(c config) *PatientClient {
	return &PatientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `patient.Hooks(f(g(h())))`.
func (c *PatientClient) Use(hooks ...Hook) {
	c.hooks.Patient = append(c.hooks.Patient, hooks...)
}

// Create returns a create builder for Patient.
func (c *PatientClient) Create() *PatientCreate {
	mutation := newPatientMutation(c.config, OpCreate)
	return &PatientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Patient.
func (c *PatientClient) Update() *PatientUpdate {
	mutation := newPatientMutation(c.config, OpUpdate)
	return &PatientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PatientClient) UpdateOne(pa *Patient) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatient(pa))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PatientClient) UpdateOneID(id int) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatientID(id))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Patient.
func (c *PatientClient) Delete() *PatientDelete {
	mutation := newPatientMutation(c.config, OpDelete)
	return &PatientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PatientClient) DeleteOne(pa *Patient) *PatientDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PatientClient) DeleteOneID(id int) *PatientDeleteOne {
	builder := c.Delete().Where(patient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PatientDeleteOne{builder}
}

// Create returns a query builder for Patient.
func (c *PatientClient) Query() *PatientQuery {
	return &PatientQuery{config: c.config}
}

// Get returns a Patient entity by its id.
func (c *PatientClient) Get(ctx context.Context, id int) (*Patient, error) {
	return c.Query().Where(patient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PatientClient) GetX(ctx context.Context, id int) *Patient {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QueryDeceasedreceives queries the deceasedreceives edge of a Patient.
func (c *PatientClient) QueryDeceasedreceives(pa *Patient) *DeceasedReceiveQuery {
	query := &DeceasedReceiveQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, id),
			sqlgraph.To(deceasedreceive.Table, deceasedreceive.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, patient.DeceasedreceivesTable, patient.DeceasedreceivesColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PatientClient) Hooks() []Hook {
	return c.hooks.Patient
}

// RelativeClient is a client for the Relative schema.
type RelativeClient struct {
	config
}

// NewRelativeClient returns a client for the Relative from the given config.
func NewRelativeClient(c config) *RelativeClient {
	return &RelativeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `relative.Hooks(f(g(h())))`.
func (c *RelativeClient) Use(hooks ...Hook) {
	c.hooks.Relative = append(c.hooks.Relative, hooks...)
}

// Create returns a create builder for Relative.
func (c *RelativeClient) Create() *RelativeCreate {
	mutation := newRelativeMutation(c.config, OpCreate)
	return &RelativeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Relative.
func (c *RelativeClient) Update() *RelativeUpdate {
	mutation := newRelativeMutation(c.config, OpUpdate)
	return &RelativeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RelativeClient) UpdateOne(r *Relative) *RelativeUpdateOne {
	mutation := newRelativeMutation(c.config, OpUpdateOne, withRelative(r))
	return &RelativeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RelativeClient) UpdateOneID(id int) *RelativeUpdateOne {
	mutation := newRelativeMutation(c.config, OpUpdateOne, withRelativeID(id))
	return &RelativeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Relative.
func (c *RelativeClient) Delete() *RelativeDelete {
	mutation := newRelativeMutation(c.config, OpDelete)
	return &RelativeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RelativeClient) DeleteOne(r *Relative) *RelativeDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RelativeClient) DeleteOneID(id int) *RelativeDeleteOne {
	builder := c.Delete().Where(relative.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RelativeDeleteOne{builder}
}

// Create returns a query builder for Relative.
func (c *RelativeClient) Query() *RelativeQuery {
	return &RelativeQuery{config: c.config}
}

// Get returns a Relative entity by its id.
func (c *RelativeClient) Get(ctx context.Context, id int) (*Relative, error) {
	return c.Query().Where(relative.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RelativeClient) GetX(ctx context.Context, id int) *Relative {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryDeceasedreceives queries the deceasedreceives edge of a Relative.
func (c *RelativeClient) QueryDeceasedreceives(r *Relative) *DeceasedReceiveQuery {
	query := &DeceasedReceiveQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relative.Table, relative.FieldID, id),
			sqlgraph.To(deceasedreceive.Table, deceasedreceive.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, relative.DeceasedreceivesTable, relative.DeceasedreceivesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RelativeClient) Hooks() []Hook {
	return c.hooks.Relative
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryDeceasedreceives queries the deceasedreceives edge of a User.
func (c *UserClient) QueryDeceasedreceives(u *User) *DeceasedReceiveQuery {
	query := &DeceasedReceiveQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(deceasedreceive.Table, deceasedreceive.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DeceasedreceivesTable, user.DeceasedreceivesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
