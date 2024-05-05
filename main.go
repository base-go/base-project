package main

import (
	"net/http"

	"github.com/base-go/baseql/graphql"
	"github.com/base-go/baseql/graphql/graphiql"
	"github.com/base-go/baseql/graphql/introspection"
	"github.com/base-project/app"
	"github.com/base-project/core/database"
)

func main() {
	database.InitDB()

	graphqlSchema := graphql.NewSchema() // Create a new GraphQL schema

	app.InitApp(graphqlSchema) // Initialize all application modules

	introspection.AddIntrospectionToSchema(graphqlSchema)

	http.Handle("/graphql", graphql.Handler(graphqlSchema))
	http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))
	if err := http.ListenAndServe(":3030", nil); err != nil {
		panic(err)
	}
}
