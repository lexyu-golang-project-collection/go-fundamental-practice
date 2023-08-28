package mypackage

import (
	"fmt"
	"unicode/utf8"

	g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"
)

func Runes() {
	rStr := "abcdefg"
	g.PL("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
