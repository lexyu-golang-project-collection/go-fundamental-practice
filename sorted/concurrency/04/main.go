package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	// numRounds := 10
	// go throwingNinjaStar(channel, numRounds)
	go throwingNinjaStar(channel)

	/*
		fmt.Println(<-channel)
		fmt.Println("------------------------------------")
		for i := 0; i < numRounds; i++ {
			fmt.Println(<-channel)
		}
	*/

	/*
		for message := range channel {
			fmt.Println(message)
		}
	*/

	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}

}

func throwingNinjaStar(channel chan string) {
	// rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))

	numRounds := 10

	for i := 0; i < numRounds; i++ {

		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)

	}

	close(channel)
}
