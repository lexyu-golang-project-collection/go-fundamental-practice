package main

import "fmt"

// An interface is a collection of methods
type Car interface {
	Drive() string
	Stop() string
}

type BMW struct {
	Model string
}

func (b BMW) Drive() string {
	return "Driving " + b.Model
}

func (b BMW) Stop() string {
	return "Stop " + b.Model
}

func main() {
	c1 := &BMW{"TestCar"}

	fmt.Println("Model =", c1.Model)
	fmt.Println(c1.Drive())
	fmt.Println(c1.Stop())
}
