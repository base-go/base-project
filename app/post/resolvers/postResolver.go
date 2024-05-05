// File: base-project/app/post/resolvers/postResolver.go
package resolvers

import (
	"base-project/app/post/types"
	"base-project/core/database"

	"github.com/graphql-go/graphql"
)

// PostResolver resolves queries and mutations related to posts
type PostResolver struct{}

func GetAllPosts(params graphql.ResolveParams) (interface{}, error) {
	var posts []types.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
