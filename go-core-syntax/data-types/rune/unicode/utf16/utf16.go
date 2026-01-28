package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	// UTF-16 編碼
	runes := []rune("Hello, 世界")
	encoded := utf16.Encode(runes)
	fmt.Printf("UTF-16 encoded: %v\n", encoded)

	// UTF-16 解碼
	decoded := utf16.Decode(encoded)
	fmt.Printf("Decoded string: %s\n", string(decoded))

	// 處理代理對
	r := rune(0x1F600) // 笑臉 emoji
	r1, r2 := utf16.EncodeRune(r)
	fmt.Printf("Surrogate pair: %X, %X\n", r1, r2)

	// 解碼代理對
	original := utf16.DecodeRune(r1, r2)
	fmt.Printf("Original rune: %X\n", original)
}
