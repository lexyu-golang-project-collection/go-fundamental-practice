package concurrency

import (
	"fmt"
	"time"
)

func printTo150() {
	for i := 1; i <= 150; i++ {
		fmt.Println("Fun 1:", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Fun 2:", i)
	}
}
func Goroutines() {
	go printTo150()
	go printTo10()
	time.Sleep(2 * time.Second)
}

func Empty() {

}
