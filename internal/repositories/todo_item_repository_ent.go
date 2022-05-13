package repositories

import (
	"context"
	"graphql-app/ent"
	"graphql-app/ent/todo"
	"graphql-app/graph/model"

	_ "google.golang.org/genproto/googleapis/rpc/status"
)

type TodoItemRepositoryEnt struct {
	DBClient *ent.Client
}

func NewTodoItemRepositoryEnt(DBClient *ent.Client) *TodoItemRepositoryEnt {
	return &TodoItemRepositoryEnt{DBClient: DBClient}
}

func (r *TodoItemRepositoryEnt) CreateTodo(item *model.NewTodo) (*model.Todo, error) {
	todoItemDB, err := r.DBClient.Todo.
		Create().
		SetText(item.Text).
		SetDone(false).
		Save(context.Background())
	if err != nil {
		return &model.Todo{}, err
	}
	return &model.Todo{
		ID:   todoItemDB.ID,
		Text: todoItemDB.Text,
		Done: todoItemDB.Done,
	}, nil
}

func (r *TodoItemRepositoryEnt) GetAllTodos() ([]*model.Todo, error) {
	var items []*model.Todo
	todoItemsDB, err := r.DBClient.Todo.Query().All(context.Background())
	if err != nil {
		return items, err
	}
	for _, todoItemDB := range todoItemsDB {
		var todoModel *model.Todo
		todoModel = &model.Todo{
			ID:   todoItemDB.ID,
			Text: todoItemDB.Text,
			Done: todoItemDB.Done,
		}
		items = append(items, todoModel)
	}
	return items, nil
}

func (r *TodoItemRepositoryEnt) GetTodoById(itemId int) (*model.Todo, error) {
	var todoModel *model.Todo
	todoItemDB, err := r.DBClient.Todo.Query().Where(todo.ID(itemId)).Only(context.Background())
	if err != nil {
		return todoModel, err
	}
	return &model.Todo{
		ID:   todoItemDB.ID,
		Text: todoItemDB.Text,
		Done: todoItemDB.Done,
	}, nil
}
