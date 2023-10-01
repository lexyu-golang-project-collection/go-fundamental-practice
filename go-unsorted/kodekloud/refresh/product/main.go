package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	// your code goes here
	var discount float64
	switch product {
	case "apples":
		// price = price * 0.9
		discount = 0.1
	case "bananas":
		// price = price * 0.8
		discount = 0.2
	default:
		discount = 0
	}
	return price * (1 - discount)
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
