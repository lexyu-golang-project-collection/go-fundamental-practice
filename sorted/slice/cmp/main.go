package main

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/slices"
)

func main() {
	sliceCmpDemo()
	fmt.Println("----------------------------------")
	reflectDemo()
}

func sliceCmpDemo() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	// Compare slices using slices.Equal
	fmt.Println(slices.Equal(a, b)) // Output: true
	fmt.Println(slices.Equal(a, c)) // Output: false
}

func reflectDemo() {
	var x []int
	y := []int{}

	fmt.Println(x == nil)                // Output: true
	fmt.Println(y == nil)                // Output: false
	fmt.Println(reflect.DeepEqual(x, y)) // Output: false
}
