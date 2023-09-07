package oop

type Vehicle struct {
	Seats int
	Color string
}

type Car struct {
	Vehicle
}

type MotorCycle struct {
	Base Vehicle
}
