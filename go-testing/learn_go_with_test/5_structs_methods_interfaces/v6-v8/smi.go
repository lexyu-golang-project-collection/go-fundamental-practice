package structs_methods_interfaces

import "math"

// interfaces
type Shape interface {
	Area() float64
}

// concretes
type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r *Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
