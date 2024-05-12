package main

import (
	"fmt"
	"sync"
)

func main() {
	var c chan int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		c <- 1
		wg.Done()
	}()

	go func() {
		val := <-c
		fmt.Println(val)
		wg.Done()
	}()
	wg.Wait()

}
