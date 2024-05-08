package main

import (
	"fmt"
	"net/http"

	"base-project/app"
	"base-project/core/database"
	"base-project/core/graphiql"

	"github.com/base-go/handler"
)

func main() {
	database.InitDB() // Initialize the database

	graphqlSchema := app.InitApp() // Initialize all application modules and get the schema

	h := handler.New(&handler.Config{
		Schema:   graphqlSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.Handle("/graphiql/",
		http.StripPrefix(
			"/graphiql/", http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					graphiql.RenderGraphiQL(w, r)
				},
			),
		),
	)

	fmt.Println("Server is running on http://localhost:3232/graphql")

	if err := http.ListenAndServe(":3232", nil); err != nil {
		panic(err)
	}
}
