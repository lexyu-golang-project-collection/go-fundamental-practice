package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Employee struct {
	name string
	age  int
}

func main() {
	employees := []Employee{
		Employee{name: "John", age: 50},
		Employee{name: "Jane", age: 22},
		Employee{name: "Doe", age: 70},
		Employee{name: "Alex", age: 35},
		Employee{name: "Peter", age: 40},
	}

	fmt.Println("Employees before sorting: ")
	fmt.Printf("%+v\n", employees)

	fmt.Println("Employees after sorting: ")
	fmt.Printf("%+v\n", sortEmployee(employees))

}

func sortEmployee(es []Employee) []Employee {
	sortEmployees := es

	slices.SortFunc(sortEmployees,
		func(a, b Employee) int {
			return cmp.Or(
				cmp.Compare(a.age, b.age),
				cmp.Compare(a.name, b.name),
			)
		})

	return sortEmployees
}
