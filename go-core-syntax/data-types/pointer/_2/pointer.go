package main

import "fmt"

func changeValWithoutPointer(f3 int) int {
	f3 += 1
	return f3
}

func changeValByPointer(myPtr *int) {
	*myPtr = 12
}

func doubleArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var NumSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / NumSize)
}

func main() {
	fmt.Println("WithoutPointer ========================================================================")
	f3 := 5
	fmt.Println("f3 before func :", f3)
	changeValWithoutPointer(f3)
	fmt.Println("f3 after func :", f3)

	fmt.Println("With Pointer ========================================================================")
	f4 := 10
	fmt.Println("f4 before func :", f4)
	changeValByPointer(&f4)
	fmt.Println("f4 after func :", f4)

	fmt.Println("========================================================================")

	var f4Ptr *int = &f4
	fmt.Println("&f4 Address :", &f4)
	fmt.Println("f4Ptr Address :", f4Ptr)
	fmt.Println("*f4Ptr Value :", *f4Ptr)
	*f4Ptr = 11
	fmt.Println("*f4Ptr Value :", *f4Ptr)

	fmt.Println("========================================================================")

	pArr := [4]int{1, 2, 3, 4}
	doubleArrVals(&pArr)
	fmt.Println(pArr)

	fmt.Println("========================================================================")

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}
