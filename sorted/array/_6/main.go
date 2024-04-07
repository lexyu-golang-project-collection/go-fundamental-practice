package main

import "fmt"

func main() {
	var h []int
	j := [2]int{1, 2}

	h = append(h, j[:]...)

	fmt.Println("h = ", h)
	fmt.Println("j = ", j)

}
