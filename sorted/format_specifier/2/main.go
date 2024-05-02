package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	person := Person{
		Name:    "John",
		Age:     30,
		Address: "123 Main St",
	}

	fmt.Printf("%%v: %v\n", person)
	fmt.Printf("%%+v: %+v\n", person)
	fmt.Printf("%%#v: %#v\n", person)
}
