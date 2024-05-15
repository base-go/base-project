package mutations

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

func Register(input types.UserInput) (*types.User, error) {
	user := &types.User{
		Name:         input.Name,
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
func registerUserResolver(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context // Use GraphQL operation context
	fmt.Println("Resolving registration")
	log.Println(p.Args) // map[input:map[email:flakerimi@gmail.com name:Flakerim Ismani password:darude username:flakerimi]]

	input, _ := p.Args["input"].(map[string]interface{})
	name := input["name"].(string)
	username, usernameOk := input["username"].(string)
	email, emailOk := input["email"].(string)
	password, passwordOk := input["password"].(string)

	fmt.Println(username, email, password)
	fmt.Println(usernameOk, emailOk, passwordOk)

	if !usernameOk || !emailOk || !passwordOk {
		return nil, fmt.Errorf("all fields must be provided and valid")
	}

	// Check if user already exists
	existingUser := types.User{}
	if err := database.DB.WithContext(ctx).Where("username = ? OR email = ?", username, email).First(&existingUser).Error; err == nil {
		if existingUser.ID != 0 {
			return nil, fmt.Errorf("username or email already exists")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create new user
	newUser := types.User{
		Name:         name,
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	if err := database.DB.WithContext(ctx).Create(&newUser).Error; err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}

	return newUser, nil // Consider returning the newUser object to confirm creation
}

func RegisterField() *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "Register a new user",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "RegisterUserInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"name":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
						"username": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
						"email":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
						"password": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
					},
				})),
			},
		},
		Resolve: registerUserResolver,
	}
}
