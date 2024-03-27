package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateData(msg string) <-chan string {
	c := make(chan string)

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println("generateData for loop ...")

			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		close(c)
	}()

	return c
}

func main() {

	data := generateData("Some Data")

	value := generateData("Some Value")

	for i := 0; i < 10; i++ {
		fmt.Println(<-data)
		fmt.Println(<-value)
	}

	// or can simply use the for range
	// for msg := range data {
	// 	fmt.Println("data for loop ...")
	// 	fmt.Println(msg)
	// }

	// for msg := range value {
	// 	fmt.Println("value for loop ...")
	// 	fmt.Println(msg)
	// }

	fmt.Println("Leaving...")
}
