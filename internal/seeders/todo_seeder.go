package seeders

import (
	"graphql-app/graph/model"
	"graphql-app/internal/repositories"
)

type TodoSeeder struct {
	repo *repositories.Repository
}

func NewTodoSeeder(repo *repositories.Repository) *TodoSeeder {
	return &TodoSeeder{repo: repo}
}

func (seeder *TodoSeeder) SeedTodos() error {
	_ ,err := seeder.repo.CreateTodo(&model.NewTodo{
		Text: "Todo",
		UserID: 1,
	})
	if err != nil {
		return err
	}
	return nil
}