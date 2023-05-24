package mypackage

func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func Closures() {
	intSum := func(x, y int) int { return x + y }
	pl("5 + 4 =", intSum(5, 4))
	sample := 1
	changeVar := func() { sample += 1 }
	changeVar()
	pl("sample =", sample)

	useFunc(sumVals, 5, 8)
}
