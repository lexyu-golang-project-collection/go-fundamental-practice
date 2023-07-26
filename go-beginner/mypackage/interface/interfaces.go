package mypackage

import g "github.com/lexyu/go-beginner/mypackage/global"

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	g.PL("cat Attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	g.PL("Cat says Hisssss")
}

func (c Cat) HappySound() {
	g.PL("Cat says Purrrrr")
}

func Interfaces() {
	var kitty Animal = Cat("Kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	g.PL("Cats Name :", kitty2.Name())
}
