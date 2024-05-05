// File: base-project/app/post/types/inputTypes.go
package types

// PostInput represents the input type for creating/updating a post
type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostInput struct {
	ID      int     `json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}
