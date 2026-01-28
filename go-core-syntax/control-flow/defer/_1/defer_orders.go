package main

import "fmt"

func printString(str string) {
	fmt.Printf("%q \n", str)
}

func printInt(i int) {
	fmt.Printf("%d \n", i)
}

func printFloat(f float64) {
	fmt.Printf("%.2f \n", f)
}

func main() {
	printString("browser") // 1
	defer printInt(32)     // 7
	defer printFloat(0.24) // 6
	printString("chrome")  // 2
	printInt(90)           // 3
	defer printFloat(89)   // 5
	printInt(900)          // 4
}
