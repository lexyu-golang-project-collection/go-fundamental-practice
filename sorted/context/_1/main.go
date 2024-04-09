package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Canceled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// Cancel the context to release resources when done,
	defer cancel()

	// Start a goroutine to perform some processing
	go doWork(ctx)

	time.Sleep(5 * time.Second)
}
