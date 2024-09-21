package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func cli_mydb(db *gorm.DB) {
	fmt.Println("|您现在位于simulation的主模块---------|")
	fmt.Println("|请输入一个数字来选择要运行的函数: |")
	fmt.Println("|1: 查看数据库'simulation'的tiger数据库的数据: |")
	fmt.Println("|112: 查看数据库里面所有的schema     |")
	fmt.Println("|111:Initialize the database  |")
	fmt.Println("|0: 退出                           |")
	fmt.Println("|----------------------------------|")

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
			fmt.Println("---------------------------------->>")
			fmt.Print("请输入要查看的table schema: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input) // 去除空白字符
			schema_show(db, input)
		case 3:
			fmt.Println("|--------------tbd------------|")
		case 112:
			table_show(db)
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
