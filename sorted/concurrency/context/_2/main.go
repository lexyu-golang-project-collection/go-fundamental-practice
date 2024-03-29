package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "Node-1", 5)
	go worker(ctx, "Node-2", 2)
	go worker(ctx, "Node-3", 10)

	time.Sleep(20 * time.Second)

	fmt.Println("Stop the goroutine")
	cancel()

	time.Sleep(3 * time.Second)
	fmt.Println("Main Over")
}

func worker(ctx context.Context, name string, workTime int) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "Get stop signal - stop channel")
				return
			default:
				fmt.Println(name, "Keep working")
				time.Sleep(time.Second * time.Duration(workTime))
			}
		}
	}()
}
