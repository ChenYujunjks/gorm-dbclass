package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 20

	arr2 := [5]int{1, 3, 4, 5}
	fmt.Println(arr)
	//for loop array
	for i := 0; i < len(arr2); i++ {
		//fmt.Println(arr2[i])
	}
	// range loop array
	for i, value := range arr2 {
		fmt.Println(i, value)
		if i == 3 {
			fmt.Println("Break from range loop!----------------!")
			break
		}
	}
	slice := []int{3, 2, 1}
	slice2 := append(slice, 223)
	fmt.Println(slice2, "----------------------_>  map")

	m := make(map[string]int) // 声明并初始化一个空的map
	m["orange"] = 4           // 插入键值对 ("orange", 4)
	m["banana"] = 5           // 更新键 "banana" 的值为 5
	m["orange"] = 3           // 插入键值对 ("orange", 4)
	value := m["apple"]       // 获取键 "apple" 对应的值
	fmt.Println(value)        // 输出: 1
	delete(m, "cherry")
	delete(m, "orange")
	// 遍历map
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	a23 := 64.0
	fmt.Println(a23)
	//a11:=5
	//a11/a23
}
