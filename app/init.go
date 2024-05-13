package app

import (
	"base-project/app/post"
	"base-project/core/registry"
)

// Init registers all application-specific modules
func Init() {
	registry.RegisterModule("post", &post.PostModule{})
}
