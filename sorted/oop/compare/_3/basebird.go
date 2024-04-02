package main

import "fmt"

type Bird interface {
	Add()
}

func Cal(bird Bird) {
	bird.Add()
}

type BaseBird struct {
	Age int
}

func (this *BaseBird) Add() {
	fmt.Printf("BaseBird before add: age=%d\n", this.Age)
	this.Age = this.Age + 1
	fmt.Printf("BaseBird after add: age=%d\n", this.Age)
}

type DerivedBird struct {
	BaseBird
}

func (this *DerivedBird) Add() {
	fmt.Printf("DerivedBird before add: age=%d\n", this.Age)
	this.Age = this.Age + 2
	fmt.Printf("DerivedBird after add: age=%d\n", this.Age)
}

func main() {
	var b1, b2 Bird
	b1 = &BaseBird{Age: 1}
	b2 = &DerivedBird{BaseBird{Age: 1}}
	Cal(b1)
	Cal(b2)
}

/**
BaseBird before add: age=1
BaseBird after add: age=2
DerivedBird before add: age=1
DerivedBird after add: age=3
*/
