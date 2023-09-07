package oop

import "fmt"

func InterfaceEg2() {
	doRequest(&ACaller{})
	doRequest(&BCaller{})
	// won't occur err
	doRequest(&CCaller{})
}

type Caller interface {
	call()
}

type ACaller struct{}

func (c *ACaller) call() {
	fmt.Println("Call API A")
}

type BCaller struct{}

func (c *BCaller) call() {
	fmt.Println("Call API B")
}

type CCaller struct{}

func (c *CCaller) call() {
	fmt.Println("Call API C")
}

func doRequest(c Caller) {
	c.call()
}
