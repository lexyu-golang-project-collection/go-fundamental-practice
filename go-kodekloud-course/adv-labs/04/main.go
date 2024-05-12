package main

import (
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1
		wg.Done()
	}()
	wg.Wait()
	close(ch) // deadlock
	<-ch
}
