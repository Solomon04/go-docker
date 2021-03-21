package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/solomon04/go-docker/graph"
	"github.com/solomon04/go-docker/graph/generated"
	"github.com/solomon04/go-docker/src/db"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize the Environment
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize database
	db.InitDB()

	// Init Mailer
	// https://github.com/sirupsen/logrus
	// https://github.com/getsentry/sentry-go
	// https://github.com/go-co-op/gocron
	// https://github.com/dmgk/faker

	// Initialize GraphQL Server
	port := os.Getenv("HTTP_PORT")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
