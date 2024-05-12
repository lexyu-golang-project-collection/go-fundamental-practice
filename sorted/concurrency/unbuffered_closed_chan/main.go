package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", val)
	}

	// Further reads from the closed channel
	// will yield the zero value of int.
	fmt.Println("Read from closed channel:", <-ch)
}
