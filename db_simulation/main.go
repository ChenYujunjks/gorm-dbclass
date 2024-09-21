package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("simulation.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	cli_mydb(db)
}
