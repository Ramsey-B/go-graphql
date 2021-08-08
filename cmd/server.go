package main

import (
	"fmt"
	"go-graphql/config"
	"go-graphql/graph"
	"go-graphql/graph/generated"
	"go-graphql/internal/db"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading config file: %w \n", err))
	}

	database, err := db.Init(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w \n", err))
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB:     database,
		Config: cfg,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := cfg.Server.Port
	log.Printf("connect to http://localhost:%v/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
