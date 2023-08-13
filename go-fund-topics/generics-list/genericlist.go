package main

type GenericList[T comparable] struct {
	data []T
}

func NewGemericList[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("Given index is too high")
	}

	for it := 0; it < len(l.data); it++ {
		if i == it {
			return l.data[it]
		}
	}

	panic("Value not found")
}

func (l *GenericList[T]) Remove(i int) {
	if i > len(l.data)-1 {
		panic("Given index is too high")
	} else if i < 0 {
		panic("Given index is too lower")
	}

	for it := 0; it < len(l.data); it++ {
		if it == i {
			l.data = append(l.data[:it], l.data[it+1:]...)
		}
	}
}

func (l *GenericList[T]) RemoveByValue(value T) {

	for i := 0; i < len(l.data); i++ {
		if value == l.data[i] {
			l.data = append(l.data[:i], l.data[i+1:]...)
		}
	}
}
