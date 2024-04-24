package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	m[3] = "four"

	if _, ok := m[1]; ok {
		fmt.Println("1 Key Exist")
	}

	if _, ok := m[999]; ok {
		fmt.Println("999 Key Exist")
	} else {
		fmt.Println("999 Key Not Exist")
	}
}
