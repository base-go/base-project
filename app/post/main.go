package post

import (
	"base-project/app/post/mutations"
	"base-project/app/post/queries"
	"base-project/app/post/types"
	"base-project/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

type PostModule struct{}

func init() {
	fmt.Println("Registering post module")
	//modules.RegisterModule("post", &postModule{})
}

func (p *PostModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PostQueries",
			Fields: graphql.Fields{
				"getAllPosts": queries.GetAllPostsField(),
				"getPostById": queries.GetPostByIDField(),
			},
		},
	)
}

func (p *PostModule) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PostMutations",
			Fields: graphql.Fields{
				"createPost": mutations.CreatePostField(),
				"updatePost": mutations.UpdatePostField(),
				"deletePost": mutations.DeletePostField(),
			},
		},
	)
}

func Migrate() {
	// Migrate the post database model
	fmt.Println("Migrating post model...")
	database.DB.AutoMigrate(&types.Post{})
	fmt.Println("Post model migration completed.")
}
