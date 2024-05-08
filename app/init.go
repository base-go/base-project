package app

import (
	"base-project/app/modules"
	"base-project/app/post"
	"base-project/app/user"
	"fmt"

	"github.com/graphql-go/graphql"
)

func InitApp() *graphql.Schema {

	modules.RegisterModule("post", &post.PostModule{})

	modules.RegisterModule("post", &user.UserModule{})

	queryFields := graphql.Fields{}

	for name, module := range modules.GetAllModules() {
		fmt.Printf("Loading module: %s\n", name)
		modQuery := module.CreateQuery()
		if modQuery == nil {
			fmt.Printf("Query object for module %s is nil\n", name)
			continue
		}
		queryFields[name] = &graphql.Field{Type: modQuery}
	}
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	})

	mutationFields := graphql.Fields{}

	for name, module := range modules.GetAllModules() {
		fmt.Printf("Loading module: %s\n", name)
		modMutation := module.CreateMutation()
		if modMutation == nil {
			fmt.Printf("Mutation object for module %s is nil\n", name)
			continue
		}
		mutationFields[name] = &graphql.Field{Type: modMutation}
	}
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutationFields,
	})

	if len(queryFields) == 0 {
		fmt.Println("No query fields have been defined. Cannot create a valid RootQuery object.")
		panic("RootQuery fields must be an object with field names as keys.")
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		fmt.Println("Failed to create GraphQL schema:", err)
		panic(err)
	}

	return &schema
}
