// Code generated by ent, DO NOT EDIT.

package professor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the professor type in the database.
	Label = "professor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHireDate holds the string denoting the hire_date field in the database.
	FieldHireDate = "hire_date"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeChats holds the string denoting the chats edge name in mutations.
	EdgeChats = "chats"
	// EdgeReceivedMessages holds the string denoting the received_messages edge name in mutations.
	EdgeReceivedMessages = "received_messages"
	// Table holds the table name of the professor in the database.
	Table = "professors"
	// ChatsTable is the table that holds the chats relation/edge.
	ChatsTable = "chats"
	// ChatsInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatsInverseTable = "chats"
	// ChatsColumn is the table column denoting the chats relation/edge.
	ChatsColumn = "professor_chats"
	// ReceivedMessagesTable is the table that holds the received_messages relation/edge.
	ReceivedMessagesTable = "messages"
	// ReceivedMessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	ReceivedMessagesInverseTable = "messages"
	// ReceivedMessagesColumn is the table column denoting the received_messages relation/edge.
	ReceivedMessagesColumn = "professor_received_messages"
)

// Columns holds all SQL columns for professor fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHireDate,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultHireDate holds the default value on creation for the "hire_date" field.
	DefaultHireDate func() time.Time
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)

// OrderOption defines the ordering options for the Professor queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHireDate orders the results by the hire_date field.
func ByHireDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHireDate, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByChatsField orders the results by chats field.
func ByChatsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatsStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceivedMessagesField orders the results by received_messages field.
func ByReceivedMessagesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceivedMessagesStep(), sql.OrderByField(field, opts...))
	}
}
func newChatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ChatsTable, ChatsColumn),
	)
}
func newReceivedMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceivedMessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ReceivedMessagesTable, ReceivedMessagesColumn),
	)
}
