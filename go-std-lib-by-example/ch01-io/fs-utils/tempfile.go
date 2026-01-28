package main

import (
	"fmt"
	"os"
)

func TempFileDemo() {
	// ioutil.TempFile("", "tempFile.txt")
	tempFile, err := os.CreateTemp("../1-2", "tempFile*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	// defer os.Remove(tempFile.Name())

	fmt.Println("Temporary file created:", tempFile.Name())

	_, err = tempFile.Write([]byte("TempFile: Hello, Go!"))
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}

	err = tempFile.Close()
	if err != nil {
		fmt.Println("Error closing temp file:", err)
		return
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Println("Error reading from temp file:", err)
		return
	}
	fmt.Println("Data read from temp file:", string(data))
}
