package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

var (
	// 32 byte long secret key
	secretKey32 string = "uuYdfqwMd875eAoLMzrVzi4nwHYX9iQX"
	secretKey16 string = "gEvTbr0amBIucnx6"
)

func encrypt(plain_text string) string {
	block, err := aes.NewCipher([]byte(secretKey32))
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(block)

	// Make a buffer the same length as plaintext
	cipher_text := make([]byte, len(plain_text))
	block.Encrypt(cipher_text, []byte(plain_text))

	return string(cipher_text)
}

func main() {

	cipher := encrypt("This is some sensitive information")
	fmt.Printf("%x \n", cipher)

	// This will cause an error since the, plaintext is less than 16 bytes.
	// cipher = encrypt("Hi")
	// fmt.Printf("%x \n", cipher)
}
