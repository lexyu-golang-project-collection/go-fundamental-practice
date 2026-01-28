package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert[S any, D any](src []S, mapper func(s S) D) []D {
	r := make([]D, 0, len(src))
	for _, v := range src {
		r = append(r, mapper(v))
	}
	return r
}

func StringToUpperAndDoPrint(arr []string) {
	slice := Convert(arr, func(str string) string { return strings.ToUpper(str) })
	for idx, v := range slice {
		fmt.Println("Index + v = ", strconv.Itoa(idx)+" "+v)
	}
}

func main() {
	ss := []string{"hello", "world", "my", "freind"}

	StringToUpperAndDoPrint(ss)
}
