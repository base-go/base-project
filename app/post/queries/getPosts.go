package queries

import (
	"base-project/app/post/types"
	"base-project/core/database"
	"errors"

	"github.com/graphql-go/graphql"
)

// GetAllPostsField returns a GraphQL field configuration for getting all posts.
func GetAllPostsField() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.PostType),
		Description: "Get all posts",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var posts []types.Post
			if err := database.DB.Find(&posts).Error; err != nil {
				return nil, err
			}
			return posts, nil
		},
	}
}

// GetAllPosts retrieves all posts
func GetPostByIDField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PostType,
		Description: "Get post by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, ok := params.Args["id"].(int)
			if !ok {
				return nil, errors.New("id must be an integer")
			}
			var post types.Post
			if err := database.DB.First(&post, id).Error; err != nil {
				return nil, err
			}
			return &post, nil
		},
	}
}
