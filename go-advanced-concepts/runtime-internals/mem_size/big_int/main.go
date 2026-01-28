package main

import (
	"fmt"
	"math/big"
	"unsafe"
)

func main() {

	// Create a big.Int value
	x, _ := new(big.Int).SetString("123456789", 10)

	y := 13456789
	fmt.Println(unsafe.Sizeof(y))

	// Calculate the size in bytes
	numBits := x.BitLen()
	numBytes := (numBits + 7) / 8

	fmt.Println("Length : ", len(x.Text(10)))
	fmt.Printf("Value: %s\n", x.String())
	fmt.Printf("Size: %d bits (%d bytes)\n", numBits, numBytes)
}
