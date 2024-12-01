package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 字元判斷
	fmt.Println(unicode.IsLetter('A')) // true
	fmt.Println(unicode.IsDigit('1'))  // true
	fmt.Println(unicode.IsSpace(' '))  // true
	fmt.Println(unicode.IsPunct('.'))  // true

	// 字元轉換
	fmt.Println(string(unicode.ToLower('A'))) // "a"
	fmt.Println(string(unicode.ToUpper('a'))) // "A"

	// 中文字元判斷
	fmt.Println(unicode.IsLetter('中')) // true
	fmt.Println(unicode.IsSymbol('￥')) // true

	// 特殊字元判斷
	fmt.Println(unicode.IsControl('\n')) // true
	fmt.Println(unicode.IsMark('̀'))     // true
}
