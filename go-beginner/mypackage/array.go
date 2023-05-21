package mypackage

func Array() {
	var arr1 [5]int
	arr1[0] = 1
	pl("arr1 Index 0 = ", arr1[0])

	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("arr2 Index 0 = ", arr2[0])
	pl("arr2 Length 0 =", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pf("arr2[%d] = %d\n", i, arr2[i])
	}
	for i, v := range arr2 {
		pf("arr2 index %d : value %d\n", i, v)
	}

	arr3 := [2][2]int{
		{5, 6},
		{7, 8},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pf("arr3[%d][%d] = %d\n", i, j, arr3[i][j])
		}
	}
}
