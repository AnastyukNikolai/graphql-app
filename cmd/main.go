package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"graphql-app/ent"
	"graphql-app/graph"
	"graphql-app/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8989"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

	client := OpenDB()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logrus.Fatal(err)
	}
	er := client.Close()
	if er != nil {
		logrus.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func OpenDB() *ent.Client {
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

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}
