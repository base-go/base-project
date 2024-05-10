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

func (p *PostModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PostQueries",
			Fields: graphql.Fields{
				"list": queries.All(),
				"show": queries.ByID(),
			},
		},
	)
}

func (p *PostModule) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PostMutations",
			Fields: graphql.Fields{
				"create": mutations.CreateField(),
				"update": mutations.UpdateField(),
				"delete": mutations.DeleteField(),
			},
		},
	)
}

func (p *PostModule) Migrate() error {
	// Migrate the post database model
	fmt.Println("Migrating post model...")
	database.DB.AutoMigrate(&types.Post{})
	if err := database.DB.AutoMigrate(&types.Post{}); err != nil {
		fmt.Println("Post model migration failed:", err)
		return err
	}
	fmt.Println("Post model migration completed.")
	return nil
}
