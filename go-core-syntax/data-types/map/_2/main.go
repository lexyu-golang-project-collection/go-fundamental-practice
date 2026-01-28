package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	m[3] = "four"

	for _, v := range m {
		fmt.Println(v)
	}
}
