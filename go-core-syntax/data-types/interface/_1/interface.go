package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}
type Square struct {
	side float64
}

type Rect struct {
	length, breadth float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

func (r Rect) area() float64 {
	return r.length * r.breadth
}

func (r Rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := Rect{length: 3, breadth: 4}
	c := Square{side: 5}
	printData(r)
	printData(c)
}
