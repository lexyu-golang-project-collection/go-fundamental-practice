package main

import "fmt"

func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int { return x + y }
	fmt.Println("5 + 4 =", intSum(5, 4))

	sample := 1
	changeVar := func() { sample += 1 }

	changeVar()
	fmt.Println("sample =", sample)

	useFunc(sumVals, 5, 8)
}
