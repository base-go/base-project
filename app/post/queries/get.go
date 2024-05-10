package queries

import (
	"base-project/app/post/types"
	"base-project/core/database"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

// All returns a GraphQL field configuration for getting all posts.
func All() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.PostType),
		Description: "Get all posts",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			fmt.Println("Executing All resolver...")
			var posts []types.Post
			if err := database.DB.Find(&posts).Error; err != nil {
				fmt.Printf("Error fetching posts: %v\n", err)
				return nil, err
			}
			fmt.Printf("Posts retrieved: %d\n", len(posts))
			for _, post := range posts {
				fmt.Printf("Post: %v\n", post)
			}
			return posts, nil
		},
	}
}

// GetAllPosts retrieves all posts
func ByID() *graphql.Field {
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
