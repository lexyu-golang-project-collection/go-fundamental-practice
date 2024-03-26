package main

import "fmt"

type contact struct {
	fname string
	lname string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact %s %s at %s", b.contact.fname, b.contact.lname, b.name)
}

func main() {
	con := contact{"Lex", "Yu", "123-456-789"}
	bus := business{"ABC Enterprise", "567 North St.", con}
	bus.info()
}
