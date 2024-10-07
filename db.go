package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConnection *gorm.DB
)

func initDatabase() {
	var err error
	DBConnection, err = gorm.Open("sqlite3", "/tmp/tasks.db")
	if err != nil {
		log.Fatal("Error: Failed to connect to database. ", err)
	}
	DBConnection.AutoMigrate(&Task{})
}
