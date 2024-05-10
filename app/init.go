package app

import (
	"base-project/app/post"
	"base-project/core/registry"
	"fmt"

	"github.com/graphql-go/graphql"
)

func InitApp() *graphql.Schema {
	registry.RegisterModule("post", &post.PostModule{})

	queryFields := graphql.Fields{}
	mutationFields := graphql.Fields{}

	// Iterate once over all modules to configure both queries and mutations
	for name, module := range registry.GetAllModules() {
		fmt.Printf("Loading module: %s\n", name)
		fmt.Printf("Loading and migrating module: %s\n", name)
		if migrator, ok := module.(interface{ Migrate() error }); ok {
			if err := migrator.Migrate(); err != nil {
				fmt.Printf("Failed to migrate module %s: %v\n", name, err)
				panic(fmt.Sprintf("Migration failed for module %s: %v", name, err))
			}
		}
		// Handle the query part
		modQuery := module.CreateQuery()
		if modQuery != nil {
			queryFields[name] = &graphql.Field{
				Type: modQuery,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// these fields primarily organize nested queries
					return struct{}{}, nil
				},
			}
		} else {
			fmt.Printf("Query object for module %s is nil\n", name)
		}

		// Handle the mutation part
		modMutation := module.CreateMutation()
		if modMutation != nil {
			mutationFields[name] = &graphql.Field{Type: modMutation,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// these fields primarily organize nested mutations
					return struct{}{}, nil
				},
			}

		} else {
			fmt.Printf("Mutation object for module %s is nil\n", name)
		}
	}

	// Create the RootQuery object
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queryFields,
	})
	if len(queryFields) == 0 {
		fmt.Println("No query fields have been defined. Cannot create a valid RootQuery object.")
		panic("RootQuery fields must be an object with field names as keys.")
	}

	// Create the RootMutation object
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutationFields,
	})

	// Create and return the schema
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
