package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("Hi, high value software engineer")
	hashAndBroadcast(bytes.NewReader(payload))
}

type hashReader struct {
	bytes.Reader
	buf *bytes.Buffer
}

type NewHashReader(b []byte) *hashReader {

}

func hashAndBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))

	// read already, will empty
	return broadcast(r)
}

func broadcast(r io.Reader) error {

	b, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	fmt.Println("string of the bytes:", string(b))

	return nil
}
