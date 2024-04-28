package main

import (
	"context"
	"fmt"
	"time"
)

func Monitor(ctx context.Context) {
	for {
		fmt.Print("monitor")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go Monitor(ctx)

	time.Sleep(20 * time.Second)
}
