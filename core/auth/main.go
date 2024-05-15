package auth

import (
	"base-project/core/auth/mutations"
	"base-project/core/auth/queries"

	"github.com/graphql-go/graphql"
)

type AuthModule struct{}

func (p *AuthModule) CreateQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AuthQueries",
			Fields: graphql.Fields{
				"me": queries.Me(),
			},
		},
	)
}

func (p *AuthModule) CreateMutation() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "AuthMutations",
			Fields: graphql.Fields{
				"login":    mutations.LoginField(),
				"register": mutations.RegisterField(),
				//"forgot":   mutations.ForgotField(),
			},
		},
	)
}

func (u *AuthModule) Resolvable() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// Default resolve function or specific logic
		return struct{}{}, nil
	}
}

func (p *AuthModule) Migrate() error {
	// Migrate the user database model
	return nil
}
