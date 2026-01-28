package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("值: %v , Type: %T\n", a, a)
}

func main() {

	show(1)
	show("abc")
	show(true)

	slice := []int{1, 2, 3}
	show(slice)
}
