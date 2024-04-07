package main

import "fmt"

func main() {
	var f [][]int
	g := []int{1, 2}

	t := [][]int{{777}, {999}}
	fmt.Println("t = ", t)

	f = append(f, g)
	fmt.Println("f = ", f)
	f = append(f, t...)
	fmt.Println("f = ", f)

	t2 := [][]int{
		{1, 2},
		{777},
		{999},
	}
	fmt.Println("t2 = ", t2)
}
