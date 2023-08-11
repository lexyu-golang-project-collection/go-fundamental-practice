package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	signal := make(chan bool)

	evilNinja := "E"
	go attack(evilNinja, signal)
	// signal <- false // deadlock
	fmt.Println(<-signal)
}

func attack(target string, signal chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
	signal <- true
}
