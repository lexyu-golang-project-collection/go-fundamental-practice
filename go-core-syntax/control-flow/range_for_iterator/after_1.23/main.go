package main

import (
	"fmt"
	"iter"
	"strings"
)

func Backward(arr []string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := len(arr) - 1; i >= 0; i-- {
			yield(strings.ToUpper(arr[i]))
		}
	}
}

func ToUppperByIter(arr []string) {
	for e := range Backward(arr) {
		fmt.Println("v = ", e)
	}
}

func main() {
	ss := []string{"hello", "world", "go"}

	ToUppperByIter(ss)
}
