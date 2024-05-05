package app

import (
	"base-project/app/post"
	"base-project/app/post/types"
	"base-project/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

// InitApp initializes all application modules and returns a GraphQL schema.
func InitApp() *graphql.Schema {

	fmt.Println("Initializing app")
	schemaConfig := graphql.SchemaConfig{
		Query:    post.CreateQuery(),
		Mutation: post.CreateMutation(),
	}

	// create database connection and initialize tables
	database.InitDB()
	database.DB.AutoMigrate(&types.Post{})

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
		panic(err)
	}

	return &schema
}
