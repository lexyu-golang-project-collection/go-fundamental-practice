package main

import "fmt"

func main() {
	var d []int
	e := []int{1, 2, 3, 4, 5}

	fmt.Println("e = ", e)

	// [:1] new slice without values after index 1
	// [2:] include index 2
	d = append(e[:1], e[2:]...)

	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

}
