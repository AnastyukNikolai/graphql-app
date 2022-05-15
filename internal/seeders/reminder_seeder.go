package seeders

import (
	"graphql-app/graph/model"
	"graphql-app/internal/repositories"
)

type ReminderSeeder struct {
	repo *repositories.Repository
}

func NewReminderSeeder(repo *repositories.Repository) *ReminderSeeder {
	return &ReminderSeeder{repo: repo}
}

func (seeder *ReminderSeeder) SeedReminders() error {
	_ ,err := seeder.repo.CreateReminder(&model.NewReminder{
		Text: "Reminder",
		TodoID: 1,
	})
	if err != nil {
		return err
	}
	return nil
}