package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(250 * time.Millisecond)
}

func start() {
	go process()
	fmt.Println("In start")
}

func process() {
	fmt.Println("In process")
}
