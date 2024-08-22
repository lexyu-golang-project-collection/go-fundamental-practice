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
				fmt.Println("Goroutine-1 canceled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine-1 is working...")
				time.Sleep(1000 * time.Millisecond)

			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine-2 canceled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine-2 is working...")
				time.Sleep(1000 * time.Millisecond)

			}
		}
	}()

	fmt.Println("Main function is working...")
	time.Sleep(5000 * time.Millisecond)

	fmt.Println("Canceling context...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main Function Done.")

}
