package main

import "fmt"

func main() {
	// %x : Base 16

	byteSlice := []byte{0x47, 0x6f, 0x70, 0x68, 0x65, 0x72}
	for _, b := range byteSlice {
		fmt.Printf("%x ", b)
	}
	fmt.Println()
}
