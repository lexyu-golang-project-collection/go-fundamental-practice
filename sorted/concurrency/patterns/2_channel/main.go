package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go boring("boring", c)

	for i := 0; i < 7; i++ {
		fmt.Printf("Say : %q\n", <-c)
	}

	fmt.Println("Leaving...")
}
