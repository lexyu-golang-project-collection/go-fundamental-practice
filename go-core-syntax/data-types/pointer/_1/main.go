package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 100
	p := &a

	fmt.Println("p =", p)
	fmt.Println("*p =", *p)

	x := 10
	y := 50
	swap(&x, &y)
	fmt.Printf("x = %d, y = %d\n", x, y)

}
