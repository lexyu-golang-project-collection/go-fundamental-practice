package main

import "fmt"

func main() {
	chs := make(chan string, 2)
	chs <- "first"
	chs <- "second"

	for ch := range chs {
		fmt.Println(ch)
	}
}
