// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"pvp/internal/data/ent/app"
	"pvp/internal/data/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AppUpdate is the builder for updating App entities.
type AppUpdate struct {
	config
	hooks    []Hook
	mutation *AppMutation
}

// Where adds a new predicate for the AppUpdate builder.
func (au *AppUpdate) Where(ps ...predicate.App) *AppUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetAppKey sets the "app_key" field.
func (au *AppUpdate) SetAppKey(s string) *AppUpdate {
	au.mutation.SetAppKey(s)
	return au
}

// SetAppScreen sets the "app_screen" field.
func (au *AppUpdate) SetAppScreen(s string) *AppUpdate {
	au.mutation.SetAppScreen(s)
	return au
}

// SetNillableAppScreen sets the "app_screen" field if the given value is not nil.
func (au *AppUpdate) SetNillableAppScreen(s *string) *AppUpdate {
	if s != nil {
		au.SetAppScreen(*s)
	}
	return au
}

// SetAppStatus sets the "app_status" field.
func (au *AppUpdate) SetAppStatus(i int8) *AppUpdate {
	au.mutation.ResetAppStatus()
	au.mutation.SetAppStatus(i)
	return au
}

// SetNillableAppStatus sets the "app_status" field if the given value is not nil.
func (au *AppUpdate) SetNillableAppStatus(i *int8) *AppUpdate {
	if i != nil {
		au.SetAppStatus(*i)
	}
	return au
}

// AddAppStatus adds i to the "app_status" field.
func (au *AppUpdate) AddAppStatus(i int8) *AppUpdate {
	au.mutation.AddAppStatus(i)
	return au
}

// SetCreateTime sets the "create_time" field.
func (au *AppUpdate) SetCreateTime(i int16) *AppUpdate {
	au.mutation.ResetCreateTime()
	au.mutation.SetCreateTime(i)
	return au
}

// AddCreateTime adds i to the "create_time" field.
func (au *AppUpdate) AddCreateTime(i int16) *AppUpdate {
	au.mutation.AddCreateTime(i)
	return au
}

// SetDeleteTime sets the "delete_time" field.
func (au *AppUpdate) SetDeleteTime(i int16) *AppUpdate {
	au.mutation.ResetDeleteTime()
	au.mutation.SetDeleteTime(i)
	return au
}

// AddDeleteTime adds i to the "delete_time" field.
func (au *AppUpdate) AddDeleteTime(i int16) *AppUpdate {
	au.mutation.AddDeleteTime(i)
	return au
}

// SetUpdateTime sets the "update_time" field.
func (au *AppUpdate) SetUpdateTime(i int16) *AppUpdate {
	au.mutation.ResetUpdateTime()
	au.mutation.SetUpdateTime(i)
	return au
}

// AddUpdateTime adds i to the "update_time" field.
func (au *AppUpdate) AddUpdateTime(i int16) *AppUpdate {
	au.mutation.AddUpdateTime(i)
	return au
}

// Mutation returns the AppMutation object of the builder.
func (au *AppUpdate) Mutation() *AppMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AppUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.AppKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldAppKey,
		})
	}
	if value, ok := au.mutation.AppScreen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldAppScreen,
		})
	}
	if value, ok := au.mutation.AppStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: app.FieldAppStatus,
		})
	}
	if value, ok := au.mutation.AddedAppStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: app.FieldAppStatus,
		})
	}
	if value, ok := au.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldCreateTime,
		})
	}
	if value, ok := au.mutation.AddedCreateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldCreateTime,
		})
	}
	if value, ok := au.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldDeleteTime,
		})
	}
	if value, ok := au.mutation.AddedDeleteTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldDeleteTime,
		})
	}
	if value, ok := au.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldUpdateTime,
		})
	}
	if value, ok := au.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AppUpdateOne is the builder for updating a single App entity.
type AppUpdateOne struct {
	config
	hooks    []Hook
	mutation *AppMutation
}

// SetAppKey sets the "app_key" field.
func (auo *AppUpdateOne) SetAppKey(s string) *AppUpdateOne {
	auo.mutation.SetAppKey(s)
	return auo
}

// SetAppScreen sets the "app_screen" field.
func (auo *AppUpdateOne) SetAppScreen(s string) *AppUpdateOne {
	auo.mutation.SetAppScreen(s)
	return auo
}

// SetNillableAppScreen sets the "app_screen" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAppScreen(s *string) *AppUpdateOne {
	if s != nil {
		auo.SetAppScreen(*s)
	}
	return auo
}

// SetAppStatus sets the "app_status" field.
func (auo *AppUpdateOne) SetAppStatus(i int8) *AppUpdateOne {
	auo.mutation.ResetAppStatus()
	auo.mutation.SetAppStatus(i)
	return auo
}

// SetNillableAppStatus sets the "app_status" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableAppStatus(i *int8) *AppUpdateOne {
	if i != nil {
		auo.SetAppStatus(*i)
	}
	return auo
}

// AddAppStatus adds i to the "app_status" field.
func (auo *AppUpdateOne) AddAppStatus(i int8) *AppUpdateOne {
	auo.mutation.AddAppStatus(i)
	return auo
}

// SetCreateTime sets the "create_time" field.
func (auo *AppUpdateOne) SetCreateTime(i int16) *AppUpdateOne {
	auo.mutation.ResetCreateTime()
	auo.mutation.SetCreateTime(i)
	return auo
}

// AddCreateTime adds i to the "create_time" field.
func (auo *AppUpdateOne) AddCreateTime(i int16) *AppUpdateOne {
	auo.mutation.AddCreateTime(i)
	return auo
}

// SetDeleteTime sets the "delete_time" field.
func (auo *AppUpdateOne) SetDeleteTime(i int16) *AppUpdateOne {
	auo.mutation.ResetDeleteTime()
	auo.mutation.SetDeleteTime(i)
	return auo
}

// AddDeleteTime adds i to the "delete_time" field.
func (auo *AppUpdateOne) AddDeleteTime(i int16) *AppUpdateOne {
	auo.mutation.AddDeleteTime(i)
	return auo
}

// SetUpdateTime sets the "update_time" field.
func (auo *AppUpdateOne) SetUpdateTime(i int16) *AppUpdateOne {
	auo.mutation.ResetUpdateTime()
	auo.mutation.SetUpdateTime(i)
	return auo
}

// AddUpdateTime adds i to the "update_time" field.
func (auo *AppUpdateOne) AddUpdateTime(i int16) *AppUpdateOne {
	auo.mutation.AddUpdateTime(i)
	return auo
}

// Mutation returns the AppMutation object of the builder.
func (auo *AppUpdateOne) Mutation() *AppMutation {
	return auo.mutation
}

// Save executes the query and returns the updated App entity.
func (auo *AppUpdateOne) Save(ctx context.Context) (*App, error) {
	var (
		err  error
		node *App
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppUpdateOne) SaveX(ctx context.Context) *App {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AppUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AppUpdateOne) sqlSave(ctx context.Context) (_node *App, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing App.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.AppKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldAppKey,
		})
	}
	if value, ok := auo.mutation.AppScreen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldAppScreen,
		})
	}
	if value, ok := auo.mutation.AppStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: app.FieldAppStatus,
		})
	}
	if value, ok := auo.mutation.AddedAppStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: app.FieldAppStatus,
		})
	}
	if value, ok := auo.mutation.CreateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldCreateTime,
		})
	}
	if value, ok := auo.mutation.AddedCreateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldCreateTime,
		})
	}
	if value, ok := auo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldDeleteTime,
		})
	}
	if value, ok := auo.mutation.AddedDeleteTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldDeleteTime,
		})
	}
	if value, ok := auo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldUpdateTime,
		})
	}
	if value, ok := auo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: app.FieldUpdateTime,
		})
	}
	_node = &App{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}