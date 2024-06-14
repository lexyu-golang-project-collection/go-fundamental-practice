package main

import (
	"crypto/aes"
	"fmt"
)

var (
	// 32 bytes long secret key
	secretKey32 string = "uuYdfqwMd875eAoLMzrVzi4nwHYX9iQX"
	secretKey16 string = "gEvTbr0amBIucnx6"
)

func encrypt(plain_text string) string {
	key := []byte(secretKey32)

	block, _ := aes.NewCipher(key)
	fmt.Println("block = ", block)

	fmt.Println("aes.BlockSize = ", aes.BlockSize)

	// cipher_text := make([]byte, aes.BlockSize+len(plain_text))

	// cipher.NewCFBEncrypter()

	return string("")
}

func main() {
	encrypt("")
}
