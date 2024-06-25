package main

import "fmt"

type Teacher struct {
	name string
}

type Student struct {
	name string
}

// Association between Teacher and Student
type Association struct {
	teacher Teacher
	student Student
}

func main() {
	// Creating a teacher and a student
	teacher := Teacher{name: "Mr. Smith"}
	student := Student{name: "Alice"}

	// Establishing an association between the teacher and the student
	association := Association{
		teacher: teacher,
		student: student,
	}

	// Displaying the association
	fmt.Printf("%s teaches %s\n", association.teacher.name, association.student.name)
}
