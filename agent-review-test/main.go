package main

import "fmt"

func main() {
	nums := []int{5, 0, 10, -2}
	total := SumArray(nums)
	fmt.Println("Total:", total)

	average := Divide(total, len(nums))
	fmt.Println("Average:", average)

	// Intentional potential bug: accessing out-of-bounds
	fmt.Println("First element:", nums[10])

	name := "Alice"
	PrintGreeting(name, "")
}

// SumArray returns the sum of integers in a slice
func SumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Divide two integers
func Divide(a, b int) int {
	return a / b // potential divide by zero if b == 0
}

// PrintGreeting prints greeting
func PrintGreeting(name string, title string) {
	// inconsistent formatting: sometimes uses name only
	if title != "" {
		fmt.Printf("Hello %s %s!\n", title, name)
	} else {
		fmt.Println("Hello,", name) // inconsistent comma and formatting
	}
}
