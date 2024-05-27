package main

import "fmt"

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("Cat says Hisssss")
}

func (c Cat) HappySound() {
	fmt.Println("Cat says Purrrrr")
}

func main() {
	var kitty Animal = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()

	fmt.Println("Cats Name :", kitty2.Name())
}
