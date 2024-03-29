package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
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
	cancel()

	time.Sleep(3 * time.Second)
	fmt.Println("Main Over")
}
