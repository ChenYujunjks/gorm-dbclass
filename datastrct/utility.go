package main

// 实现 maxInArray 函数
func maxInArray(arr []int) int {
	out := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > out {
			out = arr[i]
		}
	}
	return out
}

// 实现 average 函数
func average(slice []float64) float64 {
	if len(slice) == 0 {
		return 0.0 // 如果切片为空，返回 0.0 或者根据需求返回其他值
	}
	sum := 0.0
	for _, value := range slice {
		sum += value
	}
	return sum / float64(len(slice))
}

// 实现 reverseMap 函数
func reverseMap(m map[string]int) map[int]string {
	reversed := make(map[int]string)

	for key, value := range m {
		reversed[value] = key
	}
	return reversed
}
