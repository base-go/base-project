package mutations

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"log"
	"time"

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
	if input.Username != nil {
		user.Username = *input.Username
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.PasswordHash != nil {
		user.PasswordHash = *input.PasswordHash
	}
	if input.LastLogin != nil {
		user.LastLogin = *input.LastLogin
	}
	if input.Provider != nil {
		user.Provider = *input.Provider
	}
	if input.ProviderID != nil {
		user.ProviderID = *input.ProviderID
	}
	if input.AccessToken != nil {
		user.AccessToken = *input.AccessToken
	}
	if input.RefreshToken != nil {
		user.RefreshToken = *input.RefreshToken
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
						"username": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"email": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"passwordHash": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"lastLogin": &graphql.InputObjectFieldConfig{
							Type: graphql.DateTime, // Use dynamic GraphQL types
						},
						"provider": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"providerID": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"accessToke": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
						"refreshToken": &graphql.InputObjectFieldConfig{
							Type: graphql.String, // Use dynamic GraphQL types
						},
					},
				})),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			inputMap := p.Args["input"].(map[string]interface{})
			updateInput := types.UpdateUserInput{
				ID:           inputMap["id"].(int),
				Username:     optionalString(inputMap["username"]),
				Email:        optionalString(inputMap["email"]),
				PasswordHash: optionalString(inputMap["passwordHash"]),
				LastLogin:    optionalTime(inputMap["lastLogin"]),
				Provider:     optionalString(inputMap["provider"]),
				ProviderID:   optionalString(inputMap["providerID"]),
				AccessToken:  optionalString(inputMap["accessToke"]),
				RefreshToken: optionalString(inputMap["refreshToken"]),
			}
			return UpdateUser(updateInput)
		},
	}
}

// This needs to be adjusted if you support more types than strings.
func optionalTime(val interface{}) *time.Time {
	if t, ok := val.(time.Time); ok {
		return &t
	}
	return nil
}
func optionalString(val interface{}) *string {
	if str, ok := val.(string); ok {
		return &str
	}
	return nil
}
