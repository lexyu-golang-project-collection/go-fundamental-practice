package main

import (
	"fmt"
)

type Color int

const (
	Black Color = iota
	White
	Red
)

func (c Color) String() string {
	return [...]string{"Black", "White", "Red"}[c]
}

func main() {
	fmt.Printf("%d, %d, %d\n", Black, White, Red)
	fmt.Printf("%s, %s, %s\n", Black, White, Red)
}
