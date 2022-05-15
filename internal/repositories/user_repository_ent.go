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

func (r *UserRepositoryEnt) CreateUser(item *model.NewUser) (*model.User, error) {
	userDB, err := r.DBClient.User.
		Create().
		SetName(item.Name).
		SetCreditCard(item.CreditCard).
		Save(context.Background())
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
		ID:   userDB.ID,
		Name: userDB.Name,
		CreditCard: userDB.CreditCard,
	}, nil
}

func (r *UserRepositoryEnt) GetAllUsers() ([]*model.User, error) {
	var items []*model.User
	usersDB, err := r.DBClient.User.Query().All(context.Background())
	if err != nil {
		return items, err
	}
	for _, userDB := range usersDB {
		var userModel *model.User
		var todoItems []*model.Todo
		todosDB, err := userDB.QueryTodos().All(context.Background())
		if err != nil {
			return items, err
		}
		for _, todoDB := range todosDB {
			var todoModel *model.Todo
			var remindersItems []*model.Reminder
			remindersDB, err := todoDB.QueryReminders().All(context.Background())
			if err != nil {
				return items, err
			}
			for _, reminderDB := range remindersDB {
				var reminderModel *model.Reminder
				reminderModel = &model.Reminder{
					ID:   reminderDB.ID,
					Text: reminderDB.Text,
				}
				remindersItems = append(remindersItems, reminderModel)
			}
			todoModel = &model.Todo{
				ID:   todoDB.ID,
				Text: todoDB.Text,
				Reminders: remindersItems,
			}
			todoItems = append(todoItems, todoModel)
		}
		userModel = &model.User{
			ID:   userDB.ID,
			Name: userDB.Name,
			CreditCard: userDB.CreditCard,
			Todos: todoItems,
		}
		items = append(items, userModel)
	}
	return items, nil
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
