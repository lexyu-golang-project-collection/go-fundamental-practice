package main

import (
	"fmt"
	"strings"
)

func ReadAtDemo() {
	reader := strings.NewReader("Test Go ReadAt Method.")
	reader2 := strings.NewReader("中文字串")

	readAndPrint := func(reader *strings.Reader) {
		p := make([]byte, 6)
		n, err := reader.ReadAt(p, 0)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s == %d\n", p, n)
	}

	readAndPrint(reader)
	readAndPrint(reader2)
}
