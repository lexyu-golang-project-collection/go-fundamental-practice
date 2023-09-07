package eg4

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
