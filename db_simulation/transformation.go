package main

import (
	"fmt"

	"gorm.io/gorm"
)

func schema_show(db *gorm.DB, table_input string) {
	var tables []string
	db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables)
	fmt.Println("Tables in the database:", tables)
	fmt.Println("------------------***-----------------")

	db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name = ?", table_input).Scan(&schema)
	fmt.Printf("Schema of table %s:\n%s\n", tableName, schema)
}
