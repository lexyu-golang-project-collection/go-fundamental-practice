package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	var c **int = &b
	var x int = *b
	var y int = a

	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("y = ", &y)

	var year int = 2021
	var p *int = &year

	fmt.Println("&year = ", &year)
	fmt.Println("p = ", p)

}
