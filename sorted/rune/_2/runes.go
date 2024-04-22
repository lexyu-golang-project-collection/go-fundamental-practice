package main

import "fmt"

func main() {
	f := "中文字"
	fmt.Println([]rune(f))
	fmt.Println([]byte(f))
	/*
		[20013 25991 23383]
		[228 184 173 230 150 135 229 173 151]
	*/
}
