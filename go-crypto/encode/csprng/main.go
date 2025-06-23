package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// CSPRNG :Cryptographically Secure Pseudo-Random Number Generator
func GenerateSecretHex40() (string, error) {
	b := make([]byte, 20) // 20 bytes = 160 bits
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil // 40 hex characters
}

func main() {
	str, _ := GenerateSecretHex40()
	fmt.Println(str)
}
