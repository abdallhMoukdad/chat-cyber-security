package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Chat schema.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).Ref("chats").Unique(),
		edge.From("professor", Professor.Type).Ref("chats").Unique(),
		edge.To("messages", Message.Type).Unique(),
	}
}
