package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 500; i++ {
		fmt.Printf("Trying to produce: %d\n", i)
		ch <- i // This will block after buffer is full
		fmt.Printf("Produced: %d\n", i)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)

	go producer(ch, &wg)

	// Wait for producer to finish or deadlock
	wg.Wait()
	fmt.Println("Program finished")
}
