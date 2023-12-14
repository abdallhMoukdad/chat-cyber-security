// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject1/ent/chat"
	"awesomeProject1/ent/message"
	"awesomeProject1/ent/professor"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfessorCreate is the builder for creating a Professor entity.
type ProfessorCreate struct {
	config
	mutation *ProfessorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProfessorCreate) SetName(s string) *ProfessorCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetHireDate sets the "hire_date" field.
func (pc *ProfessorCreate) SetHireDate(t time.Time) *ProfessorCreate {
	pc.mutation.SetHireDate(t)
	return pc
}

// SetNillableHireDate sets the "hire_date" field if the given value is not nil.
func (pc *ProfessorCreate) SetNillableHireDate(t *time.Time) *ProfessorCreate {
	if t != nil {
		pc.SetHireDate(*t)
	}
	return pc
}

// SetPassword sets the "password" field.
func (pc *ProfessorCreate) SetPassword(s string) *ProfessorCreate {
	pc.mutation.SetPassword(s)
	return pc
}

// SetChatsID sets the "chats" edge to the Chat entity by ID.
func (pc *ProfessorCreate) SetChatsID(id int) *ProfessorCreate {
	pc.mutation.SetChatsID(id)
	return pc
}

// SetNillableChatsID sets the "chats" edge to the Chat entity by ID if the given value is not nil.
func (pc *ProfessorCreate) SetNillableChatsID(id *int) *ProfessorCreate {
	if id != nil {
		pc = pc.SetChatsID(*id)
	}
	return pc
}

// SetChats sets the "chats" edge to the Chat entity.
func (pc *ProfessorCreate) SetChats(c *Chat) *ProfessorCreate {
	return pc.SetChatsID(c.ID)
}

// SetReceivedMessagesID sets the "received_messages" edge to the Message entity by ID.
func (pc *ProfessorCreate) SetReceivedMessagesID(id int) *ProfessorCreate {
	pc.mutation.SetReceivedMessagesID(id)
	return pc
}

// SetNillableReceivedMessagesID sets the "received_messages" edge to the Message entity by ID if the given value is not nil.
func (pc *ProfessorCreate) SetNillableReceivedMessagesID(id *int) *ProfessorCreate {
	if id != nil {
		pc = pc.SetReceivedMessagesID(*id)
	}
	return pc
}

// SetReceivedMessages sets the "received_messages" edge to the Message entity.
func (pc *ProfessorCreate) SetReceivedMessages(m *Message) *ProfessorCreate {
	return pc.SetReceivedMessagesID(m.ID)
}

// Mutation returns the ProfessorMutation object of the builder.
func (pc *ProfessorCreate) Mutation() *ProfessorMutation {
	return pc.mutation
}

// Save creates the Professor in the database.
func (pc *ProfessorCreate) Save(ctx context.Context) (*Professor, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfessorCreate) SaveX(ctx context.Context) *Professor {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfessorCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfessorCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfessorCreate) defaults() {
	if _, ok := pc.mutation.HireDate(); !ok {
		v := professor.DefaultHireDate()
		pc.mutation.SetHireDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfessorCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Professor.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := professor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Professor.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.HireDate(); !ok {
		return &ValidationError{Name: "hire_date", err: errors.New(`ent: missing required field "Professor.hire_date"`)}
	}
	if _, ok := pc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Professor.password"`)}
	}
	if v, ok := pc.mutation.Password(); ok {
		if err := professor.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Professor.password": %w`, err)}
		}
	}
	return nil
}

func (pc *ProfessorCreate) sqlSave(ctx context.Context) (*Professor, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfessorCreate) createSpec() (*Professor, *sqlgraph.CreateSpec) {
	var (
		_node = &Professor{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(professor.Table, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(professor.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.HireDate(); ok {
		_spec.SetField(professor.FieldHireDate, field.TypeTime, value)
		_node.HireDate = value
	}
	if value, ok := pc.mutation.Password(); ok {
		_spec.SetField(professor.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if nodes := pc.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   professor.ChatsTable,
			Columns: []string{professor.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ReceivedMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   professor.ReceivedMessagesTable,
			Columns: []string{professor.ReceivedMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfessorCreateBulk is the builder for creating many Professor entities in bulk.
type ProfessorCreateBulk struct {
	config
	err      error
	builders []*ProfessorCreate
}

// Save creates the Professor entities in the database.
func (pcb *ProfessorCreateBulk) Save(ctx context.Context) ([]*Professor, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Professor, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfessorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfessorCreateBulk) SaveX(ctx context.Context) []*Professor {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfessorCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfessorCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
