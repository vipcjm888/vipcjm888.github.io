package main

import "fmt"

// Reverse函数接受一个T类型的切片，并返回一个T类型的切片，T是一个类型参数
func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	// 整数切片的反转
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始整数切片:", intSlice)
	fmt.Println("反转后的整数切片:", Reverse(intSlice))

	// 字符串切片的反转
	stringSlice := []string{"Go", "泛型", "示例"}
	fmt.Println("原始字符串切片:", stringSlice)
	fmt.Println("反转后的字符串切片:", Reverse(stringSlice))
}
