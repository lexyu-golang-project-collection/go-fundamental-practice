package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evil_ninjas := []string{"T", "J", "B", "A"}
	for _, evil_ninja := range evil_ninjas {
		go attack(evil_ninja)
	}

	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at ", target)
	time.Sleep(time.Second)
}
