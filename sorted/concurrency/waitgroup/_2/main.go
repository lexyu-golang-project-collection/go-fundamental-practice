package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter when the function exits

	fmt.Printf("Worker %d starting\n", id)
	// Simulating some work
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d working...\n", id)
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 5

	// Add number of goroutines to WaitGroup counter
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		// Pass the WaitGroup by pointer to the worker function
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers have finished")
}
