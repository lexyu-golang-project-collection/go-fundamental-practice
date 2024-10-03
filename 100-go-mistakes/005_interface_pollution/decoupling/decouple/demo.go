package main

import "fmt"

type Customer struct {
	ID    string
	Name  string
	Email string
}

// Decouple

type CustomerService interface {
	StoreCustomer(Customer) error
}

type CustomerServiceImpl struct {
	storer CustomerService
}

func (cs CustomerServiceImpl) createNewCustomer(id, name, email string) error {
	customer := &Customer{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return cs.storer.StoreCustomer(*customer)
}

// 實現 CustomerService 的不同儲存方式
type SQLDatabase struct{}

func (db SQLDatabase) StoreCustomer(customer Customer) error {
	fmt.Printf("Storing customer to SQL database: %+v\n", customer)
	return nil
}

type NoSQLDatabase struct{}

func (db NoSQLDatabase) StoreCustomer(customer Customer) error {
	fmt.Printf("Storing customer to NoSQL database: %+v\n", customer)
	return nil
}

type APIService struct{}

func (api APIService) StoreCustomer(customer Customer) error {
	fmt.Printf("Storing customer via API: %+v\n", customer)
	return nil
}

func main() {
	sqlDB := SQLDatabase{}
	nosqlDB := NoSQLDatabase{}
	apiService := APIService{}

	decoupledService1 := CustomerServiceImpl{storer: sqlDB}
	decoupledService1.createNewCustomer("2", "Jane Doe", "jane@example.com")

	decoupledService2 := CustomerServiceImpl{storer: nosqlDB}
	decoupledService2.createNewCustomer("3", "Bob Smith", "bob@example.com")

	decoupledService3 := CustomerServiceImpl{storer: apiService}
	decoupledService3.createNewCustomer("4", "Alice Johnson", "alice@example.com")
}
