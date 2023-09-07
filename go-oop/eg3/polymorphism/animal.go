package oop

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
