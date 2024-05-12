package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	// Create first goroutine here
	go printNumbers(&wg)
	// Create second goroutine here
	go printLetters(&wg)

	// Wait for both goroutines to finish here
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Println(string(i))
	}
	wg.Done()
}
