package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("Hi, high value software engineer")
	// hashAndBroadcast(bytes.NewReader(payload))
	hashAndBroadcast(newHashReader(payload))
}

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func newHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	/*
		b, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		hash := sha1.Sum(b)
	*/

	// hash := r.(*hashReader).hash()

	hash := r.hash()

	fmt.Println(hash)
	// fmt.Println(hex.EncodeToString(hash[:]))

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
