package queries

import (
    "base-project/app/user/types"
    "base-project/core/database"
    "errors"

    "github.com/graphql-go/graphql"
)

func GetAllUsersField() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.NewList(types.UserType),
        Description: "Get all users",
        Resolve: func(params graphql.ResolveParams) (interface{}, error) {
            var users []types.User
            if err := database.DB.Find(&users).Error; err != nil {
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return users, nil
        },
    }
}

func GetUserByIDField() *graphql.Field {
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
                return nil, err // Consider wrapping with a more user-friendly message
            }
            return &user, nil
        },
    }
}
