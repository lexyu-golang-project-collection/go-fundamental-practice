package mypackage

func Pointers() {
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)

	f4 := 10
	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)

	var f4Ptr *int = &f4
	pl("&f4 Address :", &f4)
	pl("f4Ptr Address :", f4Ptr)
	pl("*f4Ptr Value :", *f4Ptr)
	*f4Ptr = 11
	pl("*f4Ptr Value :", *f4Ptr)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	pf("Average : %.3f\n", getAverage(iSlice...))
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
