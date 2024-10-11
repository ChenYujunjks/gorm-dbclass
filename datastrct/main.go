package main

import (
	"fmt"
)

func main() {
	// 测试 maxInArray 函数
	fmt.Println("测试 maxInArray 函数:")
	array1 := []int{1, 5, 3, 9, 2}
	fmt.Println("数组中的最大值：", maxInArray(array1)) // 输出 9

	array2 := []int{-1, -50, -3, -9, -2}
	fmt.Println("数组中的最大值：", maxInArray(array2)) // 输出 -1

	array3 := []int{5}
	fmt.Println("数组中的最大值：", maxInArray(array3)) // 输出 5

	// 测试 average 函数
	fmt.Println("\n测试 average 函数:")
	slice1 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println("切片的平均值：", average(slice1)) // 输出 3.0

	slice2 := []float64{10.0, 20.0, 30.0}
	fmt.Println("切片的平均值：", average(slice2)) // 输出 20.0

	slice3 := []float64{0.0, 0.0}
	fmt.Println("切片的平均值：", average(slice3)) // 输出 0.0

	// 测试 reverseMap 函数
	fmt.Println("\n测试 reverseMap 函数:")
	myMap1 := map[string]int{"one": 1, "two": 2, "three": 3}
	reversed1 := reverseMap(myMap1)
	fmt.Println("键值交换后的map：", reversed1) // 输出 map[1:"one", 2:"two", 3:"three"]

	myMap2 := map[string]int{"apple": 10, "banana": 20}
	reversed2 := reverseMap(myMap2)
	fmt.Println("键值交换后的map：", reversed2) // 输出 map[10:"apple", 20:"banana"]

	myMap3 := map[string]int{"key1": 100}
	reversed3 := reverseMap(myMap3)
	fmt.Println("键值交换后的map：", reversed3) // 输出 map[100:"key1"]
}
