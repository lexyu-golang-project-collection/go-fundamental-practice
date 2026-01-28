package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	person   Person // Named field
	jobTitle string
}

func main() {
	emp := Employee{
		person: Person{
			name: "Jane Smith",
			age:  25,
		},
		jobTitle: "Product Manager",
	}

	// Accessing fields through the named field
	fmt.Println(emp.person.name) // "Jane Smith"
	fmt.Println(emp.person.age)  // 25
	fmt.Println(emp.jobTitle)    // "Product Manager"
}
