package main

import "fmt"

type Transport interface {
	getName() string
}

// Parent
type Vehicle struct {
	name string
}

func (v *Vehicle) getName() string {
	return v.name
}

// Sub
type Car struct {
	Vehicle
	wheel int
	doors int
}

// Sub
type Motocycle struct {
	Vehicle
	wheel int
}

// Other
type Printer struct{}

func (p *Printer) printName(t Transport) {
	fmt.Println("Vehicle Name:", t.getName())
}

func main() {
	vehicle := &Vehicle{name: "Tesla"}

	car := &Car{
		Vehicle: Vehicle{name: "CarName"},
		wheel:   4,
		doors:   4,
	}

	motocycle := &Motocycle{
		Vehicle: Vehicle{name: "MotocycleName"},
		wheel:   2,
	}

	printer := Printer{}

	printer.printName(vehicle)
	printer.printName(car)
	printer.printName(motocycle)
}
