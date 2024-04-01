package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutexA, mutexB sync.Mutex

	// First goroutine locks mutexA then waits before locking mutexB
	wg.Add(1)
	go func() {
		defer wg.Done()
		mutexA.Lock()
		defer mutexA.Unlock()
		fmt.Println("Goroutine 1 locked mutexA")

		// Introducing some delay to increase the likelihood of deadlock
		// In real-world scenarios, this delay could be due to various operations
		// such as IO, network calls, etc.
		for i := 0; i < 1000000000; i++ {
		}

		mutexB.Lock()
		defer mutexB.Unlock()
		fmt.Println("Goroutine 1 locked mutexB")
	}()

	// Second goroutine locks mutexB then waits before locking mutexA
	wg.Add(1)
	go func() {
		defer wg.Done()
		mutexB.Lock()
		defer mutexB.Unlock()
		fmt.Println("Goroutine 2 locked mutexB")

		// Introducing some delay
		for i := 0; i < 1000000000; i++ {
		}

		mutexA.Lock()
		defer mutexA.Unlock()
		fmt.Println("Goroutine 2 locked mutexA")
	}()

	wg.Wait()
	fmt.Println("Both goroutines finished")
}
