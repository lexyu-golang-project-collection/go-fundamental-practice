package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "繁體中文"

	fmt.Println("str idx = ", strings.Index(str, "文"))

	fmt.Println("str idx = ", UTF8Index(str, "文"))

}

func UTF8Index(str, sub string) int {
	asciiPos := strings.Index(str, sub)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}

	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++

		if totalSize == asciiPos {
			return pos
		}

	}
	return pos
}
