package main

import "fmt"

func main() {
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2 Index 0 = ", arr2[0])
	fmt.Println("arr2 Length 0 =", len(arr2))
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("arr2 index %d : value %d\n", i, v)
	}
}
