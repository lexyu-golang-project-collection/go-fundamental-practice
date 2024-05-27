package main

import (
	"fmt"
)

func modify(s []int) {
	s[0] = 10000
}

func main() {
	// Slice
	slice := []int{1, 2, 3}
	fmt.Println("slice = ", slice)

	modify(slice)
	fmt.Println("slice = ", slice)
}
