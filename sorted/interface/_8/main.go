package main

// GO111MODULE=off go test -bench .
type Empty struct{}

var bytes []byte

func (e Empty) Read(b []byte) (int, error) {
	bytes = b
	return 0, nil
}
