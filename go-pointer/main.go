package main

import (
	"fmt"
)

func main() {

	i := 100
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	var n *int
	var s *string
	fmt.Println(n) // nil
	fmt.Println(s) // nil
	fmt.Println("s = ", &s)

	var ptr_i *int = &i
	ptr_i2 := &i
	fmt.Println("Pointer = ", ptr_i)
	fmt.Println("Pointer = ", ptr_i2)

	s1 := "Hello"
	fmt.Printf("%T %v \n", s1, s1)
	var b *string = &s1
	fmt.Println("b = ", b)
	var a = &s1
	fmt.Println("a = ", a)
	c := &s1
	fmt.Println("c = ", c)
	ps := &s1
	*ps = "World"
	fmt.Printf("%T %v \n", s1, s1)

	modify(&s1)
	fmt.Println(s1)

	// Slice
	slice := []int{1, 2, 3}
	fmt.Println("slice = ", slice)
	modify2(slice)
	fmt.Println("slice = ", slice)

	// Maps
	codes := make(map[string]int)
	codes["A"] = 65
	codes["F"] = 70
	fmt.Println(codes)
	modify3(codes)
	fmt.Println(codes)

	// array
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify4(arr...)
	fmt.Println(arr)

	modify5([3]int(arr))
	fmt.Println(arr)

}

func modify(s *string) {
	*s = "change the world"
}

func modify2(s []int) {
	s[0] = 10000
}

func modify3(m map[string]int) {
	m["K"] = 75
}

func modify4(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}

func modify5(numbers [3]int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}
