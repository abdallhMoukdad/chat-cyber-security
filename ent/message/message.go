// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// EdgeChat holds the string denoting the chat edge name in mutations.
	EdgeChat = "chat"
	// EdgeSender holds the string denoting the sender edge name in mutations.
	EdgeSender = "sender"
	// EdgeReceiver holds the string denoting the receiver edge name in mutations.
	EdgeReceiver = "receiver"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// ChatTable is the table that holds the chat relation/edge.
	ChatTable = "messages"
	// ChatInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatInverseTable = "chats"
	// ChatColumn is the table column denoting the chat relation/edge.
	ChatColumn = "chat_messages"
	// SenderTable is the table that holds the sender relation/edge.
	SenderTable = "messages"
	// SenderInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	SenderInverseTable = "students"
	// SenderColumn is the table column denoting the sender relation/edge.
	SenderColumn = "student_sent_messages"
	// ReceiverTable is the table that holds the receiver relation/edge.
	ReceiverTable = "messages"
	// ReceiverInverseTable is the table name for the Professor entity.
	// It exists in this package in order to avoid circular dependency with the "professor" package.
	ReceiverInverseTable = "professors"
	// ReceiverColumn is the table column denoting the receiver relation/edge.
	ReceiverColumn = "professor_received_messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"chat_messages",
	"professor_received_messages",
	"student_sent_messages",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp func() time.Time
)

// OrderOption defines the ordering options for the Message queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByChatField orders the results by chat field.
func ByChatField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatStep(), sql.OrderByField(field, opts...))
	}
}

// BySenderField orders the results by sender field.
func BySenderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiverField orders the results by receiver field.
func ByReceiverField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiverStep(), sql.OrderByField(field, opts...))
	}
}
func newChatStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ChatTable, ChatColumn),
	)
}
func newSenderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SenderTable, SenderColumn),
	)
}
func newReceiverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ReceiverTable, ReceiverColumn),
	)
}
