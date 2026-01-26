package main

import "fmt"

func main() {
	res := Multiply(1, 2)
	fmt.Println(res)
}

func Multiply(a int, b int) int {
	result := 0
	for i := 0; i < a; i++ {
		result = result + b
	}
	return result // Test Agent Review Trigger Or Not
}
