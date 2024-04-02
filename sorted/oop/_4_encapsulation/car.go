package main

import "fmt"

// An interface is a collection of methods
type Car interface {
	Drive() string
	Stop() string
	GetID()
	GetName()
}

type BMW struct {
	id    int
	name  string
	Model string
}

func (b BMW) Drive() string {
	return "Driving " + b.Model
}

func (b BMW) Stop() string {
	return "Stop " + b.Model
}

func (b *BMW) GetID() int {
	return b.id
}

func (b *BMW) GetName() string {
	return b.name
}

func main() {
	b1 := &BMW{1, "Elec", "Test"}

	fmt.Printf("%+v\n", b1)

	fmt.Println(b1.Drive())
	fmt.Println(b1.Stop())
}
