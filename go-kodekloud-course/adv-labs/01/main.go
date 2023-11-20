package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10) // <size>
	go produce(ch, &wg)
	wg.Wait()
}

func produce(ch chan int, wg *sync.WaitGroup) {
	for i := 10; i <= 100; i += 10 {
		ch <- i
	}
	fmt.Println("Exiting Produce")
	close(ch)
	go consume(ch, wg)
	wg.Done()
}

func consume(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println("Received: ", val)
	}
	fmt.Println("Exiting Consume")
	wg.Done()
}
