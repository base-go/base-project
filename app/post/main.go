package post

import (
	"github.com/base-go/baseql/graphql"
	"github.com/base-project/app/init"
)

func init() {
	init.RegisterModule(InitPostModule)
}

func InitPostModule(schema *graphql.Schema) {
	// Assuming InitSchema setups the schema for posts
	InitSchema(schema)
}
