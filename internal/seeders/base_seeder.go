package seeders

import (
	"graphql-app/internal/repositories"
)

type Seeder struct {
	User
	Todo
	Reminder
}

type User interface {
	SeedUsers() (err error)
}

type Todo interface {
	SeedTodos() (err error)
}

type Reminder interface {
	SeedReminders() (err error)
}

func NewSeeder(repo *repositories.Repository) *Seeder {
	return &Seeder{
		User: NewUserSeeder(repo),
		Todo: NewTodoSeeder(repo),
		Reminder: NewReminderSeeder(repo),
	}
}