package main

import "fmt"

func main() {
	p := []int{1, 2, 3}
	q := 0
	q, p[q] = 1, 100
	fmt.Println("p = ", p)
	fmt.Println("q = ", q)
}
