package types

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// User struct for GORM model
type User struct {
	gorm.Model
	Name  string
	Email string
}

// UserType for GraphQL schema
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Field for Name",
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "Field for Email",
			},
		},
	},
)
