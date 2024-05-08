package mutations

import (
    "base-project/app/user/types"
    "base-project/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func DeleteUser(id int) (string, error) {
    var user types.User
    // Delete the record with the provided ID
    if err := database.DB.Delete(&user, id).Error; err != nil {
        log.Printf("Error deleting user with ID %d: %v", id, err)
        return "", err
    }
    return "User successfully deleted", nil
}

func DeleteUserField() *graphql.Field {
    return &graphql.Field{
        Type:        graphql.String,
        Description: "Delete a user",
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.Int),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            id := p.Args["id"].(int)
            // Call the delete function with the provided ID
            msg, err := DeleteUser(id)
            if err != nil {
                return nil, err
            }
            return msg, nil
        },
    }
}
