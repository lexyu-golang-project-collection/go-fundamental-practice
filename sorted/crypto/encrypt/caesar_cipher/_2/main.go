package main

import (
	"fmt"
	"strings"
)

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {

	encrypt_result := strings.Map(func(r rune) rune {
		return caesar(r, 13)
	}, "hello, world!")

	fmt.Println(encrypt_result)
}
