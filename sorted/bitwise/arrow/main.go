package main

import (
	"fmt"
)

func main() {
	num := 2 // 二進位: 10
	fmt.Printf("%b\n", num)

	result := num << 2 // 將位元向左移 2 位 (相當於 2 * 2^2 = 8)

	fmt.Println(result) // 輸出: 8 (二進位: 1000)
	fmt.Printf("%b\n", result)

	num = 8           // 二進位: 1000
	result = num >> 2 // 將位元向右移 2 位 (相當於 8 / 2^2 = 2)

	fmt.Println(result) // 輸出: 2 (二進位: 10)
}
