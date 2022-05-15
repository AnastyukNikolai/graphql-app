package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"graphql-app/internal/repositories"
	"graphql-app/internal/seeders"
)

func main() {
	client := repositories.OpenDB()
	repo := repositories.NewRepository(client)
	seeder := seeders.NewSeeder(repo)
	err := seed(seeder)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
}

func seed(seeder *seeders.Seeder) error {
	err := seeder.User.SeedUsers()
	if err !=nil {
		return err
	}
	err = seeder.Todo.SeedTodos()
	if err !=nil {
		return err
	}
	err = seeder.Reminder.SeedReminders()
	if err !=nil {
		return err
	}
	return nil
}
