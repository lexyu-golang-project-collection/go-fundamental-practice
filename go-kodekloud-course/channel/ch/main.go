package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan string)

	wg.Add(2)

	go sell(ch, &wg)

	go buy(ch, &wg)

	wg.Wait()
	// time.Sleep(2 * time.Second)
}

// send data to the channel
func sell(ch chan string, wg *sync.WaitGroup) {
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
	wg.Done()
}

// receive data from the channel
func buy(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	val := <-ch
	fmt.Println("Received data - ", val)
	wg.Done()
}
