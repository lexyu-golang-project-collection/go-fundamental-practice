package main

import "fmt"

func main() {
	res := Multiply(1, 2)
	fmt.Println(res)

	x := "!"
	fmt.Println("Hello World", x)

	result := Add(5, -3)
	fmt.Println("Result:", result)
}

func Multiply(a int, b int) int {
	return a * b
}

func Add(a, b int) int {
	return a + b
}
