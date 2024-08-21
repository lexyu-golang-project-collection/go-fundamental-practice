package main

import (
	"context"
	"fmt"
	"time"
)

func Monitor(ctx context.Context) {
	for {
		fmt.Println("monitor...")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go Monitor(ctx)

	time.Sleep(10 * time.Second)
}
