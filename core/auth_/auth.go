package auth

import (
	"base-project/core/database"
	"base-project/core/registry"
	"base-project/core/user/types"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

type AuthModule struct {
	registry.DefaultModule // Embedding default implementations
}

// Constants for JWT
const (
	SecretKey = "your_secret_key"
)

func (a *AuthModule) CreateMutation() *graphql.Object {
	fmt.Println("Creating AuthMutation")
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "AuthMutation",
		Fields: graphql.Fields{
			"register": &graphql.Field{
				Type: graphql.String, // Returning the message
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					fmt.Println("Resolving register")
					name, _ := params.Args["name"].(string)
					username, _ := params.Args["username"].(string)
					email, _ := params.Args["email"].(string)
					password, _ := params.Args["password"].(string)

					fmt.Println(name, username, email, password)
					// Check if user already exists
					existingUser := types.User{}
					database.DB.Where("username = ? OR email = ?", username, email).First(&existingUser)
					if existingUser.ID != 0 {
						return nil, fmt.Errorf("username or email already exists")
					}

					// Hash password
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
					if err != nil {
						return nil, err
					}

					// Create new user
					newUser := types.User{
						Name:         name,
						Username:     username,
						Email:        email,
						PasswordHash: string(hashedPassword),
					}
					result := database.DB.Create(&newUser)
					if result.Error != nil {
						return nil, result.Error
					}

					return "User registered successfully!", nil
				},
			},

			"login": &graphql.Field{
				Type: graphql.String, // Returning the token as a string
				Args: graphql.FieldConfigArgument{
					"login": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					login, _ := params.Args["login"].(string)
					password, _ := params.Args["password"].(string)

					if username, authenticated := Authenticate(login, password); authenticated {
						token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
							"username": username,
							"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
						})

						tokenString, err := token.SignedString([]byte(SecretKey))
						if err != nil {
							return nil, err
						}
						return tokenString, nil
					}
					return nil, fmt.Errorf("invalid credentials")
				},
			},
		},
	})
}

// CreateQuery creates the AuthQuery object

func (a *AuthModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "AuthQuery",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					tokenString, _ := params.Context.Value("token").(string)
					token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
						return []byte(SecretKey), nil
					})
					if err != nil {
						return nil, err
					}

					if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
						return claims["username"], nil
					}

					return nil, fmt.Errorf("invalid token")
				},
			},
		},
	})
}

// Authenticate checks if the user exists and the password is correct

func Authenticate(login, password string) (string, bool) {
	var u types.User
	// Check for user by username or email
	database.DB.Where("username = ? OR email = ?", login, login).First(&u)
	if u.ID == 0 {
		return "", false
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", false
	}

	// User is authenticated, return the username and true
	return u.Username, true
}
