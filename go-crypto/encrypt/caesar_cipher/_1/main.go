package main

import (
	"fmt"
	"strings"
)

/*
B
C
D
N
U
[
*/
func caesarCipherShiftRunePart1(r rune) rune {
	return r + 1
}

func caesarCipherShiftRunePart2(r rune) rune {
	return (((r - 'A') + 1) % 26) + 'A'
}

func caesarCipherFinal(r rune, shift uint) rune {
	s := rune(shift % 26)

	if s == 0 {
		return r
	}

	return (((r - 'A') + s) % 26) + 'A'
}

func caesarCipherShiftRuneDemo1() {
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('A'))
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('B'))
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('C'))
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('M'))
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('T'))
	fmt.Printf("%c\n", caesarCipherShiftRunePart1('Z'))
}

func caesarCipherShiftRuneDemo2() {
	const input = "HELLO, WORLD!"
	var output string

	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherShiftRunePart2(r)
		}
		output += string(r)
	}

	fmt.Println(output)
}

func caesarCipherShiftRuneDemo3(input string, shift uint) {
	var output string

	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherFinal(r, shift)
		}
		output += string(r)
	}

	fmt.Println(output)
}

func caesarCipherEncrypt(value string, shift uint) string {
	var builder strings.Builder
	value = strings.ToUpper(value)

	for _, r := range value {
		if r >= 'A' && r <= 'Z' {
			r = caesarCipherFinal(r, shift)
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

func caesarCipherDecrypt(value string, shift uint) string {
	return caesarCipherEncrypt(value, 26-(shift%26))
}

func main() {
	// caesarCipherShiftRuneDemo1()

	// caesarCipherShiftRuneDemo2()

	// caesarCipherShiftRuneDemo3("THIS IS A TEST", 13)

	encrypt_result := caesarCipherEncrypt("Hello, World!", 13)
	fmt.Println(encrypt_result)

	decrypt_result := caesarCipherDecrypt(encrypt_result, 13)
	fmt.Println(decrypt_result)

}
