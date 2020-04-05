package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v9"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/dickywijayaa/orders-go-graphql/graph"
	"github.com/dickywijayaa/orders-go-graphql/graph/generated"
	"github.com/dickywijayaa/orders-go-graphql/postgres"
	"github.com/dickywijayaa/orders-go-graphql/repositories"
	oauth "github.com/dickywijayaa/orders-go-graphql/middleware"
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

	UserRepo := repositories.UserRepository{DB: db}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(oauth.OAuthMiddleware(UserRepo))

	db.AddQueryHook(postgres.DBLogger{})

	c := generated.Config{Resolvers: &graph.Resolver{
		UserRepo: UserRepo,
		OrderRepo: repositories.OrderRepository{DB: db},
		OrderDetailRepo: repositories.OrderDetailRepository{DB: db},
		UserAddressRepo: repositories.UserAddressRepository{DB: db},
		ProvinceRepo: repositories.ProvinceRepository{DB: db},
		ProductRepo: repositories.ProductRepository{DB: db},
	}}

	middleware := graph.Dataloader{
		UserRepo: UserRepo,
		OrderRepo: repositories.OrderRepository{DB: db},
		OrderDetailRepo: repositories.OrderDetailRepository{DB: db},
		ProvinceRepo: repositories.ProvinceRepository{DB: db},
	}

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", middleware.LoaderMiddleware(db, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
