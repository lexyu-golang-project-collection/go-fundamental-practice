package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	fmt.Println("arr index 0 = ", arr[0])

	var c [10]int // default value 0
	fmt.Println("c = ", c)
}
