package main

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return "int"
}

func Do(i fmt.Stringer) {
	_ = i.String()
}

// GO111MODULE=off go build -gcflags='-l -m -m'
func main() {
	m := MyInt(1)
	var i fmt.Stringer = m
	Do(i)
}

/*
./main.go:11:9: parameter i leaks to {heap} with derefs=0:
./main.go:11:9:   flow: {heap} = i:
./main.go:11:9:     from i.String() (call parameter) at ./main.go:12:14
./main.go:11:9: leaking param: i
./main.go:18:23: m escapes to heap:
./main.go:18:23:   flow: i = &{storage for m}:
./main.go:18:23:     from m (spill) at ./main.go:18:23
./main.go:18:23:     from i := m (assign) at ./main.go:18:6
./main.go:18:23:   flow: {heap} = i:
./main.go:18:23:     from Do(i) (call parameter) at ./main.go:19:4
./main.go:18:23: m escapes to heap
*/
