package main

import (
	// eg2 "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-oop/eg2"
	// eg3i "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-oop/eg3/inheritance"
	// eg3p "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-oop/eg3/polymorphism"
	"fmt"

	eg4 "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-oop/eg4"
	eg5 "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-oop/eg5"
)

func main() {

	// eg2.InterfaceEg2()

	/*
		car := &eg3i.Car{
			Vehicle: eg3i.Vehicle{
				Seats: 6,
				Color: "Blue",
			},
		}

		fmt.Println(car.Seats)
		fmt.Println(car.Color)

		motorCycle := &eg3i.MotorCycle{
			Base: eg3i.Vehicle{
				Seats: 2,
				Color: "red",
			},
		}

		fmt.Println(motorCycle.Base.Seats)
		fmt.Println(motorCycle.Base.Color)

		cat := &eg3p.Cat{
			Name: "catty",
		}

		fmt.Println(cat.Name)
		fmt.Println(cat.Sound())
	*/

	// Compare eg4 with eg5
	// /*
	bb1 := eg4.BaseBird{
		Age: 1,
	}
	bb1.Add()

	db1 := eg4.DerivedBird{
		BaseBird: eg4.BaseBird{
			Age: 1,
		},
	}
	db1.Add()
	// */

	fmt.Println("-----------------------------------")

	bb2 := eg5.BaseBird{
		Age: 1,
	}
	bb2.Cal()

	db2 := eg5.DerivedBird{
		BaseBird: eg5.BaseBird{
			Age: 1,
		},
	}
	db2.Cal()
}
