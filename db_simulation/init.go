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
	// Automatically migrate Employee and Department
	db.AutoMigrate(&Employee{}, &Department{})

	// Insert initial data for Department
	initialDepartments := []Department{
		{Name: "HR"},
		{Name: "Engineering"},
		{Name: "Marketing"},
		{Name: "Sales"},
	}

	for _, department := range initialDepartments {
		var existing Department
		result := db.Where("name = ?", department.Name).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&department) // 如果不存在则插入
		}
	}

	// Insert initial data for Employee
	initialEmployees := []Employee{
		{Name: "Alice", DepartmentID: 1},   // HR
		{Name: "Bob", DepartmentID: 2},     // Engineering
		{Name: "Charlie", DepartmentID: 3}, // Marketing
		{Name: "David", DepartmentID: 4},   // Sales
	}

	for _, employee := range initialEmployees {
		var existing Employee
		result := db.Where("name = ?", employee.Name).First(&existing)
		if result.RowsAffected == 0 {
			db.Create(&employee) // 如果不存在则插入
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

func table_show(db *gorm.DB) {
	var tables []string
	db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables)
	fmt.Println("Tables in the database:", tables)
}
