package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	bytes := []byte{0x47, 0x6f, 0x70, 0x68, 0x65, 0x72}

	encoded := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("Base64 Encode String = ", encoded)

	stringConverted := string(bytes)
	fmt.Println("String Converted:", stringConverted)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Base64 Decode = ", decoded)
	fmt.Printf("Base64 Decode String = %s", []byte(decoded))
}
