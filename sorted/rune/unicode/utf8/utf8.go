package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	// 計算字串中的 rune 數量
	count := utf8.RuneCountInString(str)
	fmt.Printf("Rune count: %d\n", count)

	// 遍歷並解碼每個 rune
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c %d\n", r, size)
		i += size
	}

	// 編碼 rune
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	fmt.Printf("Encoded bytes: %v, length: %d\n", buf[:n], n)
	fmt.Printf("字元: %c\n", '世')
	fmt.Printf("Unicode 碼點（十進制）: %d\n", '世')
	fmt.Printf("Unicode 碼點（十六進制）: %X\n", '世')
	fmt.Printf("UTF-8 編碼後的位元組（十六進制）: [% X]\n", buf[:n])
	fmt.Printf("UTF-8 編碼後的位元組（二進制）: ")
	for _, b := range buf[:n] {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
}
