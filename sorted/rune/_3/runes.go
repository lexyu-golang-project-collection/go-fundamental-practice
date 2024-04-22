package main

import "fmt"

func main() {
	s1 := "截取部分中文字串"
	fmt.Println(s1[:2])

	s2 := "截取部分中文字串"
	res := []rune(s2)
	fmt.Println(string(res[:4]))
}
