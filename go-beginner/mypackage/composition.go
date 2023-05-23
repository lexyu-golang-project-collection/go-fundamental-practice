package mypackage

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
func Composition() {
	con1 := contact{"Lex", "Yu", "123-456-789"}
	bus1 := business{"LEX Corp", "567 North St.", con1}
	bus1.info()
}
