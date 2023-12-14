// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject1/ent/predicate"
	"awesomeProject1/ent/professor"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfessorDelete is the builder for deleting a Professor entity.
type ProfessorDelete struct {
	config
	hooks    []Hook
	mutation *ProfessorMutation
}

// Where appends a list predicates to the ProfessorDelete builder.
func (pd *ProfessorDelete) Where(ps ...predicate.Professor) *ProfessorDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProfessorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProfessorDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProfessorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(professor.Table, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProfessorDeleteOne is the builder for deleting a single Professor entity.
type ProfessorDeleteOne struct {
	pd *ProfessorDelete
}

// Where appends a list predicates to the ProfessorDelete builder.
func (pdo *ProfessorDeleteOne) Where(ps ...predicate.Professor) *ProfessorDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProfessorDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{professor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProfessorDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
