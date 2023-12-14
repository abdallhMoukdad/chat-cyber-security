// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject1/ent/chat"
	"awesomeProject1/ent/message"
	"awesomeProject1/ent/student"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// EnrollmentDate holds the value of the "enrollment_date" field.
	EnrollmentDate time.Time `json:"enrollment_date,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// NationalNumber holds the value of the "national_number" field.
	NationalNumber string `json:"national_number,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// HomeLocation holds the value of the "home_location" field.
	HomeLocation string `json:"home_location,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges        StudentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Chats holds the value of the chats edge.
	Chats *Chat `json:"chats,omitempty"`
	// SentMessages holds the value of the sent_messages edge.
	SentMessages *Message `json:"sent_messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChatsOrErr returns the Chats value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) ChatsOrErr() (*Chat, error) {
	if e.loadedTypes[0] {
		if e.Chats == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: chat.Label}
		}
		return e.Chats, nil
	}
	return nil, &NotLoadedError{edge: "chats"}
}

// SentMessagesOrErr returns the SentMessages value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) SentMessagesOrErr() (*Message, error) {
	if e.loadedTypes[1] {
		if e.SentMessages == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: message.Label}
		}
		return e.SentMessages, nil
	}
	return nil, &NotLoadedError{edge: "sent_messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			values[i] = new(sql.NullInt64)
		case student.FieldName, student.FieldPassword, student.FieldNationalNumber, student.FieldPhoneNumber, student.FieldHomeLocation:
			values[i] = new(sql.NullString)
		case student.FieldEnrollmentDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case student.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case student.FieldEnrollmentDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field enrollment_date", values[i])
			} else if value.Valid {
				s.EnrollmentDate = value.Time
			}
		case student.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				s.Password = value.String
			}
		case student.FieldNationalNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field national_number", values[i])
			} else if value.Valid {
				s.NationalNumber = value.String
			}
		case student.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				s.PhoneNumber = value.String
			}
		case student.FieldHomeLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_location", values[i])
			} else if value.Valid {
				s.HomeLocation = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Student.
// This includes values selected through modifiers, order, etc.
func (s *Student) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryChats queries the "chats" edge of the Student entity.
func (s *Student) QueryChats() *ChatQuery {
	return NewStudentClient(s.config).QueryChats(s)
}

// QuerySentMessages queries the "sent_messages" edge of the Student entity.
func (s *Student) QuerySentMessages() *MessageQuery {
	return NewStudentClient(s.config).QuerySentMessages(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("enrollment_date=")
	builder.WriteString(s.EnrollmentDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(s.Password)
	builder.WriteString(", ")
	builder.WriteString("national_number=")
	builder.WriteString(s.NationalNumber)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(s.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("home_location=")
	builder.WriteString(s.HomeLocation)
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student
