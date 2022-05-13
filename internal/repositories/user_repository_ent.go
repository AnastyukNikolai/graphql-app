package repositories

import (
	"context"
	"graphql-app/ent"
	"graphql-app/ent/user"
	"graphql-app/graph/model"
)

type UserRepositoryEnt struct {
	DBClient *ent.Client
}

func NewUserRepositoryEnt(DBClient *ent.Client) *UserRepositoryEnt {
	return &UserRepositoryEnt{DBClient: DBClient}
}

func (r *UserRepositoryEnt) GetUserById(userId int) (*model.User, error) {
	var userModel *model.User
	userDB, err := r.DBClient.User.Query().Where(user.ID(userId)).Only(context.Background())
	if err != nil {
		return userModel, err
	}
	return &model.User{
		ID:         userDB.ID,
		CreditCard: userDB.CreditCard,
	}, nil
}
