package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v9"
	"github.com/dickywijayaa/orders-go-graphql/graph"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
	"github.com/dickywijayaa/orders-go-graphql/postgres"
	"github.com/dickywijayaa/orders-go-graphql/repositories"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := postgres.New(&pg.Options{
		User: "dickywijaya",
		Database: "orders_dev",
	})

	defer db.Close()

	db.AddQueryHook(postgres.DBLogger{})

	c := generated.Config{Resolvers: &graph.Resolver{
		UserRepo: repositories.UserRepository{DB: db},
		OrderRepo: repositories.OrderRepository{DB: db},
		OrderDetailRepo: repositories.OrderDetailRepository{DB: db},
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
