package main

import (
	"fmt"
)

func modify1(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}

func modify2(numbers [3]int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}

func main() {

	// array
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify1(arr...)
	fmt.Println(arr)

	modify2([3]int(arr))
	fmt.Println(arr)

}
