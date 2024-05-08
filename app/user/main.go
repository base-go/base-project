package user

import (
	"base-project/app/modules"
	"base-project/app/user/mutations"
	"base-project/app/user/queries"
	"base-project/app/user/types"
	"base-project/core/database"
	"fmt"

	"github.com/graphql-go/graphql"
)

type UserModule struct{}

func init() {
	fmt.Println("Registering user module")
	modules.RegisterModule("user", &UserModule{})
}
func (p *UserModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserQueries",
			Fields: graphql.Fields{
				"getAllUsers": queries.GetAllUsersField(),
				"getUserById": queries.GetUserByIDField(),
			},
		},
	)
}

func (p *UserModule) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UserMutations",
			Fields: graphql.Fields{
				"createUser": mutations.CreateUserField(),
				"updateUser": mutations.UpdateUserField(),
				"deleteUser": mutations.DeleteUserField(),
			},
		},
	)
}

func Migrate() {
	// Migrate the user database model
	fmt.Println("Migrating user model...")
	database.DB.AutoMigrate(&types.User{})
	fmt.Println("User model migration completed.")
}
