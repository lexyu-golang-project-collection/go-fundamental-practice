package mypackage

import "fmt"

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}
func newCustAdd(c *customer, address string) {
	c.address = address
}

func Structs() {
	var ts customer
	ts.name = "Lex"
	ts.address = "5 main st"
	ts.bal = 234.56

	getCustInfo(ts)
	pl("Address :", ts.address)
	newCustAdd(&ts, "123 South st")
	pl("Address :", ts.address)
	ss := customer{"Sally Smith", "123 Main", 0.0}
	pl(ss.name)

	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())
}
