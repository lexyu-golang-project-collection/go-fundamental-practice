package main

import (
	"fmt"
)

func caesarCipherShiftRune(r rune) rune {
	return r + 1
}

func main() {
	fmt.Printf("%c\n", caesarCipherShiftRune('A')) // B
	fmt.Printf("%c\n", caesarCipherShiftRune('B')) // C
	fmt.Printf("%c\n", caesarCipherShiftRune('C')) // D
	fmt.Printf("%c\n", caesarCipherShiftRune('M')) // N
	fmt.Printf("%c\n", caesarCipherShiftRune('T')) // U
	fmt.Printf("%c\n", caesarCipherShiftRune('Z')) // [
}
