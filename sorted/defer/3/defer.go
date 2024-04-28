package main

import "fmt"

func demo() (i int) {
	i = 0
	defer func() {
		i += 10
		fmt.Println("defer-2")
	}()
	return i
}

func main() {
	fmt.Println("return", demo())
}
