package mutations

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"log"
	"time"

	"github.com/graphql-go/graphql"
)

func CreateUser(input types.UserInput) (*types.User, error) {
	user := &types.User{

		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: input.PasswordHash,
		LastLogin:    input.LastLogin,
		Provider:     input.Provider,
		ProviderID:   input.ProviderID,
		AccessToken:  input.AccessToken,
		RefreshToken: input.RefreshToken,
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
						"username": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"email": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"passwordHash": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"lastLogin": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime, // Adjusted to use dynamic type
						},
						"provider": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"providerID": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"accessToken": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
						"refreshToken": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Adjusted to use dynamic type
						},
					},
				})),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			input, _ := p.Args["input"].(map[string]interface{})
			userInput := types.UserInput{
				Username:     input["username"].(string),     // Adjusted for dynamic casting
				Email:        input["email"].(string),        // Adjusted for dynamic casting
				PasswordHash: input["passwordHash"].(string), // Adjusted for dynamic casting
				LastLogin:    input["lastLogin"].(time.Time), // Adjusted for dynamic casting
				Provider:     input["provider"].(string),     // Adjusted for dynamic casting
				ProviderID:   input["providerID"].(string),   // Adjusted for dynamic casting
				AccessToken:  input["accessToken"].(string),  // Adjusted for dynamic casting
				RefreshToken: input["refreshToken"].(string), // Adjusted for dynamic casting
			}
			return CreateUser(userInput)
		},
	}
}
