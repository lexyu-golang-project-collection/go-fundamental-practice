package oop

// An interface is a collection of methods

type Car interface {
	Drive() string
	Stop() string
}

type BMW struct {
	Model string
}

func (b BMW) Drive() string {
	return "Driving " + b.Model
}

func (b BMW) Stop() string {
	return "Stop " + b.Model
}
