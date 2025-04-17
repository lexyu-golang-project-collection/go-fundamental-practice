package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

var (
	secretKey32 string = "uuYdfqwMd875eAoLMzrVzi4nwHYX9iQX"
	secretKey16 string = "gEvTbr0amBIucnx6"
)

func encrypt(plainText string) []byte {
	block, err := aes.NewCipher([]byte(secretKey16))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block)

	// Make a buffer the same length as plaintext
	plainBytes := []byte(plainText)
	paddedBytes := pad(plainBytes, block.BlockSize())
	cipherText := make([]byte, len(paddedBytes))

	for i := 0; i < len(paddedBytes); i += block.BlockSize() {
		if i+block.BlockSize() <= len(paddedBytes) {
			block.Encrypt(cipherText[i:i+block.BlockSize()], paddedBytes[i:i+block.BlockSize()])
		}
	}

	return cipherText
}

func toBinaryString(data []byte) string {
	var binStr strings.Builder
	for _, b := range data {
		binStr.WriteString(fmt.Sprintf("%08b ", b))
	}
	return binStr.String()
}

func decrpyt(encryptedBase64 string, key []byte, isPrintBytes bool) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plainBytes := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += block.BlockSize() {
		block.Decrypt(plainBytes[i:i+block.BlockSize()], cipherText[i:i+block.BlockSize()])
	}

	if isPrintBytes {
		fmt.Println("解密後的原始位元組陣列：")
		for i, b := range plainBytes {
			fmt.Printf("位元組 %d: 值=%d (0x%02x) ASCII=%q\n", i, b, b, b)
		}

		lastBlockIndex := len(plainBytes) - block.BlockSize()
		if lastBlockIndex < 0 {
			lastBlockIndex = 0
		}

		fmt.Println("\n最後一個區塊的位元組:")
		for i := lastBlockIndex; i < len(plainBytes); i++ {
			fmt.Printf("位元組 %d: 值=%d (0x%02x) ASCII=%q\n", i, plainBytes[i], plainBytes[i], plainBytes[i])
		}

		if len(plainBytes) > 0 {
			paddingValue := plainBytes[len(plainBytes)-1]
			fmt.Printf("\n最後一個字節值(可能是填充長度): %d (0x%02x)\n", paddingValue, paddingValue)
		}
	}

	unpaddedBytes := unpad(plainBytes)
	return string(unpaddedBytes), nil
}

func pad(data []byte, blockSize int) []byte {
	// 需要添加的填充字節數是 16 - (34 % 16) = 16 - 2 = 14
	// 添加 14 個值為 14 (0x0E)
	padding := blockSize - (len(data) % blockSize)
	padText := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

func unpad(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return data
	}

	// 獲取最後一個字節，它表示填充的數量
	padding := int(data[length-1])

	// 驗證填充是否有效
	if padding > length || padding == 0 {
		return data
	}

	// 檢查所有填充字節是否一致
	for i := length - padding; i < length; i++ {
		if int(data[i]) != padding {
			return data
		}
	}

	// 返回去除填充後的數據
	return data[:length-padding]
}

func main() {
	plainText := "AES block of 16"
	// plainText := "AES"
	cipherBytes := encrypt(plainText)

	fmt.Println("Hex :")
	fmt.Printf("%x\n", cipherBytes)

	binaryStr := toBinaryString(cipherBytes)
	fmt.Println("\nBinary :")
	fmt.Println(binaryStr)

	base64Str := base64.StdEncoding.EncodeToString(cipherBytes)
	fmt.Println("\nBase64 :")
	fmt.Println(base64Str)

	fmt.Println("\nDecrypt :")
	origin, err := decrpyt(base64Str, []byte(secretKey16), false)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(origin)

	// ------------------------------------------------------

	plainText = "This is some sensitive information!!"
	cipher := encrypt(plainText)
	fmt.Println("\nHex :")
	fmt.Printf("%x\n", cipher)

	binaryStr = toBinaryString(cipher)
	fmt.Println("\nBinary :")
	fmt.Println(binaryStr)

	base64Str = base64.StdEncoding.EncodeToString([]byte(cipher))
	fmt.Println("\nBase64 :")
	fmt.Println(base64Str)

	fmt.Println("\nDecrypt :")
	origin, err = decrpyt(base64Str, []byte(secretKey16), false)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println(origin)

	// This will cause an error since the, plaintext is less than 16 bytes.
	// cipher = encrypt("Hi")
	// fmt.Printf("%x \n", cipher)

	// ------------------------------------------------------

}
