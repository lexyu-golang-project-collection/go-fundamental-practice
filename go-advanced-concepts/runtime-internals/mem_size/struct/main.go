package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	myBool  bool    // 1 byte
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
}

type myStructOptimized struct {
	myFloat float64 // 8 bytes
	myBool  bool    // 1 byte
	myInt   int32   // 4 bytes
}

func main() {

	a := myStruct{}
	fmt.Println(unsafe.Sizeof(a)) //24

	b := myStructOptimized{}
	fmt.Println(unsafe.Sizeof(b)) //16
}
