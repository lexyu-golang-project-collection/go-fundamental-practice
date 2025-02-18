package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Save private key to a file
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()

	// Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Write private key to file
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		panic(err)
	}

	// Save public key to a file
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()

	// Marshal public key to DER format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// Encode public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Write public key to file
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		panic(err)
	}
}
