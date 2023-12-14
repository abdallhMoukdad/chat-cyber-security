// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject1/ent/predicate"
	"awesomeProject1/ent/university"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UniversityUpdate is the builder for updating University entities.
type UniversityUpdate struct {
	config
	hooks    []Hook
	mutation *UniversityMutation
}

// Where appends a list predicates to the UniversityUpdate builder.
func (uu *UniversityUpdate) Where(ps ...predicate.University) *UniversityUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UniversityUpdate) SetName(s string) *UniversityUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UniversityUpdate) SetNillableName(s *string) *UniversityUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetLocation sets the "location" field.
func (uu *UniversityUpdate) SetLocation(s string) *UniversityUpdate {
	uu.mutation.SetLocation(s)
	return uu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uu *UniversityUpdate) SetNillableLocation(s *string) *UniversityUpdate {
	if s != nil {
		uu.SetLocation(*s)
	}
	return uu
}

// SetEstablishedDate sets the "established_date" field.
func (uu *UniversityUpdate) SetEstablishedDate(t time.Time) *UniversityUpdate {
	uu.mutation.SetEstablishedDate(t)
	return uu
}

// SetNillableEstablishedDate sets the "established_date" field if the given value is not nil.
func (uu *UniversityUpdate) SetNillableEstablishedDate(t *time.Time) *UniversityUpdate {
	if t != nil {
		uu.SetEstablishedDate(*t)
	}
	return uu
}

// Mutation returns the UniversityMutation object of the builder.
func (uu *UniversityUpdate) Mutation() *UniversityMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UniversityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UniversityUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UniversityUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UniversityUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UniversityUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := university.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "University.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Location(); ok {
		if err := university.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "University.location": %w`, err)}
		}
	}
	return nil
}

func (uu *UniversityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(university.Table, university.Columns, sqlgraph.NewFieldSpec(university.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(university.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Location(); ok {
		_spec.SetField(university.FieldLocation, field.TypeString, value)
	}
	if value, ok := uu.mutation.EstablishedDate(); ok {
		_spec.SetField(university.FieldEstablishedDate, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{university.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UniversityUpdateOne is the builder for updating a single University entity.
type UniversityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UniversityMutation
}

// SetName sets the "name" field.
func (uuo *UniversityUpdateOne) SetName(s string) *UniversityUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UniversityUpdateOne) SetNillableName(s *string) *UniversityUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetLocation sets the "location" field.
func (uuo *UniversityUpdateOne) SetLocation(s string) *UniversityUpdateOne {
	uuo.mutation.SetLocation(s)
	return uuo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (uuo *UniversityUpdateOne) SetNillableLocation(s *string) *UniversityUpdateOne {
	if s != nil {
		uuo.SetLocation(*s)
	}
	return uuo
}

// SetEstablishedDate sets the "established_date" field.
func (uuo *UniversityUpdateOne) SetEstablishedDate(t time.Time) *UniversityUpdateOne {
	uuo.mutation.SetEstablishedDate(t)
	return uuo
}

// SetNillableEstablishedDate sets the "established_date" field if the given value is not nil.
func (uuo *UniversityUpdateOne) SetNillableEstablishedDate(t *time.Time) *UniversityUpdateOne {
	if t != nil {
		uuo.SetEstablishedDate(*t)
	}
	return uuo
}

// Mutation returns the UniversityMutation object of the builder.
func (uuo *UniversityUpdateOne) Mutation() *UniversityMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UniversityUpdate builder.
func (uuo *UniversityUpdateOne) Where(ps ...predicate.University) *UniversityUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UniversityUpdateOne) Select(field string, fields ...string) *UniversityUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated University entity.
func (uuo *UniversityUpdateOne) Save(ctx context.Context) (*University, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UniversityUpdateOne) SaveX(ctx context.Context) *University {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UniversityUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UniversityUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UniversityUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := university.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "University.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Location(); ok {
		if err := university.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf(`ent: validator failed for field "University.location": %w`, err)}
		}
	}
	return nil
}

func (uuo *UniversityUpdateOne) sqlSave(ctx context.Context) (_node *University, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(university.Table, university.Columns, sqlgraph.NewFieldSpec(university.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "University.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, university.FieldID)
		for _, f := range fields {
			if !university.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != university.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(university.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Location(); ok {
		_spec.SetField(university.FieldLocation, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EstablishedDate(); ok {
		_spec.SetField(university.FieldEstablishedDate, field.TypeTime, value)
	}
	_node = &University{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{university.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
