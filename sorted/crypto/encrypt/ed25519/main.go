package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func main() {
	message := []byte("Hello World")

	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(publicKey)
	fmt.Println(privateKey)

	// Sign the message
	signature := ed25519.Sign(privateKey, message)
	fmt.Printf("Signature: %x\n", signature)
	fmt.Printf("Signature: %x\n", []byte("signature"))

	// Verify the signature
	valid := ed25519.Verify(publicKey, message, signature)
	fmt.Printf("Valid: %v\n", valid)
}
