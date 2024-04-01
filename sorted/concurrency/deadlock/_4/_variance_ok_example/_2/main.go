package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- "ch2 value"
		ch1 <- "ch1 value"
	}()

	fmt.Println("", <-ch2)
	fmt.Println("", <-ch1)

	time.Sleep(time.Second * 2)
}
