package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 500; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Println("n = ", n)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	start := time.Now()

	ch := make(chan int, 100)
	go producer(ch)
	consumer(ch)

	fmt.Printf("Time taken: %v\n", time.Since(start))
}
