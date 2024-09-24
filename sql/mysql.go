package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySql_test() {
	// 配置 MySQL 连接信息
	// 格式: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:1532@tcp(127.0.0.1:3306)/file-store?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// 获取通用数据库对象 sql.DB，以便进行底层操作
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting sql.DB from GORM:", err)
		return
	}
	defer sqlDB.Close()

	// 尝试 ping 数据库
	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	fmt.Println("Successfully connected to MySQL database using GORM!")
}
