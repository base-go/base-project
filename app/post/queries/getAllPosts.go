package queries

import (
	"github.com/base-project/app/post/types"
	"github.com/base-project/core/database"
)

func GetAllPosts() ([]*types.Post, error) {
	var posts []*types.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
