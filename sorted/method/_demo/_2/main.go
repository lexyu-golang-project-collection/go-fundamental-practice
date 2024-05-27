package main

import "fmt"

type PhoneFeatures interface {
	Call()
}

type Phone struct {
	UserName string
}

func NewPhone(userName string) Phone {
	return Phone{UserName: userName}
}

// Fix Option 1 - constructor return pointer
func NewPhone2(userName string) *Phone {
	return &Phone{UserName: userName}
}

func (p *Phone) Call() {
	fmt.Println(p.UserName, "Call...")
}

// Fix Option 2 - Value Receiver
// func (p Phone) Call() {
// 	fmt.Println(p.UserName, " Call...")
// }

func main() {
	// p := &Phone{UserName: "userA"}

	p := NewPhone("BetaUser")
	pc := p

	p.UserName = "Tester"

	p.Call()
	pc.Call()

	p2 := NewPhone2("BetaUser2")
	p2c := p2

	p2.UserName = "Tester2"

	p2.Call()
	p2c.Call()
}
