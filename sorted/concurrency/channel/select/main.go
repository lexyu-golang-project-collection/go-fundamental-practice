package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Get stop signal - stop channel")
				return
			default:
				fmt.Println("Keep working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	fmt.Println("Stop the goroutine")
	stop <- true

	time.Sleep(3 * time.Second)
	fmt.Println("Main Over")
}
