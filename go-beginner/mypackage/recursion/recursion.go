package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func Recursion(num int) {
	g.PF("Factorial(%d) = %d", num, factorial(num))
}
