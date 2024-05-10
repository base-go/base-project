package registry

import (
	"fmt"
	"sync"

	"github.com/graphql-go/graphql"
)

// GraphQLModule defines the interface that all modules must implement to be registered.
// Each module must provide methods to create its GraphQL queries and mutations.
type GraphQLModule interface {
	CreateQuery() *graphql.Object
	CreateMutation() *graphql.Object
	Migrate() error
}

var (
	// modulesRegistry stores all registered modules. The key is the module name.
	modulesRegistry = make(map[string]GraphQLModule)
	lock            sync.RWMutex
)

// RegisterModule attempts to register a module under a unique name. It returns an error
// if the module is already registered under that name.
func RegisterModule(name string, module GraphQLModule) error {
	lock.Lock()
	defer lock.Unlock()
	if _, exists := modulesRegistry[name]; exists {
		errMsg := fmt.Sprintf("Error: Module already registered: %s", name)
		fmt.Println(errMsg)
		return fmt.Errorf(errMsg)
	}
	modulesRegistry[name] = module
	fmt.Printf("Successfully registered module: %s\n", name)
	return nil
}

// GetAllModules retrieves a copy of the registry map, protecting it from modifications.
// It prints the current count of registered modules and each retrieval.
func GetAllModules() map[string]GraphQLModule {
	lock.RLock()
	defer lock.RUnlock()
	fmt.Println("Retrieving all modules...")
	copy := make(map[string]GraphQLModule, len(modulesRegistry))
	for key, value := range modulesRegistry {
		copy[key] = value
	}
	fmt.Printf("Retrieved all modules, count: %d\n", len(copy))
	return copy
}
