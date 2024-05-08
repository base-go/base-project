package types

import "gorm.io/gorm"

// UserInput represents the input type for creating/updating a user

type UserInput struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UpdateUserInput represents the input type for updating a user

type UpdateUserInput struct {
	ID    int     `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
