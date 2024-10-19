package main

// GO111MODULE=off go test -bench .
type Empty struct{}

func (e Empty) Read(b []byte) (int, error) {
	return 0, nil
}
