package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"

func Pointers() {
	f3 := 5
	g.PL("f3 before func :", f3)
	changeVal(f3)
	g.PL("f3 after func :", f3)

	f4 := 10
	g.PL("f4 before func :", f4)
	changeVal2(&f4)
	g.PL("f4 after func :", f4)

	var f4Ptr *int = &f4
	g.PL("&f4 Address :", &f4)
	g.PL("f4Ptr Address :", f4Ptr)
	g.PL("*f4Ptr Value :", *f4Ptr)
	*f4Ptr = 11
	g.PL("*f4Ptr Value :", *f4Ptr)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	g.PL(pArr)

	iSlice := []float64{11, 13, 17}
	g.PF("Average : %.3f\n", getAverage(iSlice...))
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
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
