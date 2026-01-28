package main

import (
	"bytes"
	"fmt"
	"io"
)

func NopCloserDemo() {
	buf := bytes.NewBufferString("Hello, Go!")

	readCloser := io.NopCloser(buf)

	data, err := io.ReadAll(readCloser)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Read data:", string(data))

	/*
		err = readCloser.Close()
		if err != nil {
			fmt.Println("Error closing:", err)
			return
		}
		fmt.Println("Closed successfully (no-op)")
	*/
}

// req, _ := http.NewRequest("POST", "https://example.com", io.NopCloser(bytes.NewBufferString("payload")))
