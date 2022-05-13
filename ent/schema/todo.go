package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").MaxLen(500),
		field.Bool("done").Default(false),
		field.Int("user_id").Optional(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("todos").
			Unique().
			Field("user_id"),
		edge.To("reminders", Reminder.Type),
	}
}
