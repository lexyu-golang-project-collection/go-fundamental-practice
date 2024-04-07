package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Remove Specific Index
	numbers = append(numbers[:2], numbers[3:]...)

	fmt.Println(numbers)
}
