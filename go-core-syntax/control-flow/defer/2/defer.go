package main

import "fmt"

func demo() int {
	i := 0
	defer func() {
		fmt.Println("defer-1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer-2")
	}()
	return i
}

func main() {
	fmt.Println("return", demo())
}
