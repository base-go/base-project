package types

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// User struct
type AuthUser struct {
	gorm.Model // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
	Username   string
	Email      string
	Avatar     string
	Password   string
}

// UserType for GraphQL schema
var AuthUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"avatar": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// UserInput for GraphQL mutations
type AuthUserInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
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
	},
})
