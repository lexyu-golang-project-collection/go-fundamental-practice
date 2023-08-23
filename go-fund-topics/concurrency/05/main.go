package main

import (
	"fmt"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go captainElect(channel1, "Ninja1")
	go captainElect(channel2, "Ninja2")

	// fmt.Println(<-channel1)
	// fmt.Println(<-channel2)

	select { // random choose one
	case msg := <-channel1:
		fmt.Println(msg)
	case msg := <-channel2:
		fmt.Println(msg)
	default:
		fmt.Println("Neither channel1 or channel2")
	}

	roughlyFair()
}

func captainElect(ninja chan string, ch string) {
	// time.Sleep(3 * time.Second)
	ninja <- ch
}

func roughlyFair() {
	ch1 := make(chan interface{})
	close(ch1)
	ch2 := make(chan interface{})
	close(ch2)

	var ch1Count, ch2Count int
	for i := 0; i < 1000; i++ {
		select { // random choose one
		case <-ch1:
			ch1Count++
		case <-ch2:
			ch2Count++
		}
	}
	fmt.Printf("ch1Count: %d, ch2Count: %d\n", ch1Count, ch2Count)
}
