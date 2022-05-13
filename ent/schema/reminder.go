package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reminder holds the schema definition for the Reminder entity.
type Reminder struct {
	ent.Schema
}

// Fields of the Reminder.
func (Reminder) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").MaxLen(500),
		field.Int("todo_id").Optional(),
	}
}

// Edges of the Reminder.
func (Reminder) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("todo", Todo.Type).
			Ref("reminders").
			Unique().
			Field("todo_id"),
	}
}
