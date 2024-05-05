// File: base-project/app/post/resolvers/mutationResolver.go
package resolvers

import (
	"base-project/app/post/mutations"
	"base-project/app/post/types"
)

// MutationResolver resolves mutations related to posts
type MutationResolver struct{}

// CreatePost resolves the mutation for creating a new post
func (r *MutationResolver) CreatePost(input types.PostInput) (*types.Post, error) {
	return mutations.CreatePost(input)
}

// UpdatePost resolves the mutation for updating an existing post
func (r *MutationResolver) UpdatePost(id int, input types.UpdatePostInput) (*types.Post, error) {
	return mutations.UpdatePost(input)
}

// DeletePost resolves the mutation for deleting a post
func (r *MutationResolver) DeletePost(id int) (string, error) {
	return mutations.DeletePost(id)
}
