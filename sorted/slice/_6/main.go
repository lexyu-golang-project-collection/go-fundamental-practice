package main

import "fmt"

func main() {
	var x []int
	var y []int

	x = append(x, 1, 2, 3, 4, 5)
	y = append(y, 6, 7, 8, 9, 100)

	copy(x[3:], y[4:len(y)])

	fmt.Println(x)
}
