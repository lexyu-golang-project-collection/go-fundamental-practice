package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

type Rectangle struct {
	length  int
	breadth int
}

type Employee struct {
	eid int
	id  int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func (e Employee) get_id() int {
	return e.eid + 10
}

func main() {

	s := Student{name: "test", grades: []int{90, 75, 80}}
	s.displayName()

	fmt.Printf("%.2f%%\n", s.calculatePercentage())

	r := Rectangle{breadth: 10, length: 5}
	fmt.Println(r.area())
	fmt.Println(r)

	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{i, i + 10}
		fmt.Println(employees[i])
	}

	employees2 := make([]Employee, 5)
	for i := range employees2 {
		employees2[i] = Employee{eid: i}
		employees2[i].id = employees2[i].get_id()
		fmt.Printf("%+v\n", employees2[i])
	}
}
