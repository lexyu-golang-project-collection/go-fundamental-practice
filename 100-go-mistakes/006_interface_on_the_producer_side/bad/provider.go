package main

import "fmt"

// Provider 定義了一個龐大的介面
type Database interface {
	Connect() error
	Query(sql string) ([]string, error)
	Insert(data string) error
	Update(data string) error
	Delete(id string) error
	Close() error
}

// Provider 的實作
type SimpleReader struct{}

func (s SimpleReader) Connect() error {
	return nil
}

func (s SimpleReader) Query(sql string) ([]string, error) {
	return []string{"data1", "data2"}, nil
}

func (s SimpleReader) Insert(data string) error {
	fmt.Println("Insert: not implemented")
	return nil
}

func (s SimpleReader) Update(data string) error {
	fmt.Println("Update: not implemented")
	return nil
}

func (s SimpleReader) Delete(id string) error {
	fmt.Println("Delete: not implemented")
	return nil
}

func (s SimpleReader) Close() error {
	return nil
}
