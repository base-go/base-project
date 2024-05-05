package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/base.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func CloseDB() {
	db, _ := DB.DB()
	db.Close()
}
