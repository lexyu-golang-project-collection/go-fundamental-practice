package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Line 1
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println("Inside Goroutine")
		// Line 2
		wg.Done()
	}()
	// Line 3
	wg.Wait()
}
