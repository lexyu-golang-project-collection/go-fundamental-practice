package main

import (
	"bufio"
	"os"
)

func ReaderFromDemo() {
	file, err := os.Open("../01/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}
