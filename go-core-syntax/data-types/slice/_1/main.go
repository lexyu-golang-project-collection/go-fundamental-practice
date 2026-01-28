package main

import "fmt"

/* Pass by Referenece
type slice struct{
	p *pointer --> array
	len int64
	cap int64
}
*/

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	slice2 := make([]int, 5)
	fmt.Println(slice2)

	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:3]
	fmt.Println(slice3)

	slice4 := make([]int, 5, 10)
	fmt.Println(slice4)

	var slice5 []int
	slice5 = append(slice5, 1, 2, 3, 4, 5)
	fmt.Println(slice5)
}
