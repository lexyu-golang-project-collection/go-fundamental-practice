package main

import (
	"fmt"
)

func modify(s *string) {
	*s = "change the world"
}

func main() {

	i := 100
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	var n *int
	var s *string
	fmt.Println(n) // nil
	fmt.Println(s) // nil
	fmt.Println("s = ", &s)

	var ptr_i *int = &i
	ptr_i2 := &i
	fmt.Println("Pointer = ", ptr_i)
	fmt.Println("Pointer = ", ptr_i2)

	s1 := "Hello"
	fmt.Printf("%T %v \n", s1, s1)
	var b *string = &s1
	fmt.Println("b = ", b)
	var a = &s1
	fmt.Println("a = ", a)
	c := &s1
	fmt.Println("c = ", c)
	ps := &s1
	*ps = "World"
	fmt.Printf("%T %v \n", s1, s1)

	modify(&s1)
	fmt.Println(s1)
}
