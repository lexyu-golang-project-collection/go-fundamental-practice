package main

import (
	"fmt"
	"strings"
)

func main() {
	transform := func(r rune) rune {
		if r == 'A' || r == 'I' {
			return '_'
		}
		return r
	}

	input := "THIS IS A DOG"
	fmt.Println(input)

	result := strings.Map(transform, input)
	fmt.Println(result)
}
