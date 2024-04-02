package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Load private key from file
	privateKeyFile, err := os.Open("private.pem")
	if err != nil {
		log.Fatal("Failed to open private key file:", err)
	}

	defer privateKeyFile.Close()

	privateKeyBytes, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		log.Fatal("Failed to read private key file:", err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		log.Fatal("Invalid private key file")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Fatal("Failed to parse private key:", err)
	}

	// Load public key from file
	publicKeyFile, err := os.Open("public.pem")
	if err != nil {
		log.Fatal("Failed to open public key file:", err)
	}

	defer publicKeyFile.Close()

	publicKeyBytes, err := ioutil.ReadAll(publicKeyFile)
	if err != nil {
		log.Fatal("Failed to read public key file:", err)
	}

	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	if publicKeyBlock == nil || publicKeyBlock.Type != "PUBLIC KEY" {
		log.Fatal("Invalid public key file")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Fatal("Failed to parse public key:", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Invalid RSA public key")
	}

	// Data to sign
	data := []byte("Hello, world!")

	// Sign the data
	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}

	// Print and encode signature to base64
	signatureEncoded := base64.StdEncoding.EncodeToString(signature)
	println("Signature:", signatureEncoded)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		println("Signature verification failed:", err)
	} else {
		println("Signature verified successfully.")
	}

}
