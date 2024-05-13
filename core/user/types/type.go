package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model   // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Name         string
	Username     string
	Email        string
	Avatar       string
	PasswordHash string
	LastLogin    time.Time
	Provider     string
	ProviderID   string
	AccessToken  string
	RefreshToken string
}

// UserType for GraphQL schema
var UserType = graphql.NewObject(
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
			"passwordHash": &graphql.Field{
				Type: graphql.String,
			},
			"lastLogin": &graphql.Field{
				Type: graphql.DateTime,
			},
			"provider": &graphql.Field{
				Type: graphql.String,
			},
			"providerID": &graphql.Field{
				Type: graphql.String,
			},
			"accessToken": &graphql.Field{
				Type: graphql.String,
			},
			"refreshToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// UserInput for GraphQL mutations
type UserInput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Avatar       string    `json:"avatar"`
	PasswordHash string    `json:"passwordHash"`
	LastLogin    time.Time `json:"lastLogin"`
	Provider     string    `json:"provider"`
	ProviderID   string    `json:"providerID"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
}

// UpdateUserInput for partial updates
type UpdateUserInput struct {
	ID           int        `json:"id"`
	Name         *string    `json:"name"`
	Username     *string    `json:"username"`
	Email        *string    `json:"email"`
	Avatar       *string    `json:"avatar"`
	PasswordHash *string    `json:"passwordHash"`
	LastLogin    *time.Time `json:"lastLogin"`
	Provider     *string    `json:"provider"`
	ProviderID   *string    `json:"providerID"`
	AccessToken  *string    `json:"accessToken"`
	RefreshToken *string    `json:"refreshToken"`
}
