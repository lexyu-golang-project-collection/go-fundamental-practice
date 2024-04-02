package main

import "fmt"

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

func main() {
	v1 := Vehicle{4, "Red"}

	fmt.Println(v1)
}
