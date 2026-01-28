package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	var h1 maphash.Hash
	h1.WriteString("Hi")
	fmt.Println("val=", h1.Sum64())

	var h2 maphash.Hash
	h2.WriteString("Hi")
	fmt.Println("val=", h2.Sum64())

	var h3 maphash.Hash
	h3.SetSeed(h1.Seed())
	h3.WriteString("Hi")
	fmt.Println("val=", h3.Sum64())
}
