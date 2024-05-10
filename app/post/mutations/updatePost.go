package mutations

import (
	"base-project/app/post/types"
	"base-project/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func UpdatePost(input types.UpdatePostInput) (*types.Post, error) {
	var post types.Post
	if err := database.DB.First(&post, input.ID).Error; err != nil {
		log.Printf("Error finding post: %v", err)
		return nil, err
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Content != nil {
		post.Content = *input.Content
	}

	if err := database.DB.Save(&post).Error; err != nil {
		log.Printf("Error updating post: %v", err)
		return nil, err
	}

	return &post, nil
}

func UpdateField() *graphql.Field {
	return &graphql.Field{
		Type:        types.PostType,
		Description: "Update an existing post",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "UpdatePostInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"id": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"title": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"content": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
					},
				})),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			inputMap := p.Args["input"].(map[string]interface{})
			updateInput := types.UpdatePostInput{
				ID:      inputMap["id"].(int),
				Title:   optionalString(inputMap["title"]),
				Content: optionalString(inputMap["content"]),
			}
			return UpdatePost(updateInput)
		},
	}
}

func optionalString(val interface{}) *string {
	if str, ok := val.(string); ok {
		return &str
	}
	return nil
}
