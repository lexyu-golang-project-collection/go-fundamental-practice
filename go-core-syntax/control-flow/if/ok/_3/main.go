package main

import "fmt"

func main() {
	var stocks map[string]float64

	// It's importatnt to initialize the map before adding values. Else it's a panic situation
	stocks = make(map[string]float64)
	stocks["APPL"] = 345.6

	fmt.Println(stocks)

	if price, ok := stocks["APPL"]; ok {
		fmt.Printf("3. %s -> $%.2f\n", "APPL", price)
	} else {
		fmt.Printf("3. %s not found\n", "APPL")
	}
}
