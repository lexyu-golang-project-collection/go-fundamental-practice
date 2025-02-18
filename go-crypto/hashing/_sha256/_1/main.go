package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "Hello, world!"

	hasher := sha256.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)

	fmt.Println(hasher)

	fmt.Println(hash)

	fmt.Println(hex.EncodeToString(hash))
}
