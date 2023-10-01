package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Split(text, " ")
	words2 := strings.Fields(text)
	frequency := make(map[string]int)
	fmt.Println(words)
	fmt.Println(words2)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))

}
