package mutations

import (
	"base-project/app/post/types"
	"base-project/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func CreatePost(input types.PostInput) (*types.Post, error) {
	post := &types.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	if err := database.DB.Create(post).Error; err != nil {
		log.Printf("Error creating post: %v", err)
		return nil, err
	}
	return post, nil
}

// CreatePostField returns a GraphQL field configuration for creating a post.
func CreateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PostType,
		Description: "Create a new post",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "PostInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"title": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"content": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				})),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			input, _ := p.Args["input"].(map[string]interface{})
			postInput := types.PostInput{
				Title:   input["title"].(string),
				Content: input["content"].(string),
			}
			return CreatePost(postInput)
		},
	}
}
