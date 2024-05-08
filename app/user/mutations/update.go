package mutations

import (
    "base-project/app/user/types"
    "base-project/core/database"
    "log"

    "github.com/graphql-go/graphql"
)

func UpdateUser(input types.UpdateUserInput) (*types.User, error) {
    var user types.User
    // Fetch the existing record to update
    if err := database.DB.First(&user, input.ID).Error; err != nil {
        log.Printf("Error finding user: %v", err)
        return nil, err
    }

    // Update only the fields that were actually provided
    if input.Name != nil {
        user.Name = *input.Name
    }
    if input.Email != nil {
        user.Email = *input.Email
    }

    // Save the updated record
    if err := database.DB.Save(&user).Error; err != nil {
        log.Printf("Error updating user: %v", err)
        return nil, err
    }

    return &user, nil
}

func UpdateUserField() *graphql.Field {
    return &graphql.Field{
        Type:        types.UserType,
        Description: "Update an existing user",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "UpdateUserInput",
                    Fields: graphql.InputObjectConfigFieldMap{
                        "id": &graphql.InputObjectFieldConfig{
                            Type: graphql.NewNonNull(graphql.Int), // Ensure this matches the ID type in your schema
                        },
                        "name": &graphql.InputObjectFieldConfig{
                            Type: graphql.String, // Use dynamic GraphQL types
                        },
                        "email": &graphql.InputObjectFieldConfig{
                            Type: graphql.String, // Use dynamic GraphQL types
                        },
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            inputMap := p.Args["input"].(map[string]interface{})
            updateInput := types.UpdateUserInput{
                ID:      inputMap["id"].(int), // Ensure casting matches the ID type
                Name: optionalString(inputMap["name"]),
                Email: optionalString(inputMap["email"]),
            }
            return UpdateUser(updateInput)
        },
    }
}

// This needs to be adjusted if you support more types than strings.

func optionalString(val interface{}) *string {
	if str, ok := val.(string); ok {
		return &str
	}
	return nil
}
