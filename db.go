package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConnection *gorm.DB
)

func initDatabase() {
	DBConnection, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		fmt.Printf("Error: Failed to connect to database. %v", DBConnection)
	}
	DBConnection.AutoMigrate(&task{})
}
