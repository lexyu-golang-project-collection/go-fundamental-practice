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

func (c Cat) Walk() { fmt.Println("貓走路") }

// func (c Cat) Quack() { fmt.Println("貓呱呱叫") }

// Function that accepts any Speaker (Duck Typing)
func introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

// ---------------------------------

type Walker interface {
	Walk()
	Quack()
}

type Duck struct{}

func (d Duck) Walk()  { fmt.Println("鴨子走路") }
func (d Duck) Quack() { fmt.Println("鴨子呱呱叫") }

type Robot struct{}

func (r Robot) Walk()  { fmt.Println("機器人走路") }
func (r Robot) Quack() { fmt.Println("機器人模仿呱呱叫") }

func MakeItWalkAndQuack(w Walker) {
	w.Walk()
	w.Quack()
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both Dog and Cat implement the Speak method, so they satisfy the Speaker interface
	introduce(dog)
	introduce(cat)

	// ---------------------------------

	duck := Duck{}
	robot := Robot{}

	MakeItWalkAndQuack(duck)
	MakeItWalkAndQuack(robot)

	// MakeItWalkAndQuack(cat) // not impl full Walker methods
}
