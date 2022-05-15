package seeders

import (
	"graphql-app/graph/model"
	"graphql-app/internal/repositories"
)

type UserSeeder struct {
	repo *repositories.Repository
}

func NewUserSeeder(repo *repositories.Repository) *UserSeeder {
	return &UserSeeder{repo: repo}
}

func (seeder *UserSeeder) SeedUsers() error {
	_ ,err := seeder.repo.CreateUser(&model.NewUser{
		Name: "User",
		CreditCard: "3334 4322 4432 3332",
	})
	if err != nil {
		return err
	}
	return nil
}