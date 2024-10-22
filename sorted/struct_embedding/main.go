package main

import "fmt"

// Define the Driver interface
type Driver interface {
	Drive()
}

// Define the Car struct which implements the Drive method
type Car struct {
	brand string
}

func (c Car) Drive() {
	fmt.Println("Driving a", c.brand)
}

// Now define a Truck struct that embeds Car
type Truck struct {
	Car   // Embedding Car
	load int
}

func main() {
	// Create an instance of Truck
	myTruck := Truck{
		Car:  Car{brand: "Ford"},
		load: 1000,
	}

	// You can call the Drive method directly on Truck because of embedding
	myTruck.Drive()

	// Truck satisfies the Driver interface because it embeds Car, which implements Drive
	var d Driver = myTruck
	d.Drive()
}
