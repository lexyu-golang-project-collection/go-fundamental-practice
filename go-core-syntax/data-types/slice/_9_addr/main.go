package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	scores := []int{1, 2, 3, 4, 5}
	fmt.Printf("&scores = %p\n", &scores)

	scoresArrayPtr := reflect.ValueOf(scores).Pointer()
	fmt.Printf("Memory address of scores' underlying array: %p\n", unsafe.Pointer(scoresArrayPtr))

	// newSlice := scores
	// newSlice := scores[0:] // `:` newSclice point to same address array
	newSlice := scores[1:]
	newSliceArrayPtr := reflect.ValueOf(newSlice).Pointer()
	fmt.Printf("Memory address of newSlice's underlying array: %p\n", unsafe.Pointer(newSliceArrayPtr))
	fmt.Printf("&newSlice = %p\n", &newSlice)

	// Printing the capacity of both slices
	fmt.Println("scores capacity:", cap(scores))
	fmt.Println("newSlice capacity:", cap(newSlice))
	fmt.Println(newSlice) // 3, 4

	newSlice[0] = 999   // 把原本 scores 中 index 值為 3 的元素改成 999
	fmt.Println(scores) // 1, 2, 999, 4, 5
}
