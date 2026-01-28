package main

import "fmt"

type Person struct {
	Name string
}

type PersonI interface{}

func main() {

	var p1, p2 PersonI = &Person{"K"}, &Person{"K"}
	var p3, p4 PersonI = Person{"K"}, Person{"K"}

	fmt.Println(p1 == p2) // false, Memory Addr are different
	fmt.Println(p3 == p4) // true, Compare Value

}
