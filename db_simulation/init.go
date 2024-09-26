package main

import (
	"fmt"

	"gorm.io/gorm"
)

// Define the tigerman struct to represent the table
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
			db.Create(&building) // Only insert if the record does not exist
		}
	}
	// relational database data
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
			db.Create(&employee)
		}
	}

	db.AutoMigrate(&Customer{}, &CustomerPhone{})

	// Insert initial data for Customer
	db.Create(&Customer{
		Email:   "alice@example.com",
		Name:    "Alice Wang",
		Street:  "Main St",
		ZipCode: "10001",
	})

	// 插入 CustomerPhone 数据
	db.Create(&CustomerPhone{
		CustomerEmail: "alice@example.com",
		Phone:         "123-456-7890",
	})

	db.Create(&CustomerPhone{
		CustomerEmail: "alice@example.com",
		Phone:         "098-765-4321",
	})
	// 新增 Customer 和 CustomerPhone 示例
	// Insert a new customer Bob Li and his phone numbers
	db.Create(&Customer{
		Email:   "bob@example.com",
		Name:    "Bob Li",
		Street:  "Market St",
		ZipCode: "94105",
	})

	db.Create(&CustomerPhone{
		CustomerEmail: "bob@example.com",
		Phone:         "555-555-5555",
	})

	db.Create(&CustomerPhone{
		CustomerEmail: "bob@example.com",
		Phone:         "555-123-4567",
	})
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
