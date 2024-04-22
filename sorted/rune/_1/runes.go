package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rStr := "abcdefg"
	fmt.Println("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
