package main

import (
	"fmt"

	"gorm.io/gorm"
)

// Define the tigerman struct to represent the table
type tigerman struct {
	ID       uint `gorm:"primaryKey"`
	Building string
}

func Initialize_seed(db *gorm.DB) {
	// Automatically migrate the schema (create the table)
	db.AutoMigrate(&tigerman{})

	// Insert initial data
	initialBuildings := []tigerman{
		{Building: "Building A"},
		{Building: "Building B"},
		{Building: "Building C"},
		{Building: "Building D"},
	}

	for _, building := range initialBuildings {
		var existing tigerman
		result := db.Where("building = ?", building.Building).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&building) // 只有在数据库中不存在该记录时才插入
		}
	}
}

func read_tiger(db *gorm.DB) {
	fmt.Printf("选定表格里面的所有元素：--->> \n")

	var tigermen []tigerman
	db.Find(&tigermen)
	fmt.Println("Buildings in tiger table:")
	for _, t := range tigermen {
		fmt.Printf("ID: %d, Building: %s\n", t.ID, t.Building)
	}
}

func schema_show(db *gorm.DB) {
	var tables []string
	db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables)
	fmt.Println("Tables in the database:", tables)
	fmt.Println("------------------***-----------------")

	var schema string
	tableName := "tigermen" // 这里你可以替换成任意表名
	db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name = ?", tableName).Scan(&schema)
	fmt.Printf("Schema of table %s:\n%s\n", tableName, schema)
}
