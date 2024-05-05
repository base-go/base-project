module base-project

go 1.22

require github.com/graphql-go/handler v0.2.3

require github.com/mattn/go-sqlite3 v1.14.22 // indirect

require (
	github.com/graphql-go/graphql v0.8.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.10
)

// replace (
// 	github.com/base-project/app => ./app
// 	github.com/base-project/core => ./core
// )
