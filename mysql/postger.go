package sql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func PostgreSQL_Test() {
	dsn := "host=localhost user=postgres password=1532 dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Successfully connected to the PostgreSQL database")

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 查询是否已经存在名为 "Yujun" 的用户
	/*var existingUser User
	result := db.Where("name = ?", "Yujun").First(&existingUser)

	if result.RowsAffected > 0 {
		fmt.Println("User already exists:", existingUser.Name)
	} else {
		// 如果用户不存在，则创建一个新的用户
		db.Create(&User{Name: "Yujun"})
		fmt.Println("New user created: Yujun")
	}
	*/
	var users []User
	db.Find(&users)
	fmt.Println("All users in the database:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}

}
