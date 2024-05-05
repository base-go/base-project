// File: base-project/app/post/queries/getPostById.go
package queries

import (
	"time"

	"github.com/base-project/app/post/types"
)

// GetPostByID retrieves a post by its ID
func GetPostByID(id int) (*types.Post, error) {
	// TODO: Implement logic for getting a post by ID
	// For now, return a placeholder post
	post := &types.Post{
		ID:        id,
		Title:     "Sample Post",
		Body:      "This is a sample post body",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return post, nil
}
