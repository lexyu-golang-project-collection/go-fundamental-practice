package main

import (
	"fmt"
	"os"
)

func main() {
	curr, _ := os.Getwd()
	fmt.Println(curr)
}
