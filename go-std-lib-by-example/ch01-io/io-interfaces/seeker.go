package main

import (
	"fmt"
	"io"
	"strings"
)

func SeekerDemo() {
	reader := strings.NewReader("這是一段繁體中文句子")
	reader.Seek(-9, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
