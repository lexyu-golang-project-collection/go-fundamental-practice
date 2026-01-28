package main

import "fmt"

func main() {
	first := "fisrt"
	fmt.Println([]rune(first))
	fmt.Println([]byte(first))
	/*
		[102 105 115 114 116]
		[102 105 115 114 116]
	*/
}
