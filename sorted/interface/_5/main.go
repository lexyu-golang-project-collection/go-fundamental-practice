package main

// GO111MODULE=off go test -bench .
type Doer interface {
	Do() uint
}

type Empty struct{}

//go:noinline
func (e Empty) Do() uint {
	return 0
}
