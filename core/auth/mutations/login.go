package mutations

import (
	"base-project/core/database"
	"base-project/core/user/types"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// SecretKey
const SecretKey = "secret"

// Authenticate the user by checking credentials
func Authenticate(login, password string) (map[string]interface{}, error) {
	var user types.User
	result := database.DB.Where("username = ? OR email = ?", login, login).First(&user)
	if result.Error != nil {
		log.Printf("Authentication failed: %v", result.Error)
		return nil, result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		log.Println("Invalid password:", err)
		return nil, err
	}

	expTime := time.Now().Add(time.Hour * 24 * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      expTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("Token generation failed:", err)
		return nil, err
	}

	return map[string]interface{}{
		"accessToken": tokenString,
		"username":    user.Username,
		"id":          user.ID,
		"avatar":      user.Avatar,
		"email":       user.Email,
		"name":        user.Name,
		"exp":         expTime.Unix(),
	}, nil
}

var LoginResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LoginResponse",
	Fields: graphql.Fields{
		"accessToken": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"exp": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"avatar": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// LoginField returns a GraphQL field configuration for logging in a user.
func LoginField() *graphql.Field {
	return &graphql.Field{
		Type:        LoginResponseType, // Make sure this matches the type you defined.
		Description: "Log in a user using either username or email",
		Args: graphql.FieldConfigArgument{
			"login": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			login := p.Args["login"].(string)
			password := p.Args["password"].(string)
			return Authenticate(login, password)
		},
	}
}
