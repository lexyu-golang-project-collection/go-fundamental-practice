package main

import (
	"fmt"
	"unsafe"
)

func main() {

	x := 1
	y := "string"

	size := unsafe.Sizeof(x)
	size2 := unsafe.Sizeof(y)
	len_y := len(y)

	fmt.Println("alloc bytes =", size)
	fmt.Println("alloc bytes =", size2)
	fmt.Println("len =", len_y)
}
