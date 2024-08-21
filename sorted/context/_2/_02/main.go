package main

import (
	"context"
	"fmt"
	"time"
)

func Monitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // This checks if the context is canceled
			fmt.Println("monitor stopped")
			return
		default:
			fmt.Println("monitor...")
			time.Sleep(1 * time.Second) // Add a delay to avoid spamming
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go Monitor(ctx)

	time.Sleep(10 * time.Second)
}
