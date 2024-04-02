package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	name string
}

type Cat struct {
	Name string
}

func (d *Dog) Sound() string {
	return "Woof!"
}

func (c *Cat) Sound() string {
	return "Meow!"
}

func main() {
	d1 := Dog{"husky"}
	c1 := Cat{"persian"}

	fmt.Println(d1.Sound())
	fmt.Println(c1.Sound())
}
