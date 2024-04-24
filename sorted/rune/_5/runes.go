package main

import "fmt"

func main() {
	str := "I Love Go"
	runes := []rune(str)

	fmt.Println(string(runes))
	fmt.Println(len(runes))
}
