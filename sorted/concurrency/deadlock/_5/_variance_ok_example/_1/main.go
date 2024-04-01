package main

import (
	"fmt"
	"time"
)

func main() {
	chs := make(chan string, 2)
	chs <- "first"
	chs <- "second"

	go func() {
		for ch := range chs {
			fmt.Println(ch)
		}
	}()

	time.Sleep(3 * time.Second)
}
