package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Professor schema.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		//field.String("email").NotEmpty(),
		field.Time("hire_date").Default(time.Now),
		field.String("password").NotEmpty(),
	}
}
func (Professor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chats", Chat.Type).Unique(),
		edge.To("received_messages", Message.Type).Unique(),
	}
}
