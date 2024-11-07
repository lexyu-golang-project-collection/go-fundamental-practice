package main

import "fmt"

// Define an interface with a method `Speak`
type Speaker interface {
	Speak() string
}

// Define a struct for Dog
type Dog struct {
	Name string
}

// Dog implements the Speak method
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Define a struct for Cat
type Cat struct {
	Name string
}

// Cat implements the Speak method
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// Function that accepts any Speaker (Duck Typing)
func introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both Dog and Cat implement the Speak method, so they satisfy the Speaker interface
	introduce(dog)
	introduce(cat)
}
