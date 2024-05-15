package queries

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"errors"

	"github.com/graphql-go/graphql"
)

// GetAllUsers retrieves all users
func Me() *graphql.Field {
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
