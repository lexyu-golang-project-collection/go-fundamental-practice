package main

import "fmt"

func main() {
	res := Multiply(1, 2)
	fmt.Println(res)

	// 故意用未宣告的變數觸發警告
	fmt.Println("Hello World", x)

	// 一個簡單函式，有潛在錯誤
	result := Add(5, -3)
	fmt.Println("Result:", result)
}

func Multiply(a int, b int) int {
	return a * b
}

func Add(a, b int) int {
	return a + b // 沒有錯誤處理
}
