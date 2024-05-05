package app

import (
	"github.com/base-go/baseql/graphql"
	"github.com/base-project/app/post"
	// Include other modules here as you create them
)

// InitApp initializes all the application modules
func InitApp(schema *graphql.Schema) {
	post.InitPostModule(schema) // Initialize the post module
	// Call other module init functions here
}
