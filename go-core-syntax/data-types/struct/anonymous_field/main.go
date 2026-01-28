package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person   // Anonymous field
	jobTitle string
}

func main() {
	emp := Employee{
		Person: Person{
			name: "John Doe",
			age:  30,
		},
		jobTitle: "Software Engineer",
	}

	// Accessing fields directly
	fmt.Println(emp.name)     // "John Doe"
	fmt.Println(emp.age)      // 30
	fmt.Println(emp.jobTitle) // "Software Engineer"
}
