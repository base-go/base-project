package core

import (
	"base-project/core/auth"
	"base-project/core/registry"
	"base-project/core/user"
)

// Init registers all core modules
func Init() {
	registry.RegisterModule("user", &user.UserModule{})
	registry.RegisterModule("auth", &auth.AuthModule{})
}
