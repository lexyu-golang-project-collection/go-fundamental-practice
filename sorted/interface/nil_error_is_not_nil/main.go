package main

import "fmt"

type MyErr struct{}

func (e *MyErr) Error() string {
	return "MyErr"
}

func main() {
	var p *MyErr = nil
	var err error = p

	fmt.Println(err == nil) // ???
	err = nil
}
