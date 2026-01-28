package main

import (
	"fmt"
)

func main() {
	a := 14 // 1110 in binary
	b := 6  // 0110 in binary

	// a &^= b is equivalent to a = a &^ b
	a &^= b // Result: a = 1110 &^ 0110 => 1000 (which is 8 in decimal)

	fmt.Println(a) // Output: 8

}
