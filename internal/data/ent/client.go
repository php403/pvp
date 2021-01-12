// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"pvp/internal/data/ent/migrate"

	"pvp/internal/data/ent/app"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// App is the client for interacting with the App builders.
	App *AppClient
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
	c.App = NewAppClient(c.config)
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
		ctx:    ctx,
		config: cfg,
		App:    NewAppClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
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
		config: cfg,
		App:    NewAppClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		App.
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
	c.App.Use(hooks...)
}

// AppClient is a client for the App schema.
type AppClient struct {
	config
}

// NewAppClient returns a client for the App from the given config.
func NewAppClient(c config) *AppClient {
	return &AppClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `app.Hooks(f(g(h())))`.
func (c *AppClient) Use(hooks ...Hook) {
	c.hooks.App = append(c.hooks.App, hooks...)
}

// Create returns a create builder for App.
func (c *AppClient) Create() *AppCreate {
	mutation := newAppMutation(c.config, OpCreate)
	return &AppCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of App entities.
func (c *AppClient) CreateBulk(builders ...*AppCreate) *AppCreateBulk {
	return &AppCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for App.
func (c *AppClient) Update() *AppUpdate {
	mutation := newAppMutation(c.config, OpUpdate)
	return &AppUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppClient) UpdateOne(a *App) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withApp(a))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppClient) UpdateOneID(id int) *AppUpdateOne {
	mutation := newAppMutation(c.config, OpUpdateOne, withAppID(id))
	return &AppUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for App.
func (c *AppClient) Delete() *AppDelete {
	mutation := newAppMutation(c.config, OpDelete)
	return &AppDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppClient) DeleteOne(a *App) *AppDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppClient) DeleteOneID(id int) *AppDeleteOne {
	builder := c.Delete().Where(app.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppDeleteOne{builder}
}

// Query returns a query builder for App.
func (c *AppClient) Query() *AppQuery {
	return &AppQuery{config: c.config}
}

// Get returns a App entity by its id.
func (c *AppClient) Get(ctx context.Context, id int) (*App, error) {
	return c.Query().Where(app.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppClient) GetX(ctx context.Context, id int) *App {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppClient) Hooks() []Hook {
	return c.hooks.App
}
