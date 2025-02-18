package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	secretKey = "secret_key"
)

func generateRandomMessage() ([]byte, error) {
	message := make([]byte, 20)
	_, err := rand.Read(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// HMAC + SHA256
func generateHMAC(secretKey string, message []byte) string {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(message)
	hmacSum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmacSum)
}

func main() {

	msg, err := generateRandomMessage()

	if err != nil {
		fmt.Println("failed to generate message = ", err)
		return
	}
	fmt.Println("Message = ", hex.EncodeToString(msg))

	token := generateHMAC(secretKey, msg)
	fmt.Println("Token = ", token)
}
