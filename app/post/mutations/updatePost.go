// File: base-project/app/post/mutations/updatePost.go
package mutations

import (
	"time"

	"github.com/base-project/app/post/types"
)

// UpdatePost updates an existing post
func UpdatePost(id int, input types.PostInput) (*types.Post, error) {
	// TODO: Implement logic for updating a post
	// For now, return a placeholder post
	post := &types.Post{
		ID:        id,
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return post, nil
}
