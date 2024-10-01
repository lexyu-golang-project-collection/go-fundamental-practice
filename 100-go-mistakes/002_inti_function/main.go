package main

import (
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("Start...")
	k, v := Store("K", "V")
	fmt.Println(k, v)
}
