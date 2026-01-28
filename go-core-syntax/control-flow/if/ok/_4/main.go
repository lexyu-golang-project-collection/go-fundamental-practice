package main

import "fmt"

func main() {
	ch := make(chan string) //initializing a channel

	go func() {
		for i := 1; i <= 3; i++ {
			msg := fmt.Sprintf("Current num is #%v", i)
			ch <- msg
		}
		close(ch) //closing the channel lets range know exactly ho much data is there
	}()

	// example 1
	for msg := range ch {
		fmt.Println("received message : ", msg)
	}

	// example 2
	msg := <-ch
	fmt.Printf("Message from a closed channel : %#v\n", msg)

	// example 3
	msg, ok := <-ch // recieving value through a channel with comma ok
	fmt.Printf("Message : %#v, (was it closed? -> %#v)\n", msg, !ok)
}
