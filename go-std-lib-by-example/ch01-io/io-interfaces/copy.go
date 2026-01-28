package main

import (
	"fmt"
	"io"
	"os"
)

func CopyStdDemo() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF -- Out")
}

func CopyNDemo() {
	io.Copy(os.Stdout, os.Stdin)
	fmt.Println("Got EOF -- Out")
}
