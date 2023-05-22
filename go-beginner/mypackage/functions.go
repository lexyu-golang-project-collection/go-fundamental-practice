package mypackage

import "fmt"

func Funcs() {
	// func funcName(parameters) returnType {Body}
	sayHello()
	pl(getSum(5, 4))
	// return Multiple
	pl(getTwo(4))
	// function error
	pl(getQuotient(1.5, 2))
	// Varadic Functions
	pl(getSum2(1, 2, 3, 4, 5, 6, 7))
	// passing Arrays
	vArr := []int{1, 2, 3, 4, 5}
	pl(getArraySum(vArr))
}

func sayHello() {
	pl("hello")
}
func getSum(x int, y int) int {
	return x + y
}
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}
func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divided by zero")
	} else {
		return x / y, nil
	}
}
func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
