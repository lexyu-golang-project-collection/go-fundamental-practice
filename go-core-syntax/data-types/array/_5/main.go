package main

import "fmt"

func main() {
	var a []int
	fmt.Printf("&a = %p\n", &a)

	b := []int{1, 2}
	fmt.Printf("&b = %p\n", &b)

	a = append(a, 100)
	a = append(a, b...) // append another array

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Printf("&a = %p\n", &a)
}
