package user

import (
	"base-project/core/database"
	"base-project/core/user/mutations"
	"base-project/core/user/queries"
	"base-project/core/user/types"
	"fmt"

	"github.com/graphql-go/graphql"
)

type UserModule struct{}

func (p *UserModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserQueries",
			Fields: graphql.Fields{
				"list": queries.All(),
				"show": queries.ByID(),
			},
		},
	)
}

func (p *UserModule) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserMutations",
			Fields: graphql.Fields{
				"create": mutations.CreateUserField(),
				"update": mutations.UpdateUserField(),
				"delete": mutations.DeleteUserField(),
			},
		},
	)
}

func (u *UserModule) Resolvable() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// Default resolve function or specific logic
		return struct{}{}, nil
	}
}

func (p *UserModule) Migrate() error {
	// Migrate the user database model
	fmt.Println("Migrating user model...")

	if err := database.DB.AutoMigrate(&types.User{}); err != nil {
		fmt.Println("User model migration failed:", err)
		return err
	}
	fmt.Println("User model migration completed.")
	return nil
}
