package main

import "fmt"

type Student struct {
	name string
}

// Class struct represents a class which aggregates students
type Class struct {
	students []Student
}

func main() {
	// Creating students
	student1 := Student{name: "Alice"}
	student2 := Student{name: "Bob"}

	// Creating a class and adding students to it
	class := Class{
		students: []Student{student1, student2},
	}

	// Displaying class students
	for _, student := range class.students {
		fmt.Println("Student:", student.name)
	}
}
