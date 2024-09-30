package main

import "fmt"

func test() (x int64, y int64) {
	x = 3
	y = 6

	return
}

func main() {

	var x int64 = 2
	fmt.Printf("x %v \n", x)

	// Redeclaration does not introduce a new variable; it just assigns a new value to the original .
	x, y := test()
	if x > 0 {
		fmt.Printf("x %v y %v \n", x, y)
	}

	fmt.Printf("x %v \n", x)
}
