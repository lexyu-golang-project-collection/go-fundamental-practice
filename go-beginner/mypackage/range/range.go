package mypackage

import g "github.com/lexyu/go-beginner/mypackage/global"

func Range() {
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		g.PL(num)
	}
}