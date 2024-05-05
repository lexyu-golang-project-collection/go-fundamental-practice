package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("../demo.txt")
	if err == nil {
		fmt.Println("File exists.")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())
}
