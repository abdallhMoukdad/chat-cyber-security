// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject1/ent/chat"
	"awesomeProject1/ent/message"
	"awesomeProject1/ent/predicate"
	"awesomeProject1/ent/student"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *StudentUpdate) SetNillableName(s *string) *StudentUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetEnrollmentDate sets the "enrollment_date" field.
func (su *StudentUpdate) SetEnrollmentDate(t time.Time) *StudentUpdate {
	su.mutation.SetEnrollmentDate(t)
	return su
}

// SetNillableEnrollmentDate sets the "enrollment_date" field if the given value is not nil.
func (su *StudentUpdate) SetNillableEnrollmentDate(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetEnrollmentDate(*t)
	}
	return su
}

// SetPassword sets the "password" field.
func (su *StudentUpdate) SetPassword(s string) *StudentUpdate {
	su.mutation.SetPassword(s)
	return su
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (su *StudentUpdate) SetNillablePassword(s *string) *StudentUpdate {
	if s != nil {
		su.SetPassword(*s)
	}
	return su
}

// SetNationalNumber sets the "national_number" field.
func (su *StudentUpdate) SetNationalNumber(s string) *StudentUpdate {
	su.mutation.SetNationalNumber(s)
	return su
}

// SetNillableNationalNumber sets the "national_number" field if the given value is not nil.
func (su *StudentUpdate) SetNillableNationalNumber(s *string) *StudentUpdate {
	if s != nil {
		su.SetNationalNumber(*s)
	}
	return su
}

// SetPhoneNumber sets the "phone_number" field.
func (su *StudentUpdate) SetPhoneNumber(s string) *StudentUpdate {
	su.mutation.SetPhoneNumber(s)
	return su
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (su *StudentUpdate) SetNillablePhoneNumber(s *string) *StudentUpdate {
	if s != nil {
		su.SetPhoneNumber(*s)
	}
	return su
}

// SetHomeLocation sets the "home_location" field.
func (su *StudentUpdate) SetHomeLocation(s string) *StudentUpdate {
	su.mutation.SetHomeLocation(s)
	return su
}

// SetNillableHomeLocation sets the "home_location" field if the given value is not nil.
func (su *StudentUpdate) SetNillableHomeLocation(s *string) *StudentUpdate {
	if s != nil {
		su.SetHomeLocation(*s)
	}
	return su
}

// SetEnycrptionKey sets the "enycrption_key" field.
func (su *StudentUpdate) SetEnycrptionKey(s string) *StudentUpdate {
	su.mutation.SetEnycrptionKey(s)
	return su
}

// SetNillableEnycrptionKey sets the "enycrption_key" field if the given value is not nil.
func (su *StudentUpdate) SetNillableEnycrptionKey(s *string) *StudentUpdate {
	if s != nil {
		su.SetEnycrptionKey(*s)
	}
	return su
}

// SetChatsID sets the "chats" edge to the Chat entity by ID.
func (su *StudentUpdate) SetChatsID(id int) *StudentUpdate {
	su.mutation.SetChatsID(id)
	return su
}

// SetNillableChatsID sets the "chats" edge to the Chat entity by ID if the given value is not nil.
func (su *StudentUpdate) SetNillableChatsID(id *int) *StudentUpdate {
	if id != nil {
		su = su.SetChatsID(*id)
	}
	return su
}

// SetChats sets the "chats" edge to the Chat entity.
func (su *StudentUpdate) SetChats(c *Chat) *StudentUpdate {
	return su.SetChatsID(c.ID)
}

// SetSentMessagesID sets the "sent_messages" edge to the Message entity by ID.
func (su *StudentUpdate) SetSentMessagesID(id int) *StudentUpdate {
	su.mutation.SetSentMessagesID(id)
	return su
}

// SetNillableSentMessagesID sets the "sent_messages" edge to the Message entity by ID if the given value is not nil.
func (su *StudentUpdate) SetNillableSentMessagesID(id *int) *StudentUpdate {
	if id != nil {
		su = su.SetSentMessagesID(*id)
	}
	return su
}

// SetSentMessages sets the "sent_messages" edge to the Message entity.
func (su *StudentUpdate) SetSentMessages(m *Message) *StudentUpdate {
	return su.SetSentMessagesID(m.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearChats clears the "chats" edge to the Chat entity.
func (su *StudentUpdate) ClearChats() *StudentUpdate {
	su.mutation.ClearChats()
	return su
}

// ClearSentMessages clears the "sent_messages" edge to the Message entity.
func (su *StudentUpdate) ClearSentMessages() *StudentUpdate {
	su.mutation.ClearSentMessages()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StudentUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := student.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Student.name": %w`, err)}
		}
	}
	return nil
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.EnrollmentDate(); ok {
		_spec.SetField(student.FieldEnrollmentDate, field.TypeTime, value)
	}
	if value, ok := su.mutation.Password(); ok {
		_spec.SetField(student.FieldPassword, field.TypeString, value)
	}
	if value, ok := su.mutation.NationalNumber(); ok {
		_spec.SetField(student.FieldNationalNumber, field.TypeString, value)
	}
	if value, ok := su.mutation.PhoneNumber(); ok {
		_spec.SetField(student.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := su.mutation.HomeLocation(); ok {
		_spec.SetField(student.FieldHomeLocation, field.TypeString, value)
	}
	if value, ok := su.mutation.EnycrptionKey(); ok {
		_spec.SetField(student.FieldEnycrptionKey, field.TypeString, value)
	}
	if su.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.ChatsTable,
			Columns: []string{student.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.ChatsTable,
			Columns: []string{student.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SentMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.SentMessagesTable,
			Columns: []string{student.SentMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SentMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.SentMessagesTable,
			Columns: []string{student.SentMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableName(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetEnrollmentDate sets the "enrollment_date" field.
func (suo *StudentUpdateOne) SetEnrollmentDate(t time.Time) *StudentUpdateOne {
	suo.mutation.SetEnrollmentDate(t)
	return suo
}

// SetNillableEnrollmentDate sets the "enrollment_date" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableEnrollmentDate(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetEnrollmentDate(*t)
	}
	return suo
}

// SetPassword sets the "password" field.
func (suo *StudentUpdateOne) SetPassword(s string) *StudentUpdateOne {
	suo.mutation.SetPassword(s)
	return suo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillablePassword(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetPassword(*s)
	}
	return suo
}

// SetNationalNumber sets the "national_number" field.
func (suo *StudentUpdateOne) SetNationalNumber(s string) *StudentUpdateOne {
	suo.mutation.SetNationalNumber(s)
	return suo
}

// SetNillableNationalNumber sets the "national_number" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableNationalNumber(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetNationalNumber(*s)
	}
	return suo
}

// SetPhoneNumber sets the "phone_number" field.
func (suo *StudentUpdateOne) SetPhoneNumber(s string) *StudentUpdateOne {
	suo.mutation.SetPhoneNumber(s)
	return suo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillablePhoneNumber(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetPhoneNumber(*s)
	}
	return suo
}

// SetHomeLocation sets the "home_location" field.
func (suo *StudentUpdateOne) SetHomeLocation(s string) *StudentUpdateOne {
	suo.mutation.SetHomeLocation(s)
	return suo
}

// SetNillableHomeLocation sets the "home_location" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableHomeLocation(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetHomeLocation(*s)
	}
	return suo
}

// SetEnycrptionKey sets the "enycrption_key" field.
func (suo *StudentUpdateOne) SetEnycrptionKey(s string) *StudentUpdateOne {
	suo.mutation.SetEnycrptionKey(s)
	return suo
}

// SetNillableEnycrptionKey sets the "enycrption_key" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableEnycrptionKey(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetEnycrptionKey(*s)
	}
	return suo
}

// SetChatsID sets the "chats" edge to the Chat entity by ID.
func (suo *StudentUpdateOne) SetChatsID(id int) *StudentUpdateOne {
	suo.mutation.SetChatsID(id)
	return suo
}

// SetNillableChatsID sets the "chats" edge to the Chat entity by ID if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableChatsID(id *int) *StudentUpdateOne {
	if id != nil {
		suo = suo.SetChatsID(*id)
	}
	return suo
}

// SetChats sets the "chats" edge to the Chat entity.
func (suo *StudentUpdateOne) SetChats(c *Chat) *StudentUpdateOne {
	return suo.SetChatsID(c.ID)
}

// SetSentMessagesID sets the "sent_messages" edge to the Message entity by ID.
func (suo *StudentUpdateOne) SetSentMessagesID(id int) *StudentUpdateOne {
	suo.mutation.SetSentMessagesID(id)
	return suo
}

// SetNillableSentMessagesID sets the "sent_messages" edge to the Message entity by ID if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableSentMessagesID(id *int) *StudentUpdateOne {
	if id != nil {
		suo = suo.SetSentMessagesID(*id)
	}
	return suo
}

// SetSentMessages sets the "sent_messages" edge to the Message entity.
func (suo *StudentUpdateOne) SetSentMessages(m *Message) *StudentUpdateOne {
	return suo.SetSentMessagesID(m.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearChats clears the "chats" edge to the Chat entity.
func (suo *StudentUpdateOne) ClearChats() *StudentUpdateOne {
	suo.mutation.ClearChats()
	return suo
}

// ClearSentMessages clears the "sent_messages" edge to the Message entity.
func (suo *StudentUpdateOne) ClearSentMessages() *StudentUpdateOne {
	suo.mutation.ClearSentMessages()
	return suo
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StudentUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := student.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Student.name": %w`, err)}
		}
	}
	return nil
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.EnrollmentDate(); ok {
		_spec.SetField(student.FieldEnrollmentDate, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Password(); ok {
		_spec.SetField(student.FieldPassword, field.TypeString, value)
	}
	if value, ok := suo.mutation.NationalNumber(); ok {
		_spec.SetField(student.FieldNationalNumber, field.TypeString, value)
	}
	if value, ok := suo.mutation.PhoneNumber(); ok {
		_spec.SetField(student.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := suo.mutation.HomeLocation(); ok {
		_spec.SetField(student.FieldHomeLocation, field.TypeString, value)
	}
	if value, ok := suo.mutation.EnycrptionKey(); ok {
		_spec.SetField(student.FieldEnycrptionKey, field.TypeString, value)
	}
	if suo.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.ChatsTable,
			Columns: []string{student.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.ChatsTable,
			Columns: []string{student.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SentMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.SentMessagesTable,
			Columns: []string{student.SentMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SentMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   student.SentMessagesTable,
			Columns: []string{student.SentMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
