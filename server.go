package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fracartdev/samplecrud/graph"
	"github.com/fracartdev/samplecrud/graph/generated"
	"github.com/fracartdev/samplecrud/postgres"
)

const defaultPort = "8080"

var dbUserName = "postgres"
var dbPassword = "root"
var dbHost = "127.0.0.1"
var dbName = "samplecrud"
var dbPort = 5432

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	booksService := &postgres.BooksService{
		DbUserName: dbUserName,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbName:     dbName,
		DbPort:     dbPort,
	}

	err := booksService.Init()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Books: booksService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
