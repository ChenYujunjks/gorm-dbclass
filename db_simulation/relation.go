package main

import (
	"fmt"

	"gorm.io/gorm"
)

var results []struct {
	EmployeeName   string
	DepartmentName string
}

func relation_join(db *gorm.DB) {
	db.Table("employees").Select("employees.name as EmployeeName, departments.name as DepartmentName").
		Joins("JOIN departments ON employees.department_id = departments.id").
		Scan(&results)

	// 打印查询结果
	for _, result := range results {
		fmt.Printf("Employee: %s, Department: %s\n", result.EmployeeName, result.DepartmentName)
	}
}
