package main

import "fmt"

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func main() {

	rect1 := rectangle{10.0, 15.0}
	fmt.Println("Rect Area :", rect1.Area())
}
