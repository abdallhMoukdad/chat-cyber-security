package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Message schema.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),
		field.Time("timestamp").Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).Ref("messages").Unique(),
		edge.From("sender", Student.Type).Ref("sent_messages").Unique(),
		edge.From("receiver", Professor.Type).Ref("received_messages").Unique(),
	}
}
