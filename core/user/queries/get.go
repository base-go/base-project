package queries

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

// All returns a GraphQL field configuration for getting all users.
func All() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.UserType),
		Description: "Get all users",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Executing All resolver...")
			var users []types.User
			if err := database.DB.Find(&users).Error; err != nil {
				fmt.Printf("Error fetching users: %v\n", err)
				return nil, err
			}
			fmt.Printf("Users retrieved: %d\n", len(users))
			for _, user := range users {
				fmt.Printf("User: %v\n", user)
			}
			return users, nil
		},
	}
}

// GetAllUsers retrieves all users
func ByID() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "Get user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(int)
			if !ok {
				return nil, errors.New("id must be an integer")
			}
			var user types.User
			if err := database.DB.First(&user, id).Error; err != nil {
				return nil, err
			}
			return &user, nil
		},
	}
}
