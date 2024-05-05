package post

import (
	"base-project/app/post/mutations"
	"base-project/app/post/queries"
	"base-project/app/post/types"
	"base-project/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

func InitPostModule(schema *graphql.Schema) {
	// Initialize the post module
	fmt.Println("Initializing post module")
	Migrate()
	CreateSchema()
	CreateQuery()
	CreateMutation()
	fmt.Println("Post module initialized")
}
func Migrate() {
	// Migrate the post module
	database.DB.AutoMigrate(&types.Post{})
}
func CreateSchema() graphql.Schema {
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    CreateQuery(),
			Mutation: CreateMutation(),
		},
	)
	return schema
}

func CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllPosts": queries.GetAllPostsField(),
				"getPostById": queries.GetPostByIDField(),
			},
		},
	)
}

func CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createPost": mutations.CreatePostField(),
				"updatePost": mutations.UpdatePostField(),
				"deletePost": mutations.DeletePostField(),
			},
		},
	)
}
