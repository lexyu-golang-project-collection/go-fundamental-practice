package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	fmt.Println(i * i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go calculateSquare(i, &wg)
	}

	elapsed := time.Since(start)
	wg.Wait()

	fmt.Println("Function took: ", elapsed)
}
