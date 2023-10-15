package main

import "fmt"

// Declare the Expense struct here
type Expense struct {
	name   string
	amount float64
	date   string
}

// Implement the Total method to calculate the total amount spent
func Total(earr []Expense) (total float64) {
	for _, v := range earr {
		total += v.amount
	}

	return total
}

// Implement the getName method on the Expense struct here
func (e *Expense) getName() string {

	return e.name
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
