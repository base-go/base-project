// File: base-project/app/post/types/inputTypes.go
package types

// PostInput represents the input type for creating/updating a post
type PostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
