package main

import "fmt"

func swap(x, y *int) {

}

func main() {
	x := 100
	p := &x

	fmt.Println("p =", p)
	fmt.Println("*p =", *p)
}
