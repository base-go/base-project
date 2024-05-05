// File: base-project/app/post/types/postType.go

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
