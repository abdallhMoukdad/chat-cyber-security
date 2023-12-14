package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// University schema.
type University struct {
	ent.Schema
}

// Fields of the University.
func (University) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("location").NotEmpty(),
		field.Time("established_date").Default(time.Now),
	}
}
