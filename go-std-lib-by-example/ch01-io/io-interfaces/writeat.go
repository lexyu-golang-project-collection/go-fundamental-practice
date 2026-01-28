package main

import (
	"fmt"
	"os"
)

func WriteAtDemo() {
	file, err := os.Create("../01/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Go中文討論區－這裡是額外句子")
	n, err := file.WriteAt([]byte("GoTo覆蓋字符串XX"), 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
