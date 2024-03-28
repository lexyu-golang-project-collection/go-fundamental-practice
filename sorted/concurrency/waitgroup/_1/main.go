package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(3 * time.Second)

		fmt.Println("Job - 1 Done.")

		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)

		fmt.Println("Job - 2 Done.")

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All Job Done.")
}
