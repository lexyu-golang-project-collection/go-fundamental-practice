package main

import "fmt"

type BaseBird struct {
	Age int
}

func (this *BaseBird) Add() {
	fmt.Printf("BaseBird before add: Age=%d\n", this.Age)
	this.Age = this.Age + 1
	fmt.Printf("BaseBird after add: Age=%d\n", this.Age)
}

type DerivedBird struct {
	BaseBird
}

func (this *DerivedBird) Add() {
	fmt.Printf("DerivedBird before add: Age=%d\n", this.Age)
	this.Age = this.Age + 2
	fmt.Printf("DerivedBird after add: Age=%d\n", this.Age)
}

func main() {
	var b1 BaseBird
	var b2 DerivedBird

	b1 = BaseBird{Age: 1}
	b1.Add()

	b2 = DerivedBird{BaseBird{1}}
	b2.Add()
}

/**
BaseBird before add: Age=1
BaseBird after add: Age=2
DerivedBird before add: Age=1
DerivedBird after add: Age=3
*/
