package oop

type Car struct {
	id   int
	name string
}

func (c *Car) GetID() int {
	return c.id
}

func (c *Car) GetName() string {
	return c.name
}
