package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "First Msg"
	channel <- "Second Msg"

	// go func() {
	// 	channel <- "First Msg - func"
	// }()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
