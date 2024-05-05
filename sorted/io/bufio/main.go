package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	readFile, err := os.Open("../../os/new-demo.txt")
	if err != nil {
		fmt.Println("Error Opening", err)
		return
	}
	defer readFile.Close()

	reader := bufio.NewReader(readFile)
	for {

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error Reading", err)
			return
		}

		fmt.Print(line)
	}
}
