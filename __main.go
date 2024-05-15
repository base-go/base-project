package main

import (
	"base-project/app/post/mutations"
	"base-project/core/auth/mutations"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/base-go/handler"
	"github.com/graphql-go/graphql"
)

type User struct {
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"name":  &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})

func main() {
	fields := graphql.Fields{
		"register": mutations.RegisterField(),
	}

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: fields,
	})

	// Creating a dummy rootQuery for the schema, as the schema expects at least a stub
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"dummy": &graphql.Field{ // Dummy field to satisfy the schema requirement
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return "Query placeholder", nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Configure HTTP routes
	http.Handle("/graphql", h)
	fmt.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
