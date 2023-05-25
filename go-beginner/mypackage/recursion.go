package mypackage

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func Recursion(num int) {
	pf("Factorial(%d) = %d", num, factorial(num))
}
