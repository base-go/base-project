package mutations

import (
	"base-project/app/user/types"
	"base-project/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func CreateUser(input types.UserInput) (*types.User, error) {
	user := &types.User{
		
		Name: input.Name,
		Email: input.Email,
	}
	if err := database.DB.Create(user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}
// CreateUserField returns a GraphQL field configuration for creating a user.
func CreateUserField() *graphql.Field {
    return &graphql.Field{
        Type:        types.UserType,
        Description: "Create a new user",
        Args: graphql.FieldConfigArgument{
            "input": &graphql.ArgumentConfig{
                Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
                    Name: "UserInput",
                    Fields: graphql.InputObjectConfigFieldMap{
                        "name": &graphql.InputObjectFieldConfig{
                            Type: graphql.String, // Adjusted to use dynamic type
                        },
                        "email": &graphql.InputObjectFieldConfig{
                            Type: graphql.String, // Adjusted to use dynamic type
                        },
                    },
                })),
            },
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            input, _ := p.Args["input"].(map[string]interface{})
            userInput := types.UserInput{
                Name: input["name"].(string), // Adjusted for dynamic casting
                Email: input["email"].(string), // Adjusted for dynamic casting
            }
            return CreateUser(userInput)
        },
    }
}
