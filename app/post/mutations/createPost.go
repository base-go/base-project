// File: base-project/app/post/mutations/createPost.go
package mutations

import (
	"github.com/base-project/app/post/types"
	"github.com/base-project/core/database"
)

func CreatePost(input types.PostInput) (*types.Post, error) {
	post := &types.Post{Title: input.Title, Body: input.Body}
	result := database.DB.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}
