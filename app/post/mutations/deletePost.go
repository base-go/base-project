package mutations

import (
	"base-project/app/post/types"
	"base-project/core/database"
	"log"

	"github.com/graphql-go/graphql"
)

func DeletePost(id int) (string, error) {
	var post types.Post
	if err := database.DB.Delete(&post, id).Error; err != nil {
		log.Printf("Error deleting post: %v", err)
		return "", err
	}
	return "Post successfully deleted", nil
}
func DeleteField() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.String,
		Description: "Delete a post",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(int)
			return DeletePost(id)
		},
	}
}
