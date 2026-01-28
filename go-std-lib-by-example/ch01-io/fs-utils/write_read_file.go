package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteAndReadFileDemo() {

	/*
		baseDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting the current working directory:", err)
			return
		}
	*/

	// filePath := filepath.Join(baseDir, "01", "writeFile.txt")

	relativePath := filepath.Join("..", "..", "01", "writeFile.txt")

	err := os.MkdirAll(filepath.Dir(relativePath), os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	data := []byte("Hello, Go")

	// err := ioutil.WriteFile("/01/writeFile.txt", data, 0644)
	// err := os.WriteFile("../../01/writeFile.txt", data, 0644)
	err = os.WriteFile(relativePath, data, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to file successfully.")

	// readData, err := ioutil.ReadFile("example.txt")
	readData, err := os.ReadFile(relativePath)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("Data read from file:", string(readData))
}
