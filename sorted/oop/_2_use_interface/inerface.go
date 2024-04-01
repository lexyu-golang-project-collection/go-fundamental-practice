package main

import "fmt"

type A_Caller struct{}
type B_Caller struct{}
type C_Caller struct{}

type Caller interface {
	call()
}

func (c *A_Caller) call() {
	fmt.Println("Call API A")
}

func (c *B_Caller) call() {
	fmt.Println("Call API B")
}

func (c *C_Caller) call() {
	fmt.Println("Call API C")
}

func doRequest(c Caller) {
	c.call()
}

func CallAll() {
	doRequest(&A_Caller{})
	doRequest(&B_Caller{})
	// won't occur err
	doRequest(&C_Caller{})
}

func main() {
	CallAll()
}
