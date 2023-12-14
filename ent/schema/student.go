package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type Student struct {
	ent.Schema
}

// Fields of the User.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Default("unknown"),
		//field.String("email").NotEmpty().Default("unknown"),
		field.Time("enrollment_date").Default(time.Now),
		field.String("password").Default("not registered yet"),
		field.String("national_number").Default("unknown"),

		field.String("phone_number").Default("unknown"),
		field.String("home_location").Default("unknown"),
		field.String("enycrption_key").Default("unknown"),
	}
}

// Edges of the User.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chats", Chat.Type).Unique(),
		edge.To("sent_messages", Message.Type).Unique(),
	}
}
