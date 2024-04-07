package main

import "fmt"

func main() {
	var stocks map[string]float64
	sym := "Tesla"
	price := stocks[sym]
	fmt.Printf("1. %s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok { // using comma ok to check if key exists
		fmt.Printf("2. %s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("2. %s not found\n", sym)
	}
}
