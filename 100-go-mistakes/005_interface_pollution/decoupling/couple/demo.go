package main

import "fmt"

type Customer struct {
	ID    string
	Name  string
	Email string
}

// Couple

type Database struct{}

func (db *Database) StoreCustomer(customer Customer) error {
	fmt.Printf("Storing customer to database: %+v\n", customer)
	return nil
}

type CustomerService struct {
	db Database
}

func (cs CustomerService) createNewCustomer(id, name, email string) error {
	c := &Customer{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return cs.db.StoreCustomer(*c)
}

func main() {
	coupledService := CustomerService{db: Database{}}
	coupledService.createNewCustomer("1", "John Doe", "john@example.com")
}
