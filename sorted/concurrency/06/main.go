package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type goroutineIDKey struct{}

var goroutineIDCounter uint64

func nextGoroutineID() uint64 {
	return atomic.AddUint64(&goroutineIDCounter, 1)
}

func withGoroutineID(ctx context.Context) context.Context {
	return context.WithValue(ctx, goroutineIDKey{}, nextGoroutineID())
}

func goroutineID(ctx context.Context) uint64 {
	id, _ := ctx.Value(goroutineIDKey{}).(uint64)
	return id
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	id := goroutineID(ctx)
	fmt.Printf("Goroutine %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			ctx := withGoroutineID(context.Background())
			worker(ctx, &wg)
		}()
	}
	wg.Wait()
}
