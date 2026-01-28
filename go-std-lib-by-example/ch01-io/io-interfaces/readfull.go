package main

import (
	"bytes"
	"fmt"
	"io"
)

func ReadFullDemo() {
	data := "Hello, Go!"
	reader := bytes.NewReader([]byte(data))

	buf := make([]byte, 10)

	n, err := io.ReadFull(reader, buf)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(buf))
	}
}
