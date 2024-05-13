package main

import (
	"log"
	"net/http"

	"base-project/app"
	"base-project/core"
	"base-project/core/database"
	"base-project/core/registry"

	"github.com/base-go/handler"
	"github.com/graphql-go/graphql"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	// Register all modules from app and core packages
	app.Init()
	core.Init()

	// Initialize GraphQL schema
	schema, err := initGraphQLSchema()
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Setup GraphQL HTTP handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Configure HTTP routes
	http.Handle("/graphql", h)
	log.Println("Server is running on http://localhost:8080/graphql")

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initGraphQLSchema() (graphql.Schema, error) {
	queryFields := graphql.Fields{}
	mutationFields := graphql.Fields{}

	// Configure both queries and mutations for all modules
	for name, module := range registry.GetAllModules() {
		if query := module.CreateQuery(); query != nil {
			queryFields[name] = &graphql.Field{Type: query}
		}
		if mutation := module.CreateMutation(); mutation != nil {
			mutationFields[name] = &graphql.Field{Type: mutation}
		}
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutationFields,
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
