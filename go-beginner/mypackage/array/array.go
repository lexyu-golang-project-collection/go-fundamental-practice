package mypackage

import g "github.com/lexyu-golang-project-collection/go-fundamental-practice/go-beginner/mypackage/global"

func Array() {
	var arr1 [5]int
	arr1[0] = 1
	g.PL("arr1 Index 0 = ", arr1[0])

	arr2 := [5]int{1, 2, 3, 4, 5}
	g.PL("arr2 Index 0 = ", arr2[0])
	g.PL("arr2 Length 0 =", len(arr2))
	for i := 0; i < len(arr2); i++ {
		g.PF("arr2[%d] = %d\n", i, arr2[i])
	}
	for i, v := range arr2 {
		g.PF("arr2 index %d : value %d\n", i, v)
	}

	arr3 := [2][2]int{
		{5, 6},
		{7, 8},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			g.PF("arr3[%d][%d] = %d\n", i, j, arr3[i][j])
		}
	}

	g.PL("----------------------------------------------------------------")

	p := []int{1, 2, 3}
	qwe := 0
	qwe, p[qwe] = 1, 2
	g.PL("Exchange ", p, qwe)
}
