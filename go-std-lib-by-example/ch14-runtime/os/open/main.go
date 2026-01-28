package main

import (
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("../demo.txt")
	if err != nil {
		fmt.Println("Error Opening", err)
		return
	}
	defer readFile.Close()

	data := make([]byte, 100)
	_, err = readFile.Read(data)
	if err != nil {
		fmt.Println("Error Reading", err)
		return
	}
	fmt.Println("File Content:", string(data))

	file, err := os.OpenFile("../demo.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error Reading", err)
		return
	}
	fmt.Println("File Content:", string(data))

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Working Directory:", wd)
}
