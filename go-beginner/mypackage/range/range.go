package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/mypackage/global"

func Range() {
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		g.PL(num)
	}
}
