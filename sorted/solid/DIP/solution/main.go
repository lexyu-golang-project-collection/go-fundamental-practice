package main

import "fmt"

// Interfaces
type Oven interface {
	Heat() string
}

type Ingredient interface {
	Mix() string
}

// Impl
type GasOven struct{}

func (g *GasOven) Heat() string {
	return "Heating with gas oven!"
}

type ElectricOven struct{}

func (e *ElectricOven) Heat() string {
	return "Heating with electric oven!"
}

type Flour struct{}

func (f *Flour) Mix() string {
	return "Mixing with flour!"
}

type Sugar struct{}

func (s *Sugar) Mix() string {
	return "Mixing sugar!"
}

type Butter struct{}

func (b *Butter) Mix() string {
	return "Adding butter!"
}

// DI
type Bakery struct {
	oven        Oven
	ingredients []Ingredient
}

func (b *Bakery) Bake() {
	fmt.Println(b.oven.Heat())
	for _, ingredient := range b.ingredients {
		fmt.Println(ingredient.Mix())
	}
	fmt.Println("Baking an awesome pastry!")
}

func main() {
	gasOven := &GasOven{}
	electricOven := &ElectricOven{}
	flour := &Flour{}
	sugar := &Sugar{}
	butter := &Butter{}

	bakery := &Bakery{oven: gasOven, ingredients: []Ingredient{flour, sugar}}
	bakery.Bake()

	fmt.Println("--------------------------------------------------------")

	bakery = &Bakery{oven: electricOven, ingredients: []Ingredient{sugar, butter}}
	bakery.Bake()
}
