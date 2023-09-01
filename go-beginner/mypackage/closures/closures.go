package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"

func useFunc(f func(int, int) int, x, y int) {
	g.PL("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func Closures() {
	intSum := func(x, y int) int { return x + y }
	g.PL("5 + 4 =", intSum(5, 4))
	sample := 1
	changeVar := func() { sample += 1 }
	changeVar()
	g.PL("sample =", sample)

	useFunc(sumVals, 5, 8)
}
