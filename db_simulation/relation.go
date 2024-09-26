package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func show_all(db *gorm.DB) {
	reader := bufio.NewReader(os.Stdin)
	table_show(db)
	fmt.Println("---------------------------------->>")
	fmt.Print("请输入要查看的table: ")
	table_inpu1, _ := reader.ReadString('\n')
	table_inpu1 = strings.TrimSpace(table_inpu1)

	// **动态查询表中的所有记录
	var results []map[string]interface{}
	err := db.Table(table_inpu1).Find(&results).Error
	if err != nil {
		fmt.Printf("Error querying table %s: %v\n", table_inpu1, err)
		return
	}
	// 打印查询结果
	if len(results) == 0 {
		fmt.Printf("Table %s is empty or does not exist.\n", table_inpu1)
		return
	}
	fmt.Printf("Records in table %s:\n", table_inpu1)
	for _, row := range results {
		fmt.Printf("%s \n", row)
	}
}
