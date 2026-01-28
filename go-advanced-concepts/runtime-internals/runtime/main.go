package main

import (
	"fmt"
	"runtime"
)

func main() {
	cores := runtime.NumCPU()
	fmt.Println(cores)
}
