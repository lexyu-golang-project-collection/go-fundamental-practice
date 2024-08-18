package structs_methods_interfaces

import "math"

// v2
type Rectangle struct {
	Width  float64
	Height float64
}

// v3
type Circle struct {
	Radius float64
}

// v5
type Triangle struct {
	Base   float64
	Height float64
}

// v1
/*
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}
*/

// v2
func Perimeter(r *Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// func Area(r *Rectangle) float64 {
// 	return r.Width * r.Height
// }

// v4
type Shape interface {
	Area() float64
}

// v3
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// v5
func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
