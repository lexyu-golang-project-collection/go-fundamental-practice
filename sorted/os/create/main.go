package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("Hi, there!")
}
