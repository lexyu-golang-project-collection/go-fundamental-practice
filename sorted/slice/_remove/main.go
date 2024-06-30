package main

import "fmt"

func removeFromSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeFromSlice2(slice []int, index int) []int {
	sliceLength := len(slice)
	slice[sliceLength-1], slice[index] = slice[index], slice[sliceLength-1]
	return slice[:sliceLength-1]
}

func deleteElement(slice []int, index int) []int {
	copied := []int{}
	copied = append(copied, slice[:index]...)
	return append(copied, slice[index+1:]...)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// Remove Specific Index
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

	nums := []int{6, 7, 8, 9, 10}
	after := removeFromSlice(nums, 2)
	fmt.Println(nums)
	fmt.Println(after)

	nums2 := []int{100, 99, 98, 97, 96}
	after2 := deleteElement(nums2, 2)
	fmt.Println(nums2)
	fmt.Println(after2)

	nums3 := []int{51, 52, 53, 54, 55}
	after3 := removeFromSlice2(nums3, 2)
	fmt.Println(nums3)
	fmt.Println(after3)
}
