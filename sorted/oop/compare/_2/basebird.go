package main

import "fmt"

type BaseBird struct {
	Age int
}

func (bb *BaseBird) Cal() {
	bb.Add()
}
func (bb *BaseBird) Add() {
	fmt.Printf("BaseBird before add: Age=%d\n", bb.Age)
	bb.Age = bb.Age + 1
	fmt.Printf("BaseBird after add: Age=%d\n", bb.Age)
}

type DerivedBird struct {
	BaseBird
}

func (db *DerivedBird) Add() {
	fmt.Printf("DerivedBird before add: Age=%d\n", db.Age)
	db.Age = db.Age + 2
	fmt.Printf("DerivedBird after add: Age=%d\n", db.Age)
}

func main() {
	var b1 BaseBird
	var b2 DerivedBird

	b1 = BaseBird{Age: 1}
	b1.Cal()

	b2 = DerivedBird{BaseBird{1}}
	b2.Cal()
}

/**
BaseBird before add: Age=1
BaseBird after add: Age=2
BaseBird before add: Age=1
BaseBird after add: Age=2
*/
