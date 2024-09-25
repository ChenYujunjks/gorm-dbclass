package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func controller(db *gorm.DB) {
	fmt.Println("|您现在位于simulation的主模块-------|")
	fmt.Println("|请输入一个数字来选择要运行的函数: |")
	fmt.Println("|1: 查看数据库'simulation'的tiger数据库的数据: |")
	fmt.Println("|2: 查看table schema |")
	fmt.Println("|3: 查看一个table里面的所有元素 |")
	fmt.Println("|4: Week3 Relation Operation |")
	fmt.Println("|0: 退出                           |")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("---------------------------------->>")
		fmt.Print("请输入选项: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // 去除空白字符

		// 尝试将输入转换为整数
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("请输入有效的数字！")
			continue
		}
		switch choice {
		case 1:
			//Initialize_seed(db)
			read_tiger(db)
		case 2:
			var schema1 string
			table_show(db)
			fmt.Println("---------------------------------->>")
			fmt.Print("请输入要查看的table schema: ")
			table_input, _ := reader.ReadString('\n')
			table_input = strings.TrimSpace(table_input) // 去除空白字符
			db.Raw("SELECT sql FROM sqlite_master WHERE type='table' AND name = ?", table_input).Scan(&schema1)
			fmt.Printf("Schema of table %s:\n%s\n", table_input, schema1)
		case 3:
			show_all(db)
		case 4:
			relation_join(db)
		case 111:
			Initialize_seed(db)
		case 0:
			fmt.Println("|--------------退出程序------------|")
			return
		default:
			fmt.Println("无效输入，请重新输入")
		}
	}
}
