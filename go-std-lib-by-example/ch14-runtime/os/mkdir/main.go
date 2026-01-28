package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("folder", 777)
	if err != nil {
		fmt.Println("Error creating:", err)
	} else {
		fmt.Println("created successfully")
	}
}
