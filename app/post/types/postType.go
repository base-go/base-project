package types

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// Post struct for GORM model
type Post struct {
	gorm.Model
	Title   string
	Content string
}

// PostType for GraphQL schema
var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
