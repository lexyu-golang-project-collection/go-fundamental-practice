package main

import "fmt"

func main() {
	var dividend, divisor int
	fmt.Print("Enter the dividend: ")
	fmt.Scanln(&dividend)
	fmt.Print("Enter the divisor: ")
	fmt.Scanln(&divisor)

	quotient := divide(dividend, divisor)
	fmt.Println("Quotient:", quotient)
}

func divide(dividend, divisor int) int {
	quotient := 0
	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}
	return quotient
}
