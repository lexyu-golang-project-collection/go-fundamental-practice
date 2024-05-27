package main

import "fmt"

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

func main() {
	var c customer
	c.name = "Lex"
	c.address = "5 main st"
	c.bal = 234.56

	getCustInfo(c)
	fmt.Println("Address :", c.address)

	newCustAdd(&c, "123 South st")
	fmt.Println("Address :", c.address)

	ss := customer{"Sally Smith", "123 Main", 0.0}
	fmt.Println(ss.name)
}
