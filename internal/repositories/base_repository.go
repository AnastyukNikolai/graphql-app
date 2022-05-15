package repositories

import (
	"database/sql"
	"fmt"
	"graphql-app/ent"
	"graphql-app/graph/model"
	"os"

	"github.com/joho/godotenv"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Repository struct {
	User
	TodoItem
	Reminder
}

type User interface {
	GetUserById(userId int) (user *model.User, err error)
	GetAllUsers() ([]*model.User, error)
	CreateUser(item *model.NewUser) (*model.User, error)
}

type TodoItem interface {
	CreateTodo(item *model.NewTodo) (*model.Todo, error)
	GetAllTodos() ([]*model.Todo, error)
	GetTodoById(itemId int) (*model.Todo, error)
}

type Reminder interface {
	CreateReminder(item *model.NewReminder) (*model.Reminder, error)
	GetAllReminders() ([]*model.Reminder, error)
	GetReminderById(itemId int) (*model.Reminder, error)
}

func NewRepository(DBClient *ent.Client) *Repository {
	return &Repository{
		User:     NewUserRepositoryEnt(DBClient),
		TodoItem: NewTodoItemRepositoryEnt(DBClient),
		Reminder: NewReminderRepositoryEnt(DBClient),
	}
}

// Open new connection
func OpenDB() *ent.Client {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		viper.GetString("db.username"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
	)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
